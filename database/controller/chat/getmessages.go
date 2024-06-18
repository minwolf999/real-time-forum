package chat

import (
	"errors"
	"forum_perso/database/controller/users"
	"forum_perso/database/initialisation"
	"forum_perso/structure"
)

/*
This function takes 2 strings and 1 int has arguments:
  - the uuid of the people who send the notification
  - the uuid of the people gonna receive the notification
  - the quantity of message is always to 10 message and the int has to objectives to deplace the selection 10 by 10 in the BDD

The objective of this function is to get a maximum off 10 row in the tab Chat of the BDD.

The function gonna return:
	- an array of the structure Message
	- an error
*/
func GetMessages(id1, id2 string, maxMessage int) ([]structure.Message, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	if maxMessage%10 != 0 {
		return nil, errors.New("the third parameter isn't a multiple of 10")
	}

	datas, err := db.Query("SELECT * FROM `Chat` WHERE (`SenderId` = ? AND `ReceverId` = ?) OR (`ReceverId` = ? AND `SenderId` = ?) ORDER BY `Date` DESC LIMIT ? OFFSET ?", id1, id2, id1, id2, maxMessage, maxMessage-10)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var res []structure.Message
	for datas.Next() {
		var data structure.Message
		var userId1, userId2 string
		err = datas.Scan(&data.Id, &data.Message, &userId1, &userId2, &data.Date)
		if err != nil {
			return nil, err
		}

		user1, err := users.GetUserDatasById(userId1)
		if err != nil {
			return nil, err
		}

		data.Sender = user1

		user2, err := users.GetUserDatasById(userId2)
		if err != nil {
			return nil, err
		}

		data.Recever = user2

		res = append(res, data)
	}

	return res, nil
}

/*
This function takes 1 string has argument:
  - the uuid of the user you want to get all this message

The objective of this function is to get all the message of a user in the tab Chat of the BDD.

The function gonna return:
	- an array of the structure Message
	- an error
*/
func GetAllMessages(userId string) ([]structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}

	datas, err := db.Query("SELECT `SenderId`, `ReceverId` FROM `Chat` WHERE `SenderId` = ? OR `ReceverId` = ?", userId, userId)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var res []structure.Users
	for datas.Next() {
		var id1, id2 string
		var tmp structure.Users
		datas.Scan(&id1, &id2)

		if id1 == userId {
			tmp, err = users.GetUserDatasById(id2)
			if err != nil {
				return nil, err
			}
		} else {
			tmp, err = users.GetUserDatasById(id1)
			if err != nil {
				return nil, err
			}

		}
		tmp.Notification, err = users.GetNotificationQuantityBetween2Users(userId, tmp.Id)
		if err != nil {
			return nil, err
		}

		res = append(res, tmp)
	}

	return res, nil
}

/*
This function takes 1 string has argument:
  - the uuid of the people gonna receive the notification

The objective of this function is to get the quantity of row in the tab Notification of the BDD for the user link to the uuid give as argument.

The function gonna return:
	- an int (the quantity of notification)
	- an error
*/
func GetNotificationQuantity(userId string) (int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return 0, err
	}

	datas, err := db.Query("SELECT COUNT(*) FROM `Notification` WHERE `Utilisator` = ?", userId)
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
