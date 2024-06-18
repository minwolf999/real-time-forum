package likespost

import "forum_perso/database/initialisation"

/*
This function takes 2 string has arguments:
  - the uuid of the post you want to remove a like
  - the uuid of the user you want to delete his like

The objective of this function is to remove a row in the tab LikesPosts of the BDD.

The function gonna return:
  - an error
*/
func RemoveLikePost(idPost, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM `LikesPosts` WHERE `IdPost`=? AND `IdUser`=?", idPost, idUser)
	return err
}

/*
This function takes 2 string has arguments:
  - the uuid of the post you want to remove a dislike
  - the uuid of the user you want to delete his dislike

The objective of this function is to remove a row in the tab DislikesPosts of the BDD.

The function gonna return:
  - an error
*/
func RemoveDislikePost(idPost, idUser string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM `DislikesPosts` WHERE `IdPost`=? AND `IdUser`=?", idPost, idUser)
	return err
}
