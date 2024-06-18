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

The objective of this function is to manage the modification user page.

The function gonna return:
	- an error
*/
func ModificationUser(w http.ResponseWriter, modificationForm map[string]string, conn *websocket.Conn) error {
	response := make(map[string]map[string]any)
	response["modification"] = make(map[string]any)

	id := modificationForm["id"]
	username := modificationForm["username"]
	age := modificationForm["age"]
	gender := modificationForm["gender"]
	firstName := modificationForm["firstName"]
	lastName := modificationForm["lastName"]
	email := modificationForm["email"]
	password := modificationForm["password"]

	if modificationForm["currentUsername"] == username && modificationForm["currentAge"] == age && modificationForm["currentGender"] == gender && modificationForm["firstName"] == firstName && modificationForm["lastName"] == lastName && modificationForm["currentEmail"] == email && bcrypt.CompareHashAndPassword([]byte(modificationForm["currentPassword"]), []byte(password)) == nil {
		response["modification"]["error"] = "There is no modification !"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	if password != "" && !verificationfunction.PasswordVerif(password) {
		response["modification"]["error"] = "Invalid new Password"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	if username == "" || age == "" || (gender != "Male" && gender != "Female") || firstName == "" || lastName == "" || email == "" || password == "" {
		response["modification"]["error"] = "There is an empty field"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	err := users.UpdateUser(id, username, age, gender, firstName, lastName, email, password)
	if err != nil {
		response["modification"]["error"] = err.Error()

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return errors.New("error")
	}

	structure.Connected[username] = structure.Connected[modificationForm["currentUsername"]]
	delete(structure.Connected, modificationForm["currentUsername"])

	response["modification"]["success"] = structure.Html.User

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)

	return nil
}
