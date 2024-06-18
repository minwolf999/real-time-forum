package initialisation

import "forum_perso/structure"

/*
This function takes no argument

The objective of this function is to update to 0 all column Connected in the row who is in the tab Users of the BDD.

The function gonna return:
  - an error
*/
func DisconnectAllUsers() error {
	db, err := OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE `Users` SET `Connected`=?", 0)
	if err != nil {
		return err
	}

	structure.Html.User = structure.Users{}
	return nil
}
