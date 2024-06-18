package categories

import (
	"forum_perso/database/initialisation"
	"forum_perso/structure"
)

/*
This function takes no argument:

The objective of this function is to get all row in the tab Categories of the BDD.

The function gonna return:
	- an error
*/
func GetCategories() error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Categories`")
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Categories = []structure.Categorie{}

	for datas.Next() {
		var tmp structure.Categorie

		err = datas.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			return err
		}

		structure.Html.Categories = append(structure.Html.Categories, tmp)
	}

	return nil
}

/*
This function takes 1 int has argument:
	- an id

The objective of this function is to get a categorie name by his id in the tab Notification of the BDD.

The function gonna return:
	- a string who is the name of a categorie
	- an error
*/
func GetCategoriesById(id int) (string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return "", err
	}

	datas, err := db.Query("SELECT `Name` FROM `Categories` WHERE `Id`=?", id)
	if err != nil {
		return "", err
	}
	defer datas.Close()

	for datas.Next() {
		var res string

		err = datas.Scan(&res)
		return res, err
	}

	return "", nil
}
