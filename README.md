# Cotação USD/BRL - Go

Cliente e servidor HTTP com controle de timeout e persistência em SQLite.

---

## 🚀 Como rodar

### 1. Clonar o repositório

```bash
git clone https://github.com/adlermo/fullcycle-client-server-api.git
cd fullcycle-client-server-api
```

### 2. Instalar dependências

```bash
go mod tidy
```

---

## ▶️ Executar

### Subir o servidor

```bash
go run cmd/server/main.go
```

Servidor disponível em:
http://localhost:8080/cotacao

---

### Rodar o cliente

Em outro terminal:

```bash
go run cmd/client/main.go
```

---

## 📄 Saída

Arquivo gerado:

```
cotacao.txt
```

Conteúdo:

```
Dólar: {valor}
```

---

## ⏱️ Timeouts

* API externa: 200ms
* Banco de dados: 10ms
* Cliente: 300ms
