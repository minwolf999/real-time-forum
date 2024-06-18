package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the register page.

The function gonna return:
	- an error
*/
func RegisterHandle(w http.ResponseWriter, registerForm map[string]string, conn *websocket.Conn) error {
	structure.User = structure.Users{}

	response := make(map[string]map[string]any)
	response["register"] = make(map[string]any)

	if registerForm["password"] != registerForm["confirm"] {
		response["register"]["error"] = "Password and Password Confirmation don't match !"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	if !verificationfunction.PasswordVerif(registerForm["password"]) {
		response["register"]["error"] = "Incorrect password ! The password must contain 8 characters, 1 uppercase letter, 1 special character, 1 number"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	if registerForm["username"] == "" || registerForm["age"] == "" || registerForm["gender"] == "" || registerForm["firstName"] == "" || registerForm["lastName"] == "" || registerForm["email"] == "" || registerForm["password"] == "" {
		response["register"]["error"] = "There is an empty field 1"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	if registerForm["gender"] != "Male" && registerForm["gender"] != "Female" {
		response["register"]["error"] = "There is an empty field 2"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(registerForm["password"]), 12)
	err := users.AddUser(registerForm["username"], registerForm["age"], registerForm["gender"], registerForm["firstName"], registerForm["lastName"], registerForm["email"], string(cryptPassword))
	if err != nil {
		fmt.Println(err)
		response["register"]["error"] = "Username or Email is already used by someone !"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	structure.User = structure.Users{}
	err = users.GetUser(registerForm["email"], string(cryptPassword))
	if err != nil {
		response["register"]["error"] = err.Error()

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return errors.New("error")
	}

	structure.Connected[structure.User.Username] = conn

	users.ConnectUser(structure.User.Id)
	structure.Html.User = structure.User

	response["register"]["success"] = structure.Html.User

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)
	return nil
}
