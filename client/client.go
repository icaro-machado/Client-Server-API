package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	price, err := findPrice(ctx)
	if err != nil {
		log.Fatal("Erro ao buscar cotação:", err)
	}

	if err := savePriceTxt(price); err != nil {
		log.Fatal("Erro ao salvar cotação no arquivo:", err)
	}

	fmt.Println("Cotação salva com sucesso!")
}

func findPrice(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseBody map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	value, ok := responseBody["bid"]
	if !ok {
		return "", fmt.Errorf("Campo 'bid' não encontrado")
	}
	return value, nil
}

func savePriceTxt(price string) error {
	content := fmt.Sprintf("Dólar: %s", price)
	return ioutil.WriteFile("cotacao.txt", []byte(content), 0644)
}
