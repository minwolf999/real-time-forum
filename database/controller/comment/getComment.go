package comment

import (
	"forum_perso/database/controller/likesComment"
	"forum_perso/database/controller/users"
	"forum_perso/database/initialisation"
	"forum_perso/structure"
)

/*
This function takes 1 string has argument:
  - the uuid of the post link to all the comments you want to get

The objective of this function is to get all the row in the tab Comments of the BDD for a post.

The function gonna return:
	- an error
*/
func GetCommentsByPostId(id string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Comments` WHERE `IdPost`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Comments = []structure.Comment{}
	for datas.Next() {
		var tmp structure.Comment
		var idCreator string

		err = datas.Scan(&tmp.Id, &tmp.IdPost, &idCreator, &tmp.Text, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, creatorImage, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImagePath = creatorImage

		tmp.Likes.UserId, _ = likesComment.GetLikesComment(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likesComment.GetDislikesComment(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		structure.Html.Comments = append(structure.Html.Comments, tmp)
	}

	return nil
}
