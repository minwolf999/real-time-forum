package users

import (
	"errors"
	"fmt"
	"forum_perso/database/initialisation"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"

	"golang.org/x/crypto/bcrypt"
)

/*
This function takes 8 strings has arguments:
  - an id
  - an username
  - an age
  - a gender
  - a first name
  - a last name
  - an email
  - a password

The objective of this function is to update the row in the tab Users of the BDD where the id is equal to the first argument.

The function gonna return:
  - an error
*/
func UpdateUser(id, username, age, gender, firstName, lastName, email, password string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	var request string
	if verificationfunction.PasswordVerif(password) {
		cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

		request = fmt.Sprintf("UPDATE `Users` SET `Username`='%s', `Age`='%s', `Gender`='%s', `FirstName`='%s', `LastName`='%s', `Email`='%s', `Password`='%s' WHERE `Id`=?", username, age, gender, firstName, lastName, email, cryptPassword)
	} else {
		return errors.New("incorrect password ! The password must contain 8 characters, 1 uppercase letter, 1 special character, 1 number")
	}

	_, err = db.Exec(request, id)
	if err != nil {
		return err
	}

	err = GetUserByEmail(email)
	structure.Html.User = structure.User
	return err
}

/*
This function takes 1 string has argument:
  - an id

The objective of this function is to update to 1 the column Connected in the row who is in the tab Users of the BDD where the id is equal to the argument.

The function gonna return:
  - an error
*/
func ConnectUser(id string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE `Users` SET `Connected`=1 WHERE `Id`=?", id)
	if err != nil {
		structure.Html.User = structure.Users{}
		return err
	}

	return nil
}

/*
This function takes 1 string has argument:
  - an id

The objective of this function is to update to 0 the column Connected in the row who is in the tab Users of the BDD where the id is equal to the argument.

The function gonna return:
  - an error
*/
func DisconnectUser(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE `Users` SET `Connected`=? WHERE `Id`=?", 0, id)
	if err != nil {
		return err
	}

	structure.Html.User = structure.Users{}
	return nil
}

/*
This function takes 1 string has argument:
  - an username

The objective of this function is to update to 0 the column Connected in the row who is in the tab Users of the BDD where the username is equal to the argument.

The function gonna return:
  - an error
*/
func DisconnectUserByName(username string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE `Users` SET `Connected`=? WHERE `Username`=?", 0, username)
	if err != nil {
		return err
	}

	structure.Html.User = structure.Users{}
	return nil
}
