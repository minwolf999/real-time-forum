package users

import (
	"forum_perso/database/initialisation"
	"forum_perso/structure"
)

/*
This function takes 2 strings has arguments:
  - the email of the new user
  - the password crypted of the new user

The objective of this function is to get all row in the tab Users of the BDD where the email and the password are equal to the arguments and stock them in the var User of the package structure.

The function gonna return:
  - an error
*/
func GetUser(email, password string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Email`=? AND `Password`=?", email, password)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&structure.User.Id, &structure.User.Username, &structure.User.Age, &structure.User.Gender, &structure.User.FirstName, &structure.User.LastName, &structure.User.Email, &structure.User.Password, &structure.User.RegistrationDate, &structure.User.Role, &structure.User.Connected, &structure.User.ImagePath)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
This function takes 1 string has argument:
  - an uuid

The objective of this function is to get a row in the tab Users of the BDD where the id is equal to the argument.

The function gonna return:
	- a string who contain the username of the user found 
	- a string who contain the image Path of the user found
	- an error
*/
func GetUserById(id string) (string, string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return "", "", err
	}

	datas, err := db.Query("SELECT `Username`, `ImagePath` FROM `Users` WHERE `Id`=?", id)
	if err != nil {
		return "", "", err
	}
	defer datas.Close()

	for datas.Next() {
		var name, imagePath string
		err = datas.Scan(&name, &imagePath)
		return name, imagePath, err
	}
	return "", "", nil
}

/*
This function takes 1 string has argument:
  - an uuid

The objective of this function is to get a row in the tab Users of the BDD where the id is equal to the argument.

The function gonna return:
	- a structure Users who contain the informations of the user found
	- an error
*/
func GetUserDatasById(id string) (structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return structure.Users{}, err
	}

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Id`=?", id)
	if err != nil {
		return structure.Users{}, err
	}
	defer datas.Close()

	for datas.Next() {
		var user structure.Users
		err = datas.Scan(&user.Id, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath)
		if err != nil {
			return structure.Users{}, err
		}

		user.Notification, err = GetNotificationQuantityBetween2Users(id, user.Id)
		return user, err
	}
	return structure.Users{}, nil
}

/*
This function takes 1 string has argument:
  - an email

The objective of this function is to get a row in the tab Users of the BDD where the email is equal to the argument and stock them in the var User of the package structure.

The function gonna return:
	- an error
*/
func GetUserByEmail(email string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Email`=?", email)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&structure.User.Id, &structure.User.Username, &structure.User.Age, &structure.User.Gender, &structure.User.FirstName, &structure.User.LastName, &structure.User.Email, &structure.User.Password, &structure.User.RegistrationDate, &structure.User.Role, &structure.User.Connected, &structure.User.ImagePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUserByEmailOrUsername(email, username string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Email`=? OR `Username`=?", email, username)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&structure.User.Id, &structure.User.Username, &structure.User.Age, &structure.User.Gender, &structure.User.FirstName, &structure.User.LastName, &structure.User.Email, &structure.User.Password, &structure.User.RegistrationDate, &structure.User.Role, &structure.User.Connected, &structure.User.ImagePath)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
This function takes 1 string has argument:
  - an uuid

The objective of this function is to get all row in the tab Users of the BDD where the id is equal to the argument.

The function gonna return:
	- an array of Users (who contain all user's datas found)
	- an error
*/
func GetAllUsers(id string) ([]structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	var AllUser []structure.Users
	datas, err := db.Query("SELECT * FROM `Users` WHERE `Id` != ?", id)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	for datas.Next() {
		var User structure.Users
		err = datas.Scan(&User.Id, &User.Username, &User.Age, &User.Gender, &User.FirstName, &User.LastName, &User.Email, &User.Password, &User.RegistrationDate, &User.Role, &User.Connected, &User.ImagePath)
		if err != nil {
			return nil, err
		}

		User.Notification, err = GetNotificationQuantityBetween2Users(id, User.Id)
		if err != nil {
			return nil, err
		}

		AllUser = append(AllUser, User)
	}

	return AllUser, nil
}

/*
This function takes 2 strings has arguments:
  - two uuid

The objective of this function is to get the quantity of row in the tab Notification of the BDD where the Utilisator and the NotifForm are equal to the arguments.

The function gonna return:
	- a int
	- an error
*/
func GetNotificationQuantityBetween2Users(UtilisatorId, otherId string) (int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return 0, err
	}

	datas, err := db.Query("SELECT COUNT(*) FROM `Notification` WHERE `Utilisator` = ? AND `NotifFrom` = ?", UtilisatorId, otherId)
	if err != nil {
		return 0, err
	}
	defer datas.Close()

	for datas.Next() {
		res := 0

		err = datas.Scan(&res)

		return res, err
	}

	return 0, nil
}
