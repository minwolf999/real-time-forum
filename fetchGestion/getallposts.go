package fetchgestion

import (
	"encoding/json"
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
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts.GetPosts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structure.Html)
}
