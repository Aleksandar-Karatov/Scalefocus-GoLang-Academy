package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	var bb BartenderBot
	bb.Start()
}

type Drink struct {
	Name   string `json:"strDrink"`
	Recipe string `json:"strInstructions"`
}
type DrinkPayload struct {
	Drinks []Drink
}
type BartenderBot struct {
}

func (bb *BartenderBot) Start() {

	for {
		fmt.Print("Enter a cocktail name:")

		in := bufio.NewReader(os.Stdin)

		input, err := in.ReadString('\n')
		if strings.ToLower(input) == "nothing" {
			return
		}
		var splitInput = strings.Split(input, " ")
		input = ""
		for i := 0; i < len(splitInput); i++ {
			if splitInput[i] == "\n" {
				break
			}
			if i != len(splitInput)-1 {
				input += splitInput[i] + "%20"
			} else {
				input += splitInput[i]
			}

		}
		input = strings.Trim(strings.ToLower(input), "\n")
		input = strings.Trim(strings.ToLower(input), "\r")

		resp, err := http.Get("https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			continue
		}
		payload := DrinkPayload{}

		json.NewDecoder(resp.Body).Decode(&payload)
		sentences := strings.Split(payload.Drinks[0].Recipe, ". ")
		sentences[len(sentences)-1] = strings.TrimRight(sentences[len(sentences)-1], ".")
		fmt.Println(payload.Drinks[0].Name, " recipe: ")
		for _, sentence := range sentences {
			fmt.Println(sentence + ".")
		}
	}
}
