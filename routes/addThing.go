package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/flexIms-go/auth"
	"github.com/blhack/flexIms-go/things"
	"github.com/blhack/flexIms-go/utils"
	"net/http"
)

func AddThing(w http.ResponseWriter, r *http.Request) {

	userName := auth.WhoAmI(r).Status
	this := utils.Success{false, ""}

	if auth.WhoAmI(r).Success == true {
		var _thing things.Thing
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&_thing)
		utils.CheckErr(err)
		fmt.Printf("%+v\n",_thing)
		this = things.AddThing(_thing.Title, userName, "",_thing.Attributes)
	} else {
		this.Success = false
		this.Status = "not authenticated"
	}
	_this, _ := json.Marshal(this)
	fmt.Fprintf(w, string(_this))
}
