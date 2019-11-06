package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/auth"
	"net/http"
)

func WhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application-json")

	userName := auth.WhoAmI(r)
	status, _ := json.Marshal(userName)
	fmt.Fprintf(w, string(status))

}
