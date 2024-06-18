package structure

import (
	"html/template"

	"github.com/gorilla/websocket"
)

type IDContextKey string
type DataContextKey string

const IdCtx IDContextKey = "id"
const DataCtx DataContextKey = "data"

var (
	User Users
	Html Htmls
	Tpl  *template.Template

	Connected = map[string]*websocket.Conn{}
)

type Message struct {
	Id      string
	Sender  Users
	Recever Users
	Message string
	Date    string
}

type Htmls struct {
	User       Users
	Categories []Categorie
	Posts      []Post
	Comments   []Comment

	Error MyError
}

type Users struct {
	Id               string
	Username         string
	
	Age              string
	Gender           string
	FirstName        string
	LastName         string

	Email            string
	Password         string
	RegistrationDate string
	Role             int
	Connected        int

	ImagePath    string
	Notification int
}

type Categorie struct {
	Id   int
	Name string
}

type Post struct {
	Id                  string
	NameCreator         string
	CreatorImageProfile string
	NameCategories      []string
	Name                string
	Description         string
	CreationDate        string
	Likes               Like
	Dislikes            Like
	ImagePath           any
}

type Comment struct {
	Id           string
	IdPost       string
	Text         string
	CreationDate string
	ImagePath    any
	Likes        Like
	Dislikes     Like

	NameCreator      string
	CreatorImagePath string
}

type Like struct {
	Quantity int
	UserId   []string
}

type MyError struct {
	CreatPostError        string
	FilterError           string
	ModificationUserError string
	CreatCommentError     string
	LikeError             string
}
