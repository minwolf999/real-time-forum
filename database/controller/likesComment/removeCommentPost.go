package likesComment

import "forum_perso/database/initialisation"

/*
This function takes 2 string has arguments:
  - the uuid of the comment you want to remove a like
  - the uuid of the user you want to delete his like

The objective of this function is to remove a row in the tab LikesComments of the BDD.

The function gonna return:
  - an error
*/
func RemoveLikeComment(idComment, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM `LikesComments` WHERE `IdComment`=? AND `IdUser`=?", idComment, idUser)
	return err
}

/*
This function takes 2 string has arguments:
  - the uuid of the comment you want to remove a dislike
  - the uuid of the user you want to delete his dislike

The objective of this function is to remove a row in the tab DislikesComments of the BDD.

The function gonna return:
  - an error
*/
func RemoveDislikeComment(idComment, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM `DislikesComments` WHERE `IdComment`=? AND `IdUser`=?", idComment, idUser)
	return err
}
