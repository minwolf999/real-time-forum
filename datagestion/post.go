package datagestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum_perso/database/controller/comment"
	"forum_perso/database/controller/likesComment"
	likespost "forum_perso/database/controller/likesPost"
	"forum_perso/database/controller/posts"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
This function takes 3 arguments:
	- a Response Writer from the import net/http
	- a map who contain the datas send by the websocket
	- a connection to an open websocket

The objective of this function is to manage the post detail page.

The function gonna return:
	- an error
*/
func PostHandle(w http.ResponseWriter, commentForm map[string]string, conn *websocket.Conn) error {
	response := make(map[string]map[string]any)
	response["comment"] = make(map[string]any)

	if commentForm["type"] == "likePost" {
		PostId := commentForm["id"]
		userId := commentForm["userId"]

		if commentForm["value"] == "like" {
			likespost.AddLikePost(PostId, userId)
		} else if commentForm["value"] == "dislike" {
			likespost.AddDislikePost(PostId, userId)
		}

		posts.GetPostsById(PostId)
		comment.GetCommentsByPostId(PostId)

		response["comment"]["success"] = "success|" + PostId

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
		return nil
	}

	if commentForm["type"] == "LikeComment" {
		commentId := commentForm["id"]
		PostId := commentForm["postId"]
		userId := commentForm["userId"]

		if commentForm["value"] == "like" {
			likesComment.AddLikeComment(commentId, userId)
		} else if commentForm["value"] == "dislike" {
			likesComment.AddDislikeComment(commentId, userId)
		}

		posts.GetPostsById(PostId)
		comment.GetCommentsByPostId(PostId)

		response["comment"]["success"] = "success|" + PostId

		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)
	}

	if commentForm["type"] == "sendComment" {
		message := commentForm["text"]
		PostId := commentForm["id"]
		UserId := commentForm["userId"]

		if message == "" {
			response["profile"]["error"] = "The comment you tried to create contains empty text"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)

			return errors.New("error")
		}

		if !verificationfunction.IsValidMessage(message) && message != "" {
			response["profile"]["error"] = "The comment you tried to create contains empty text"

			formated, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
			}

			conn.WriteMessage(websocket.TextMessage, formated)

			return errors.New("error")
		}

		comment.CreatNewComment(PostId, UserId, message, "")

		posts.GetPostsById(PostId)
		comment.GetCommentsByPostId(PostId)

		response["comment"]["success"] = "success|" + PostId
		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return nil
	}

	if commentForm["type"] == "filtre" {
		PostId := commentForm["id"]

		posts.GetPostsById(PostId)
		comment.GetCommentsByPostId(PostId)

		if commentForm["value"] == "like" {
			for i := 0; i < len(structure.Html.Comments)-1; i++ {
				for y := len(structure.Html.Comments) - 1; y > i; y-- {
					if structure.Html.Comments[i].Likes.Quantity < structure.Html.Comments[y].Likes.Quantity {
						structure.Html.Comments[i], structure.Html.Comments[y] = structure.Html.Comments[y], structure.Html.Comments[i]
					}
				}
			}
		} else if commentForm["value"] == "dislike" {
			for i := 0; i < len(structure.Html.Comments)-1; i++ {
				for y := len(structure.Html.Comments) - 1; y > i; y-- {
					if structure.Html.Comments[i].Dislikes.Quantity < structure.Html.Comments[y].Dislikes.Quantity {
						structure.Html.Comments[i], structure.Html.Comments[y] = structure.Html.Comments[y], structure.Html.Comments[i]
					}
				}
			}
		}
		
		response["comment"]["success"] = "filtered|" + PostId
		formated, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, formated)

		return nil
	}

	return nil
}
