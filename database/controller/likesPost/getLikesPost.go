package likespost

import "forum_perso/database/initialisation"

/*
This function takes 1 string has argument:
  - the uuid of the post you want to add a like

The objective of this function is to get all row in the tab LikesPosts of the BDD.

The function gonna return:
  - an array of string (who contain the ids of the users who like the post)
  - an error
*/
func GetLikesPost(id string) ([]string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	datas, err := db.Query("SELECT `IdUser` FROM `LikesPosts` WHERE `IdPost`=?", id)
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
  - the uuid of the post you want to add a dislike

The objective of this function is to get all row in the tab DislikesPosts of the BDD.

The function gonna return:
  - an array of string (who contain the ids of the users who dislike the post)
  - an error
*/
func GetDislikesPost(id string) ([]string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	datas, err := db.Query("SELECT `IdUser` FROM `DislikesPosts` WHERE `IdPost`=?", id)
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
