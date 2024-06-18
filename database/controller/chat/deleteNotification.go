package chat

import "forum_perso/database/initialisation"

/*
This function takes 2 string has argument:
	- the uuid of the people have send the notification  
	- the uuid of the people have receive the notification

The objective of this function is to delete a of row in the tab Notification of the BDD.

The function gonna return:
	- an error
*/
func DeteleNotification(forId, fromId string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM `Notification` WHERE `Utilisator` = ? AND `NotifFrom` = ?", forId, fromId)
	return err
}