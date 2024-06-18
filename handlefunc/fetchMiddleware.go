package handlefunc

import (
	"context"
	"encoding/json"
	"forum_perso/database/controller/users"
	"forum_perso/structure"
	"io"
	"net/http"
	"strings"
)

/*
This function takes 1 arguments:
	- a function who managed a page

The objective of this function is to verify if the user who make a request is well connected.

The function gonna return a function
*/
func FetchMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var datas []string
		json.Unmarshal(body, &datas)

		if len(datas) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized 1")
			return
		}

		id := datas[0]

		data, err := users.GetUserDatasById(id)
		_, isConnected := structure.Connected[data.Username]
		if err != nil || !isConnected || data.Username != datas[1] || data.Age != datas[2] || data.Gender != datas[3] || data.FirstName != datas[4] || data.LastName != datas[5] || data.Email != datas[6] || data.Password != datas[7] {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized 2")

			return
		}

		dataConverted := strings.Join(datas, "\r")

		ctx := context.WithValue(r.Context(), structure.IdCtx, id)
		newReq := r.WithContext(ctx)

		ctx2 := context.WithValue(newReq.Context(), structure.DataCtx, dataConverted)
		newReq2 := newReq.WithContext(ctx2)

		next.ServeHTTP(w, newReq2)
	})
}
