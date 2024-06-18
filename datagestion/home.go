package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/posts"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the home page.

The function gonna return:
	- an error
*/
func HomeHandle(w http.ResponseWriter, homeForm map[string]string, conn *websocket.Conn) error {
	structure.Html.Error = structure.MyError{}

	response := make(map[string]map[string]any)
	response["home"] = make(map[string]any)

	if homeForm["type"] == "post" {
		idCreator := homeForm["idCreator"]
		categories := homeForm["categorie"]
		message := homeForm["message"]
		description := homeForm["description"]

		if categories == "" {
			response["home"]["error"] = "No categorie selected"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)

			return errors.New("error")
		} else if message == "" {
			response["home"]["error"] = "Empty message"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)

			return errors.New("error")
		} else if description == "" {
			response["home"]["error"] = "Empty description"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)

			return errors.New("error")
		}

		if verificationfunction.IsValidMessage(message) && verificationfunction.IsValidMessage(description) {
			err := posts.CreateNewPost(categories, idCreator, message, description, "")
			if err != nil {
				fmt.Println(err)
				response["home"]["error"] = "Invalid message or description"

				formated, err := json.Marshal(response)
				if err != nil {
					fmt.Println(err)
				}

				conn.WriteMessage(websocket.TextMessage, formated)
				return errors.New("error")
			}

			posts.GetPosts()

			response["home"]["success"] = "nothing|now reload"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)
			return nil
		}
	} else if homeForm["filterType"] != "All" {
		i, _ := strconv.Atoi(homeForm["filterType"])

		err := posts.GetPostsByCategorieId(i)
		if err != nil {
			fmt.Println(err)
		}

		response["home"]["success"] = "filtered|now reload"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return nil
	}

	posts.GetPosts()

	response["home"]["success"] = "nothing|now reload"

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)
	return nil
}
