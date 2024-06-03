package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CepService struct {
	Name     string
	Url      string
	Response map[string]interface{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <CEP>")
		return
	}

	cep := os.Args[1]
	fmt.Println("CEP:", cep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	services := []CepService{
		{
			Name: "brasilapi",
			Url:  "https://brasilapi.com.br/api/cep/v1/" + cep,
		},
		{
			Name: "viacep",
			Url:  "https://viacep.com.br/ws/" + cep + "/json/",
		},
	}

	ch := make(chan CepService, 1)
	for _, service := range services {
		go fetch(ctx, service, ch)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Request timeout and was canceled by the context")
	case result := <-ch:
		fmt.Println("Resultado mais rápido:", result.Name)
		s, _ := json.MarshalIndent(result.Response, "", "\t")
		fmt.Print(string(s))
	}
}

func fetch(ctx context.Context, service CepService, ch chan<- CepService) {
	fmt.Printf("Buscando em %s...\n", service.Name)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, service.Url, nil)
	if err != nil {
		fmt.Println("Error building request:", err)
		return
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Printf("Error doing request to %s: %v\n", service.Name, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(res.Body)

	response := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	service.Response = response
	select {
	case ch <- service:
	default:
	}
}
