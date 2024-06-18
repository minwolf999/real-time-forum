package websocketGestion

import (
	"encoding/json"
	"forum_perso/datagestion"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the websocket transfert datas.

The function gonna return:
	- an error
*/
func WebsocketGestion(w http.ResponseWriter, conn *websocket.Conn) error {
	_, tempMessage, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	tab := make(map[string]map[string]string)
	json.Unmarshal(tempMessage, &tab)

	loginForm, isOk := tab["login"]
	if isOk {
		datagestion.LogInHandle(w, loginForm, conn)
	}

	registerForm, isOk := tab["register"]
	if isOk {
		datagestion.RegisterHandle(w, registerForm, conn)
	}

	homeForm, isOk := tab["home"]
	if isOk {
		datagestion.HomeHandle(w, homeForm, conn)
	}

	filterForm, isOk := tab["filter"]
	if isOk {
		datagestion.FilterHandle(w, filterForm, conn)
	}

	profilForm, isOk := tab["profile"]
	if isOk {
		datagestion.ProfileHandle(w, profilForm, conn)
	}

	modificationForm, isOk := tab["modification"]
	if isOk {
		datagestion.ModificationUser(w, modificationForm, conn)
	}

	commentForm, isOk := tab["comment"]
	if isOk {
		datagestion.PostHandle(w, commentForm, conn)
	}

	chatForm, isOk := tab["chatWS"]
	if isOk {
		datagestion.ChatHandle(w, chatForm, conn)
	}

	return nil
}
