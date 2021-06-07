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

func main() {
	for {
		var res string
		var rStdin *bufio.Reader = bufio.NewReader(os.Stdin)

		fmt.Println("1. Add user")
		fmt.Println("2. Show users")
		fmt.Println("3. Exit")
		_, err := fmt.Scanln(&res)
		if err != nil {
			log.Fatal(err)
		}

		switch string(res) {
		case "1":

			var user user
			var res map[string]interface{}

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

			var data map[string]string = map[string]string{
				"name": user.Name,
				"pass": user.Pass,
			}
			dataJSON, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}

			response, err := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer(dataJSON))
			if response.StatusCode != 200 {
				log.Fatal(response.Status)
			}
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			err = json.NewDecoder(response.Body).Decode(&res)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(res["text"])

		case "2":

			var users []user
			response, err := http.Get("http://localhost:8080/show")
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
