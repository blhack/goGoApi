package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/things"
	"github.com/blhack/goGoApi/utils"
	"github.com/blhack/goGoApi/auth"
	"net/http"
)

func GetThings(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application-json")
	userName := auth.WhoAmI(r).Status

	var response []byte
	if userName != "" {
		things := things.GetThings(userName)
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
