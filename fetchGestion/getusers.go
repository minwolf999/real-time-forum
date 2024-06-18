package fetchgestion

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/chat"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	"net/http"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the route for get all the post.

The function gonna return nothing normally
The function gonna return by the websocket the array of Users
*/
func GetUsers(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))

	allUsers, err := users.GetAllUsers(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	for i := 0; i < len(allUsers)-1; i++ {
		for y := len(allUsers) - 1; y > i; y-- {
			if allUsers[i].Username > allUsers[y].Username {
				allUsers[i], allUsers[y] = allUsers[y], allUsers[i]
			}
		}
	}

	allMessages, err := chat.GetAllMessages(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	for i, y := 0, len(allMessages)-1; i < y; i, y = i+1, y-1 {
		allMessages[i], allMessages[y] = allMessages[y], allMessages[i]
	}

	var res []structure.Users

	res = append(res, allMessages...)
	res = append(res, allUsers...)

retry:
	for i := 0; i < len(res)-1; i++ {
		for y := len(res) - 1; y > i; y-- {
			if res[i].Username == res[y].Username {
				res = append(res[:y], res[y+1:]...)
				goto retry
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
