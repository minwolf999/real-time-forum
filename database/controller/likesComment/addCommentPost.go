package likesComment

import (
	"forum_perso/database/initialisation"
	"forum_perso/structure"

	"github.com/gofrs/uuid"
)

/*
This function takes 2 strings has arguments:
  - the uuid of the comment who gonna be liked
  - the uuid of the user who want to liked a comment

The objective of this function is to add a new row in the tab LikesComments of the BDD.

The function gonna return:
	- an error
*/
func AddLikeComment(idComment, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	for _, v := range structure.Html.Comments {
		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idComment {
				return RemoveLikeComment(idComment, idUser)
			}
		}

		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idComment {
				RemoveDislikeComment(idComment, idUser)
			}
		}
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `LikesComments`(`Id`, `IdComment`, `IdUser`) VALUES(?,?,?)", uuid.String(), idComment, idUser)
	return err
}

/*
This function takes 2 strings has arguments:
  - the uuid of the comment who gonna be disliked
  - the uuid of the user who want to disliked a comment

The objective of this function is to add a new row in the tab DislikesComments of the BDD.

The function gonna return:
	- an error
*/
func AddDislikeComment(idComment, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	for _, v := range structure.Html.Comments {
		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idComment {
				return RemoveDislikeComment(idComment, idUser)
			}
		}

		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idComment {
				RemoveLikeComment(idComment, idUser)
			}
		}
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `DislikesComments`(`Id`, `IdComment`, `IdUser`) VALUES(?,?,?)", uuid.String(), idComment, idUser)
	return err
}
