package likesComment

import "forum_perso/database/initialisation"

/*
This function takes 1 string has argument:
  - the uuid of the comment you want to get all this like

The objective of this function is to get all new row in the tab LikesComments of the BDD link to a comment.

The function gonna return:
	- an array of string who contain only the ids of the users has like the comment
	- an error
*/
func GetLikesComment(id string) ([]string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	datas, err := db.Query("SELECT `IdUser` FROM `LikesComments` WHERE `IdComment`=?", id)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var tab []string
	for datas.Next() {
		var i string
		err = datas.Scan(&i)
		if err != nil {
			return nil, err
		}

		tab = append(tab, i)
	}

	return tab, nil
}

/*
This function takes 1 string has argument:
  - the uuid of the comment you want to get all this dislike

The objective of this function is to get all new row in the tab DislikesComments of the BDD link to a comment.

The function gonna return:
	- an array of string who contain only the ids of the users has dislike the comment
	- an error
*/
func GetDislikesComment(id string) ([]string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	datas, err := db.Query("SELECT `IdUser` FROM `DislikesComments` WHERE `IdComment`=?", id)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var tab []string
	for datas.Next() {
		var i string
		err = datas.Scan(&i)
		if err != nil {
			return nil, err
		}

		tab = append(tab, i)
	}

	return tab, nil
}