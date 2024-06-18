package comment

import (
	"forum_perso/database/initialisation"
	"time"

	"github.com/gofrs/uuid"
)

/*
This function takes 4 strings has arguments:
  - the uuid of the post for link the new comment at this post
  - the uuid of the people who has send the comment
  - the message of the comment

The objective of this function is to create a new row in the tab Comments of the BDD.

The function gonna return:
	- an error
*/
func CreatNewComment(idPost, idCreator, text, imagePath string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	uuid, _ := uuid.NewV7()

	_, err = db.Exec("INSERT INTO `Comments`(`Id`, `IdPost`, `IdCreator`, `Text`, `CreationDate`, `ImagePath`) VALUES(?,?,?,?,?,?)", uuid.String(), idPost, idCreator, text, time.Now().Format("2006-01-02 15:04:05"), imagePath)
	return err
}
