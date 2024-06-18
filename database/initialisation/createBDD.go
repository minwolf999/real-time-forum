package initialisation

/*
This function takes no argument

The objective of this function is to create the BDD and fill the Categorie tab.

The function gonna return:
  - an error
*/
func CreateBDD() error {
	db, err := OpenBDD()
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Users` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `Username` TEXT UNIQUE NOT NULL, `Age` TEXT NOT NULL, `Gender` TEXT NOT NULL, `FirstName` TEXT NOT NULL, `LastName` TEXT NOT NULL, `Email` TEXT UNIQUE NOT NULL, `Password` TEXT NOT NULL, `RegistrationDate` TEXT NOT NULL, `Role` INTEGER NOT NULL DEFAULT 0, `Connected` INTEGER NOT NULL DEFAULT 0, `ImagePath` TEXT NOT NULL DEFAULT '/template/image/imagesUser/default.png');")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Categories` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `Name` TEXT UNIQUE NOT NULL);")
	if err != nil {
		return err
	}

	var categorie = []string{"Beginner", "Intermediate", "Confirm", "Expert", "Public Project"}
	for i, v := range categorie {
		_, err = db.Exec("INSERT INTO Categories VALUES(?,?);", i, v)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("CREATE TABLE `Posts` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdCategories` TEXT NOT NULL, IdCreator TEXT NOT NULL REFERENCES`Users`(`Id`), `Name` TEXT UNIQUE NOT NULL, `Description` TEXT NOT NULL, `CreationDate` TEXT NOT NULL, `ImagePath` TEXT);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `LikesPosts` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdPost` TEXT NOT NULL REFERENCES `Posts`(`Id`), `IdUser` TEXT NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `DislikesPosts` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdPost` TEXT NOT NULL REFERENCES `Posts`(`Id`), `IdUser` TEXT NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Comments` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdPost` TEXT NOT NULL REFERENCES `Posts`(`Id`), `IdCreator` TEXT NOT NULL REFERENCES `Users`(`Id`), `Text` TEXT NOT NULL, `CreationDate` TEXT NOT NULL, `ImagePath` TEXT);")
	if err != nil {
		return err
	}
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `LikesComments` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdComment` TEXT NOT NULL REFERENCES `Comments`(`Id`), `IdUser` TEXT NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `DislikesComments` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `IdComment` TEXT NOT NULL REFERENCES `Comments`(`Id`), `IdUser` TEXT NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Chat` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `Message` TEXT NOT NULL, `SenderId` TEXT NOT NULL REFERENCES `Users`(`Id`), `ReceverId` TEXT NOT NULL REFERENCES `Users`(`Id`), `Date` TEXT NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Notification` (`Id` TEXT PRIMARY KEY UNIQUE NOT NULL, `Utilisator` TEXT NOT NULL REFERENCES `Users`(`Id`), `NotifFrom` TEXT NOT NULL REFERENCES `Users`(`Id`));")
	return err
}
