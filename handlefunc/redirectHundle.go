package handlefunc

import "net/http"

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to redirect if the user try to access to a route who isn't create in the function main.

The function gonna return nothing
*/
func RedirectHundle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}