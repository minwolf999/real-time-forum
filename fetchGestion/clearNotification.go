package fetchgestion

import (
	"fmt"
	"forum_perso/database/controller/chat"
	"forum_perso/structure"
	"net/http"
	"strings"
)

/*
This function takes 2 arguments:
	- a Response Writer from the import net/http
	- a Request from the import net/http

The objective of this function is to manage the clear notification route.

The function gonna return nothing
*/
func ClearNotification(w http.ResponseWriter, r *http.Request) {
	datassConverted := fmt.Sprintf("%s", r.Context().Value(structure.DataCtx))
	datas := strings.Split(datassConverted, "\r")

	chat.DeteleNotification(datas[0], datas[len(datas)-1])
}
