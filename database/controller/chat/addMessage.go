package chat

import (
	"forum_perso/database/initialisation"
	"time"

	"github.com/gofrs/uuid"
)

/*
This function takes 3 strings has arguments:
	- a message
	- the uuid of the people have send the notification  
	- the uuid of the people have receive the notification

The objective of this function is to create a new row in the tab Chat of the BDD.

The function gonna return:
	- an error
*/
func AddMessage(message, senderId, receverId string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `Chat`(`Id`, `Message`, `SenderId`, `ReceverId`, `Date`) VALUES(?,?,?,?,?)", uuid.String(), message, senderId, receverId, time.Now().Format("2006-01-02 15:04:05"))
	return err
}

/*
This function takes 2 string has argument:
	- the uuid of the people have send the notification  
	- the uuid of the people have receive the notification

The objective of this function is to create a new row in the tab Notification of the BDD.

The function gonna return:
	- an error
*/
func AddNotification(senderId, receverId string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `Notification`(`Id`, `Utilisator`, `NotifFrom`) VALUES(?,?,?)", uuid.String(), receverId, senderId)
	return err
}
