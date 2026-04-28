package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"fullcycle-client-server-api/internal/database"
	"fullcycle-client-server-api/internal/entity"
)

func CotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctxAPI, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		req, err := http.NewRequestWithContext(ctxAPI, "GET",
			"https://br.dolarapi.com/v1/cotacoes/usd", nil)
		if err != nil {
			log.Println("erro ao criar request:", err)
			http.Error(w, "erro interno", http.StatusInternalServerError)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("API timeout:", err)
			http.Error(w, "timeout API", http.StatusGatewayTimeout)
			return
		}
		defer resp.Body.Close()

		var cotacao entity.Cotacao
		if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
			log.Println("erro decode:", err)
			http.Error(w, "erro ao processar resposta", http.StatusInternalServerError)
			return
		}

		if cotacao.Compra == 0 {
			log.Println("cotação inválida (compra = 0)")
			http.Error(w, "cotação inválida", http.StatusInternalServerError)
			return
		}

		valor := fmt.Sprintf("%.4f", cotacao.Compra)

		log.Println("Cotação recebida:", valor)

		ctxDB, cancelDB := context.WithTimeout(r.Context(), 10*time.Millisecond)
		defer cancelDB()

		if err := database.SaveCotacao(ctxDB, db, valor); err != nil {
			log.Println("DB timeout:", err)
		} else {
			log.Println("Salvo no banco:", valor)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"bid": valor,
		})
	}
}
