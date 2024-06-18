package users

import (
	"forum_perso/database/initialisation"
	"time"

	"github.com/gofrs/uuid"
)

/*
This function takes 7 strings has arguments:
	- the username of the new user
	- the age of the new user
	- the gender of the new user
	- the first name of the new user
	- the last name of the new user
	- the email of the new user
	- the password crypted of the new user

The objective of this function is to create a row in the tab Users of the BDD.

The function gonna return:
  - an error
*/
func AddUser(username, age, gender, firstName, lastName, email, password string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `Users`(`Id`, `Username`, `Age`, `Gender`, `FirstName`, `LastName`, `Email`, `Password`, `RegistrationDate`) VALUES(?,?,?,?,?,?,?,?,?)", uuid.String(), username, age, gender, firstName, lastName, email, password, time.Now().Format("2006-01-02 15:04:05"))
	return err
}
