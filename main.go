package main

import (
	"fmt"
	"forum_perso/database/controller/categories"
	"forum_perso/database/controller/posts"
	"forum_perso/database/initialisation"
	fetchgestion "forum_perso/fetchGestion"
	"forum_perso/handlefunc"
	"forum_perso/structure"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// os.Remove("structure/database.db")
	// os.RemoveAll("template/image/imageComments")
	// os.Mkdir("template/image/imageComments", os.ModePerm)

	initialisation.CreateBDD()
	initialisation.DisconnectAllUsers()
	categories.GetCategories()
	posts.GetPosts()

	structure.User.Username = "Visitor"

	// Parsing all html files
	structure.Tpl = template.Must(template.New("").ParseGlob("template/html/**/*.html"))

	// Listen statics files
	fs := http.FileServer(http.Dir("template"))
	http.Handle("/template/", http.StripPrefix("/template", fs))
}

func main() {
	fmt.Print("\033[96m")
	fmt.Println("\033[96mServer started at: http://localhost:8080\033[0m")
	fmt.Print("\033[0m")

	http.HandleFunc("/", handlefunc.RedirectHundle)
	http.HandleFunc("/login", handlefunc.Handle)

	http.HandleFunc("/getcategories05842365165", handlefunc.FetchMiddleware(fetchgestion.GetCategories))
	http.HandleFunc("/getallposts05842365165", handlefunc.FetchMiddleware(fetchgestion.GetAllPosts))
	http.HandleFunc("/getposts05842365165", handlefunc.FetchMiddleware(fetchgestion.GetUserPost))
	http.HandleFunc("/getcomments05842365165", handlefunc.FetchMiddleware(fetchgestion.GetComments))
	http.HandleFunc("/getusers05842365165", handlefunc.FetchMiddleware(fetchgestion.GetUsers))
	http.HandleFunc("/getmessages05842365165", handlefunc.FetchMiddleware(fetchgestion.GetMessages))
	http.HandleFunc("/getnotificationQuantity05842365165", handlefunc.FetchMiddleware(fetchgestion.GetNotificationQuantity))

	http.HandleFunc("/clearNotification05842365165", handlefunc.FetchMiddleware(fetchgestion.ClearNotification))
	http.HandleFunc("/disconnect05842365165", handlefunc.FetchMiddleware(fetchgestion.DisconnectHandle))

	http.ListenAndServe(":8080", nil)
}

