package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ApiResponse struct {
	ShouldIDeploy bool   `json:"shouldideploy"`
	Message       string `json:"message"`
	Timezone      string `json:"timezone"`
	Date          string `json:"date"`
}

func main() {
    url := "https://shouldideploy.today/api"

	// Obter o timezone local    
	tz, exists := os.LookupEnv("TZ")
	if exists {
        url = fmt.Sprintf("https://shouldideploy.today/api?tz=%s", tz)        
	}

	// Request to the API
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na requisição GET: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Erro ao ler o corpo da resposta: %s\n", err)
		return
	}

	// JSON
	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("Erro ao decodificar a resposta JSON: %s\n", err)
		return
	}

    // Output
	fmt.Printf("Should I Deploy Today?")
	fmt.Printf("\n%s\n", apiResponse.Message)
}
