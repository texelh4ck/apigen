package random

import (
	"apigen/services/databases"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Generate names
func rName() string {
	name := databases.NamesDB[rand.Intn(len(databases.NamesDB))]
	lastname := databases.LastNamesDB[rand.Intn(len(databases.LastNamesDB))]
	return fmt.Sprintf("%s %s", name, lastname)
}

// Generate phone number
func rPhone() string {
	phone := databases.PhoneDB[rand.Intn(len(databases.PhoneDB))]
	phone += " "
	for i := 0; i < 8; i++ {
		phone += strconv.Itoa(rand.Intn(10))
	}
	return phone
}

// Generate email
func rEmail() string {
	u, n := databases.UsersDB[rand.Intn(len(databases.UsersDB))], databases.EmailsDB[rand.Intn(len(databases.EmailsDB))]
	return fmt.Sprintf("%s@%s", u, n)
}

// Generate username
func rUsername() string {
	return databases.UsersDB[rand.Intn(len(databases.UsersDB))]
}

// Generate score
func rScore() float64 {
	return math.Round(rand.Float64() * 5)
}

// Generate priority
func rPriority() string {
	var priorities []string = []string{
		"hight",
		"meddium",
		"low",
	}
	return priorities[rand.Intn(len(priorities))]
}
