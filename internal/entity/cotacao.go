package entity

type Cotacao struct {
	Moeda           string  `json:"moeda"`
	Nome            string  `json:"nome"`
	Compra          float64 `json:"compra"`
	Venda           float64 `json:"venda"`
	FechoAnterior   float64 `json:"fechoAnterior"`
	DataAtualizacao string  `json:"dataAtualizacao"`
}
