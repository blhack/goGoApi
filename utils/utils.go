package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

var LastNewStory = 0

type Success struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

func CheckErr(err error) {

	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func GetAge(timestamp int) string {
	var ageString string
	t := time.Now()
	age := int(t.Unix()) - timestamp

	if age < 60 {
		ageString = fmt.Sprintf("%v seconds ago", age)
	} else if age >= 60 && age < 3600 {
		ageString = fmt.Sprintf("%v minutes ago", age/60)
	} else if age >= 3600 && age < 86400 {
		ageString = fmt.Sprintf("%v hours ago", age/3600)
	} else {
		ageString = fmt.Sprintf("%v days ago", age/86400)
	}

	return (ageString)
}

func GenerateUuid() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func SayHi() {
	fmt.Println("wuba luba DUB DUB")
}
