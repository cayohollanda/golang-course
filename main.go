package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {
	log.Println("Iniciando script...")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição.")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao executar requisição.")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var userResponse []User

	json.Unmarshal([]byte(bodyBytes), &userResponse)

	for _, user := range userResponse {
		fmt.Printf("Usuário: %d\n", user.ID)
		fmt.Printf("Nome: %s\n", user.Name)
		fmt.Printf("Usuário: %s\n", user.Username)
		fmt.Printf("E-mail: %s\n", user.Email)
		fmt.Printf("Endereço: Rua %s, Suite %s, Cidade %s, Zipcode %s, Latitude %s, Longitude %s\n\n", user.Address.Street, user.Address.Suite, user.Address.City, user.Address.Zipcode, user.Address.Geo.Lat, user.Address.Geo.Lng)
	}

}
