package fetchgestion

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/comment"
	"forum_perso/database/controller/posts"
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
The function gonna return by the websocket the Html variable from the package structure
*/
func GetComments(w http.ResponseWriter, r *http.Request) {
	datasConverted := fmt.Sprintf("%s", r.Context().Value(structure.DataCtx))
	datas := strings.Split(datasConverted, "\r")

	filtered, err := strconv.Atoi(datas[9])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	postId := datas[8]

	if filtered == 0 {
		err := posts.GetPostsById(postId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}

		err = comment.GetCommentsByPostId(postId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structure.Html)
}
