package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET",
		"http://localhost:8080/cotacao", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro/timeout:", err)
		return
	}
	defer resp.Body.Close()

	var data map[string]string
	json.NewDecoder(resp.Body).Decode(&data)

	file, _ := os.Create("cotacao.txt")
	defer file.Close()

	file.WriteString(fmt.Sprintf("Dólar: %s", data["bid"]))
}