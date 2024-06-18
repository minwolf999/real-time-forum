package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/chat"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the chat page.

The function gonna return:
	- an error
*/
func ChatHandle(w http.ResponseWriter, chatForm map[string]string, conn *websocket.Conn) error {
	response := make(map[string]map[string]any)
	response["chat"] = make(map[string]any)

	senderId := chatForm["senderId"]
	receverId := chatForm["interlocutorId"]

	if chatForm["type"] == "isWritting" {
		dataInterlocutor, _ := users.GetUserDatasById(receverId)

		interLocConn, isOk := structure.Connected[dataInterlocutor.Username]
		if isOk {
			var res []any
			res = append(res, senderId)
			res = append(res, "typing")
			response["chat"]["success"] = res

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			interLocConn.WriteMessage(websocket.TextMessage, formated)
		}

		return nil
	}

	message := chatForm["message"]

	if message == "" {
		response["chat"]["error"] = "Invalid Message1"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	if !verificationfunction.IsValidMessage(message) {
		response["chat"]["error"] = "Invalid Message2"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	chat.AddMessage(message, senderId, receverId)
	chat.AddNotification(senderId, receverId)

	var res []any
	res = append(res, receverId)
	response["chat"]["success"] = res

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)

	recever, _ := users.GetUserDatasById(receverId)

	receverConn, isOk := structure.Connected[recever.Username]
	if isOk {
		var res []any
		res = append(res, senderId)
		response["chat"]["success"] = res

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		receverConn.WriteMessage(websocket.TextMessage, formated)
	}

	return nil
}
