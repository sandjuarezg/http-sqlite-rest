package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type message struct {
	Text string `json:"text"`
}

var client *http.Client = &http.Client{}

func main() {
	for {
		var reply []byte = make([]byte, 1024)
		var rStdin *bufio.Reader = bufio.NewReader(os.Stdin)

		fmt.Println("1. Add user")
		fmt.Println("2. Show users")
		fmt.Println("3. Exit")
		reply, _, err := rStdin.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		switch string(reply) {
		case "1":

			reply = make([]byte, 1024)
			var user user

			fmt.Println("Enter a name")
			reply, _, err := rStdin.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			user.Name = string(reply)

			fmt.Println("Enter a password")
			reply, _, err = rStdin.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			user.Pass = string(reply)

			usersB, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}

			response, body := createNewRequest("POST", "http://localhost:8080/add", bytes.NewReader(usersB))

			var status int = response.StatusCode
			if status != 200 {
				log.Fatal(response.Status)
			}
			defer response.Body.Close()

			var mess message
			json.Unmarshal(body, &mess)
			fmt.Printf("%s\n\n", mess.Text)

		case "2":

			response, body := createNewRequest("GET", "http://localhost:8080/show", nil)

			var status int = response.StatusCode
			if status != 200 {
				log.Fatal(response.Status)
			}
			defer response.Body.Close()

			var users []user
			json.Unmarshal(body, &users)
			fmt.Printf("|%-7s|%-15s|%-15s|\n", "id", "Name", "Password")
			fmt.Println("_________________________________________")
			for i := 0; i < len(users); i++ {
				fmt.Printf("|%-7d|%-15s|%-15s|\n", users[i].Id, users[i].Name, users[i].Pass)
			}
			fmt.Println()

		case "3":

			fmt.Println("E X I T I N G . . .")
			os.Exit(0)

		default:

			fmt.Println("404 page not found")

		}
	}
}

func createNewRequest(method string, url string, content io.Reader) (response *http.Response, body []byte) {
	request, err := http.NewRequest(method, url, content)
	if err != nil {
		log.Fatal(err)
	}

	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return
}
