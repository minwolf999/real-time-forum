package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/posts"
	"forum_perso/structure"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the filter page.

The function gonna return:
	- an error
*/
func FilterHandle(w http.ResponseWriter, filterForm map[string]string, conn *websocket.Conn) error {
	response := make(map[string]map[string]any)
	response["filter"] = make(map[string]any)

	if filterForm["type"] == "filterByName" {
		namePost := filterForm["value"]
		posts.GetPostsByName(namePost)

		if len(structure.Html.Posts) == 0 {
			response["filter"]["error"] = "There is no post named: " + namePost

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)
			return errors.New("error")
		}

		response["filter"]["success"] = "success"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return nil
	}

	posts.GetPosts()
	if filterForm["type"] == "filterByLikes" {
		for i := 0; i < len(structure.Html.Posts)-1; i++ {
			for y := len(structure.Html.Posts) - 1; y > i; y-- {
				if structure.Html.Posts[i].Likes.Quantity < structure.Html.Posts[y].Likes.Quantity {
					structure.Html.Posts[i], structure.Html.Posts[y] = structure.Html.Posts[y], structure.Html.Posts[i]
				}
			}
		}

		response["filter"]["success"] = "success"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return nil
	}

	if filterForm["type"] == "filterByDislikes" {
		for i := 0; i < len(structure.Html.Posts)-1; i++ {
			for y := len(structure.Html.Posts) - 1; y > i; y-- {
				if structure.Html.Posts[i].Likes.Quantity < structure.Html.Posts[y].Likes.Quantity {
					structure.Html.Posts[i], structure.Html.Posts[y] = structure.Html.Posts[y], structure.Html.Posts[i]
				}
			}
		}

		response["filter"]["success"] = "success"

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return nil
	}

	return nil
}
