package dataprocess

import (
	. "DirectoryApi/utils"
	"encoding/json"
	"io/ioutil"
	"os"

	"DirectoryApi/models"
)

func Loadusers() []models.User {
	bytes, err := ioutil.ReadFile("json/users.json")
	CheckErr(err)
	var users []models.User

	json.Unmarshal(bytes, &users)

	return users
}
func Saveuser(a models.User) {
	var users []models.User = Loadusers()
	newUsers := append(users, a)
	bytes, err := json.Marshal(newUsers)
	CheckErr(err)

	f, err := os.OpenFile("json/users.json", os.O_WRONLY, 0600)
	CheckErr(err)

	defer f.Close()

	_, err = f.Write(bytes)
	CheckErr(err)

	return
}
