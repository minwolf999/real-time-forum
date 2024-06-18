package posts

import (
	"forum_perso/database/initialisation"
	"time"

	"github.com/gofrs/uuid"
)

/*
This function takes 5 strings has arguments:
  - the string of the ids of the categories who as be joined by a "|"
  - the uuid of the user who create the post
  - the title of the post
  - the description of the post
  - the image path if the posthas an image (if there is not image give an empty string)

The objective of this function is to create a new row in the tab Posts of the BDD.

The function gonna return:
  - an error
*/
func CreateNewPost(idCategories, idCreator, title, description, imagePath string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `Posts`(`Id`, `IdCategories`, `IdCreator`, `Name`, `Description`, `CreationDate`, `ImagePath`) VALUES(?,?,?,?,?,?,?)", uuid.String(), idCategories, idCreator, title, description, time.Now().Format("2006-01-02 15:04:05"), imagePath)
	return err
}
