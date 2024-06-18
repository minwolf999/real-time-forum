package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the login page.

The function gonna return:
	- an error
*/
func LogInHandle(w http.ResponseWriter, loginForm map[string]string, conn *websocket.Conn) error {
	structure.User = structure.Users{}

	response := make(map[string]map[string]any)
	response["login"] = make(map[string]any)

	err := users.GetUserByEmail(loginForm["email"])
	if err != nil {
		response["login"]["error"] = err.Error()

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	if structure.User.Username == "" || bcrypt.CompareHashAndPassword([]byte(structure.User.Password), []byte(loginForm["password"])) != nil {
		response["login"]["error"] = "Email or Password invalid !"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	} else if structure.User.Connected == 1 {
		response["login"]["error"] = "Already connected on another device !"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	structure.Connected[structure.User.Username] = conn

	users.ConnectUser(structure.User.Id)
	structure.User.Connected = 1
	structure.Html.User = structure.User

	response["login"]["success"] = structure.Html.User

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)
	return nil
}
