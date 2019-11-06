package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/things"
	"github.com/blhack/goGoApi/utils"
	"net/http"
)

func GetThings(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application-json")
	listId := r.FormValue("listId")

	var response []byte
	if listId != "" {
		things := things.GetThings(listId)
		if len(things) > 0 {
			response, _ = json.Marshal(things)
		} else {
			response = []byte("[]")
		}

	} else {
		response, _ = json.Marshal(utils.Success{false, "specify a ?listId=foo"})
	}

	fmt.Fprintf(w, string(response))

}
