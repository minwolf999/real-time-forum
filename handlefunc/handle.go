package handlefunc

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/posts"
	"forum_perso/structure"
	"forum_perso/websocketGestion"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to load the html file and start the infinite loop to keep the websocket open.

The function gonna return nothing
*/
func Handle(w http.ResponseWriter, r *http.Request) {
	structure.Html.User = structure.Users{}

	if r.Header.Get("Upgrade") != "websocket" {
		posts.GetPosts()
		html, err := json.Marshal(structure.Html)
		if err != nil {
			fmt.Println(err)
			return
		}

		structure.Tpl.ExecuteTemplate(w, "login.html", string(html))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		err = websocketGestion.WebsocketGestion(w, conn)
		if err != nil {
			break
		}
	}
}
