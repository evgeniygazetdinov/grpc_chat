package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grpcchat/database"
	"grpcchat/variable"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/bxcodec/faker/v4"
)

type SomeUser struct {
	Gender      string
	Name        string
	Location    string
	Email       string
	AppVersion  string
	Bio         string
	Token       string
	UserHeaders string
	IdUser      int
	Interview   bool
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
func getAppVersion() string {
	return "1.5.7"
}

func getUserHeaders(token string, host string) map[string]string {
	userHeaderData := make(map[string]string)
	userHeaderData["Authorization"] = fmt.Sprintf("JWT %s", token)
	userHeaderData["HOST"] = host
	return userHeaderData
}

func addHeadersOnRequest(someRequest *http.Request, host string) *http.Request {
	someRequest.Header.Add("Content-Type", "application/json")
	someRequest.Header.Add("Accept-Language", "ru")
	someRequest.Header.Add("accept", "application/json")
	someRequest.Header.Add("HOST", host)
	return someRequest
}

func getToken(email string, host string) string {
	token := ""
	client := &http.Client{}
	payloadEmail := map[string]string{"email": email}

	path := "/api/email/send/"
	postBody, _ := json.Marshal(payloadEmail)

	responseBody := bytes.NewBuffer(postBody)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", host, path), responseBody)
	requestWIthHeaders := addHeadersOnRequest(request, host)
	resp, err := client.Do(requestWIthHeaders)

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO correct parsing
	fmt.Println(body)
	sb := string(body)
	fmt.Printf(sb)
	return token
}

func (user *SomeUser) updateUnsettedData(host string) {
	// call that after user init
	user.Token = getToken(user.Email, host)

	// Store user in database
	userHeadersJSON, _ := json.Marshal(getUserHeaders(user.Token, host))
	err := database.StoreUser(
		user.Gender,
		user.Name,
		user.Location,
		user.Email,
		user.AppVersion,
		user.Bio,
		user.Token,
		string(userHeadersJSON),
		user.IdUser,
		user.Interview,
	)
	if err != nil {
		log.Printf("Error storing user in database: %v", err)
	}
}

func main() {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.DB.Close()

	newUser := SomeUser{
		Gender:      getRandomGender(),
		Name:        faker.Name(),
		Location:    getRandomLocation(),
		Email:       faker.Email(),
		AppVersion:  getAppVersion(),
		Bio:         faker.Paragraph(),
		Token:       "",
		UserHeaders: "",
		IdUser:      0,
		Interview:   false,
	}
	newUser.updateUnsettedData(variable.HOST)
}
