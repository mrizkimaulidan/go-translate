package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ResponseStruct is a struct for return the JSON response
type ResponseStruct struct {
	Data struct {
		Origin string
		Result string
	}
}

func main() {
	fmt.Print("Masukkan kalimat : ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	// Parsed the space characters to %20
	parsedText := strings.ReplaceAll(text, " ", "%20")

	url := "https://api-translate.azharimm.tk/translate?engine=google&text=" + parsedText + "&to=en"

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	responseByte, _ := io.ReadAll(response.Body)

	var responseStruct ResponseStruct

	json.Unmarshal(responseByte, &responseStruct)

	fmt.Println("Kata yang dimasukkan : ", text)
	fmt.Println("Hasil translate : ", responseStruct.Data.Result)
}
