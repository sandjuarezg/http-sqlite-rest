package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

var client *http.Client = &http.Client{}

func main() {
	for {
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

			dataJSON, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}

			request, err := http.NewRequest("POST", "http://localhost:8080/add", bytes.NewBuffer(dataJSON))
			if err != nil {
				log.Fatal(err)
			}

			request.Header.Set("Accept", "application/json")

			response, err := client.Do(request)
			if response.StatusCode != 200 {
				log.Fatal(response.Status)
			}
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			var data map[string]string
			err = json.NewDecoder(response.Body).Decode(&data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(data["text"])

		case "2":

			var users []user

			request, err := http.NewRequest("GET", "http://localhost:8080/show", nil)
			if err != nil {
				log.Fatal(err)
			}

			request.Header.Set("Accept", "application/json")

			response, err := client.Do(request)
			if response.StatusCode != 200 {
				log.Fatal(response.Status)
			}
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			err = json.NewDecoder(response.Body).Decode(&users)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("|%-7s|%-15s|%-15s|\n", "id", "Name", "Password")
			fmt.Println("_________________________________________")
			for i := 0; i < len(users); i++ {
				fmt.Printf("|%-7d|%-15s|%-15s|\n", users[i].Id, users[i].Name, users[i].Pass)
			}
			fmt.Println()

		case "3":

			fmt.Println("E X I T I N G . . .")
			os.Exit(0)

		}
	}
}
