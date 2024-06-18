package posts

import (
	"fmt"
	"forum_perso/database/controller/categories"
	likespost "forum_perso/database/controller/likesPost"
	"forum_perso/database/controller/users"
	"forum_perso/database/initialisation"
	"forum_perso/structure"
	verificationfunction "forum_perso/verificationFunction"
	"strconv"
	"strings"
)

/*
This function takes no argument

The objective of this function is to get all row in the tab Posts of the BDD and stock them in the var Html.Posts in the structure package.

The function gonna return:
  - an error
*/
func GetPosts() error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Posts`")
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator string

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

/*
This function takes an int has argument:
	- the int is the id of the categorie you want to get all post who is link with

The objective of this function is to get all row in the tab Posts of the BDD when the post is link with the categorie's id passed has argument and stock them in the var Html.Posts in the structure package.

The function gonna return:
  - an error
*/
func GetPostsByCategorieId(idCategorie int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Posts`")
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator string

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		if verificationfunction.TabNotContain(idCategoriesTabString, strconv.Itoa(idCategorie)) {
			continue
		}

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

/*
This function takes an string has argument:
	- the string is the title of the post you want to get

The objective of this function is to get all row in the tab Posts of the BDD when the post's title is equal to the string passed has argument and stock them in the var Html.Posts in the structure package.

The function gonna return:
  - an error
*/
func GetPostsByName(s string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	s += "%"
	datas, err := db.Query("SELECT * FROM `Posts` WHERE `Name` LIKE ?", s)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator string

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

/*
This function takes an string has argument:
	- the string is the uuid of the user you want to get the post he has write

The objective of this function is to get all row in the tab Posts of the BDD when the post is link with the user's id passed has argument and stock them in the var Html.Posts in the structure package.

The function gonna return:
  - an error
*/
func GetPostsByUserId(id string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Posts` WHERE `IdCreator`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator string

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

/*
This function takes an string has argument:
	- the string is the uuid of the post you want to get

The objective of this function is to get a row in the tab Posts of the BDD when the post's id is equal to the string passed has parameter and stock them in the var Html.Posts in the structure package.

The function gonna return:
  - an error
*/
func GetPostsById(id string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}

	datas, err := db.Query("SELECT * FROM `Posts` WHERE `Id`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator string

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}
