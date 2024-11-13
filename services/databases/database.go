package databases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	UsersDB     []string
	PhoneDB     []string
	EmailsDB    []string
	NamesDB     []string
	LastNamesDB []string
)

func InitDB() {
	loadFile("services/databases/users.json", &UsersDB)
	loadFile("services/databases/names.json", &NamesDB)
	loadFile("services/databases/lastnames.json", &LastNamesDB)
	loadFile("services/databases/emails.json", &EmailsDB)
	loadFile("services/databases/phones.json", &PhoneDB)
}

func loadFile(file string, db *[]string) {
	f, err := os.Open(file)

	if err != nil {
		fmt.Println("[!] Error read data:", file)
		return
	}

	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal([]byte(byteValue), db)
}
