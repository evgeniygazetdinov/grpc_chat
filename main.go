package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v4"
)

type SomeUser struct {
	Gender     string
	Name       string
	Location   string
	Email      string
	AppVersion string
	Bio        string
	Token      string
	UserTokens string
	IdUser     int
	Interview  false
}

func getRandomLocation() string {
	// Generate a random Float64 type number
	lati := rand.Float64() * 200
	lot := rand.Float64() * 200
	latString := strconv.FormatFloat(lati, 'f', 6, 64)
	longiString := strconv.FormatFloat(lot, 'f', 6, 64)
	return fmt.Sprintf("POINT(%s %s)", latString, longiString)
}

func getRandomGender() string {
	n := rand.Intn(2)
	result := ""
	switch n {
	case 0:
		result = "f"
	case 1:
		result = "m"

	}
	return result
}
func getRandomLanguage() string {
	n := rand.Intn(2)
	result := ""
	switch n {
	case 0:
		result = "RU"
	case 1:
		result = "EN"

	}
	return result
}

func generateRandomLocation() string {
	return "someLocationString"
}

func main() {
	newUser := SomeUser{
		Name:     faker.Name(),
		Gender:   getRandomGender(),
		Email:    faker.Email(),
		Location: getRandomLocation(),
	}
	fmt.Println(newUser)
}
