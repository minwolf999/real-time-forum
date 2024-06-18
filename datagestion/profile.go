package datagestion

import (
	"encoding/json"
	"fmt"
	"forum_perso/database/controller/comment"
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

The objective of this function is to manage the profile page.

The function gonna return:
	- an error
*/
func ProfileHandle(w http.ResponseWriter, profileForm map[string]string, conn *websocket.Conn) error {
	response := make(map[string]map[string]any)
	response["profile"] = make(map[string]any)

	if profileForm["filterType"] == "myPost" {
		id := profileForm["userId"]
		posts.GetPostsByUserId(id)
	}

	if profileForm["filterType"] == "myComment" {
		posts.GetPosts()

	triByComment:
		for i, v := range structure.Html.Posts {
			hasSendComment := false
			comment.GetCommentsByPostId(v.Id)
			for _, w := range structure.Html.Comments {
				if w.NameCreator == structure.Html.User.Username {
					hasSendComment = true
					break
				}
			}

			if hasSendComment {
				continue
			}

			if i < len(structure.Html.Posts)-1 {
				structure.Html.Posts = append(structure.Html.Posts[:i], structure.Html.Posts[i+1:]...)
				goto triByComment
			} else {
				structure.Html.Posts = structure.Html.Posts[:i]
				goto triByComment
			}
		}
	}

	if profileForm["filterType"] == "myLikePost" {
		posts.GetPosts()

	triByLikePost:
		for i, v := range structure.Html.Posts {
			has_like_or_dislike := false
			for _, w := range v.Likes.UserId {
				if w == structure.Html.User.Id {
					has_like_or_dislike = true
					break
				}
			}

			if !has_like_or_dislike {
				if i < len(structure.Html.Posts)-1 {
					structure.Html.Posts = append(structure.Html.Posts[:i], structure.Html.Posts[i+1:]...)
					goto triByLikePost
				} else {
					structure.Html.Posts = structure.Html.Posts[:i]
					goto triByLikePost
				}
			}
		}
	}

	if profileForm["filterType"] == "myDislikePost" {
		posts.GetPosts()

	triByDislikePost:
		for i, v := range structure.Html.Posts {
			has_like_or_dislike := false

			for _, w := range v.Dislikes.UserId {
				if w == structure.Html.User.Id {
					has_like_or_dislike = true
					break
				}
			}

			if !has_like_or_dislike {
				if i < len(structure.Html.Posts)-1 {
					structure.Html.Posts = append(structure.Html.Posts[:i], structure.Html.Posts[i+1:]...)
					goto triByDislikePost
				} else {
					structure.Html.Posts = structure.Html.Posts[:i]
					goto triByDislikePost
				}
			}
		}
	}

	response["profile"]["success"] = "success"

	formated, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	conn.WriteMessage(websocket.TextMessage, formated)

	return nil
}
