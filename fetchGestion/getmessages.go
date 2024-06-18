package fetchgestion

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/chat"
	"forum_perso/structure"
	"net/http"
	"strconv"
	"strings"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the route for get all the post.

The function gonna return nothing normally
The function gonna return by the websocket the array of Message return by the GetMessages function
*/
func GetMessages(w http.ResponseWriter, r *http.Request) {
	datassConverted := fmt.Sprintf("%s", r.Context().Value(structure.DataCtx))
	datass := strings.Split(datassConverted, "\r")

	userId := datass[0]
	interlocutorId := datass[8]
	maxMessage, _ := strconv.Atoi(datass[9])

	datas, err := chat.GetMessages(userId, interlocutorId, maxMessage)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}
