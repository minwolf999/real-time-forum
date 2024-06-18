package fetchgestion

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/posts"
	"forum_perso/structure"
	"net/http"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the route for get all the post.

The function gonna return nothing normally
The function gonna return by the websocket the Html variable from the package structure
*/
func GetUserPost(w http.ResponseWriter, r *http.Request) {
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))

	err := posts.GetPostsByUserId(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structure.Html)
}
