package fetchgestion

import (
	"fmt"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	"net/http"
	"strings"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the disconnection route.

The function gonna return nothing
*/
func DisconnectHandle(w http.ResponseWriter, r *http.Request) {
	datassConverted := fmt.Sprintf("%s", r.Context().Value(structure.DataCtx))
	datas := strings.Split(datassConverted, "\r")

	delete(structure.Connected, datas[1])
	users.DisconnectUserByName(datas[1])
}
