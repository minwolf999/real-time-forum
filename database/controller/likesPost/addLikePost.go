package likespost

import (
	"forum_perso/database/initialisation"
	"forum_perso/structure"

	"github.com/gofrs/uuid"
)

/*
This function takes 2 string has arguments:
  - the uuid of the post you want to add a like
  - the uuid of the user who want to like 

The objective of this function is to add a row in the tab LikesPosts of the BDD.

The function gonna return:
  - an error
*/
func AddLikePost(idPost, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	for _, v := range structure.Html.Posts {
		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idPost {
				return RemoveLikePost(idPost, idUser)
			}
		}

		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idPost {
				RemoveDislikePost(idPost, idUser)
			}
		}
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `LikesPosts`(`Id`, `IdPost`, `IdUser`) VALUES(?,?,?)", uuid.String(), idPost, idUser)
	return err
}

/*
This function takes 2 string has arguments:
  - the uuid of the post you want to add a dislike
  - the uuid of the user who want to dislike

The objective of this function is to add a row in the tab DislikesPosts of the BDD.

The function gonna return:
  - an error
*/
func AddDislikePost(idPost, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	for _, v := range structure.Html.Posts {
		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idPost {
				return RemoveDislikePost(idPost, idUser)
			}
		}

		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idPost {
				RemoveLikePost(idPost, idUser)
			}
		}
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `DislikesPosts`(`Id`, `IdPost`, `IdUser`) VALUES(?,?,?)", uuid.String(), idPost, idUser)
	return err
}
