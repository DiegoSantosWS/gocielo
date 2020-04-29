# GOCIELO

[![GoDoc](https://godoc.org/github.com/DiegoSantosWS/gocielo?status.svg)](https://godoc.org/github.com/DiegoSantosWS/gocielo) [![Build Status](https://travis-ci.org/DiegoSantosWS/gocielo.svg?branch=master)](https://travis-ci.org/DiegoSantosWS/gocielo)

Lib **gocielo** foi desenvolvida pensando em facilitar a integração com a cielo, e podendo ser ficar disponivel para a comunidade.

- Visão
	* O metodo ***CreatePayment*** monta os dados de pagamento um cartão com *tokenizado*. veja a implementação no exemplo [Exemplo 1](README.md#Exemplo-1)

	* O metodo ***AddCreditCard*** é usado para criar/add um cartão com *tokenizado*. veja a implementação no exemplo [Exemplo 2](README.md#Exemplo-2)

	* O metodo ***GetCreditCard*** é usado para mondar os dados do cartão com ou sem *token*, essa informção será usada para executar o pagamento. veja a implementação no exemplo [Exemplo 3](README.md#Exemplo-3)

	* O metodo ***ExecPaymentCreditCard*** é responsável por executar o pagamento, seu retornor e o objeto *Order* onde contem informação do aconteceu. [Exemplo 4](README.md#Exemplo-4)

----

# Observação

> A lib contem algumas funções que podem ajudar uma na implementação, uma delas é **ConvertFloatToCents** converte ``float64`` para centavos em ``int64`` póis é neste formato que a cielo espera recever o valor do pagamento. 
Para ver mais função, pode ser encontrada no aquivo utilscielo.go

## Como usar

```bash

$ mkdir your-project 
$ cd your-project
$ go mod github.com/[your-username]/your-project
$ go get github.com/DiegoSantosWS/gocielo

```

## Exemplo 1

#### Criando Pagamento

```go
func main() {
	dataPayment := &typescielo.DataToPayment{
		TypePayment:         typescielo.CC,
		BillingOrderID:      "122",
		NameCustomer:        "Name from your client",
		ValuePayment:        utilscielo.ConvertFloatToCents(1500.50),
		NumberInstallments:  typescielo.NumInstallments,
		InvoicesDescription: "Test payment",
		CardNumber:          "1234123412341231",
		// TokenCard:           "",
		NamePrintedOnCard: "JOÃO D O SILVA",
		ExpirationDate:    "12/2030",
		SaveCard:          typescielo.ConsSavedCard,
		Brand:             "Visa",
		CodeCVC:           "123",
	}
	pay := execute.CreatePayment(dataPayment)

	utilscielo.DisplayObjectFormatJSON(pay)

	// Output:
	// - typescielo.Payment
	// {
	// 	"Amount": 150050,
	// 	"Installments": 1,
	// 	"SoftDescriptor": "Test payment",
	// 	"Provider":"Cielo"
	// }
}
```

## Exemplo 2

### Criando hash do cartão

```go
func main() {
	request := typescielo.ReqDataCard{
		Card: typescielo.RequestCreditCard{
			CustomerName:   "JOAO DA SILVA",
			CardNumber:     "<NAMBER_FROM_CARD>",
			Holder:         "DIEGO DOS SANTOS",
			ExpirationDate: "11/2024",
			Brand:          "ELO",
		},
		CodeCVC:       "123",
		TypePayment:   "CreditCard",
		ChangeDefault: true,
	}

	ok, body, err := execute.AddCreditCard(req)
	if !ok || err != nil {
		log.Printf("Cannot add the card [%s]", err)
	}

	exp := typescielo.ResultGenerated{}
	err = json.Unmarshal(body, &exp)
	if err != nil {
		log.Printf("Cannot add the card [%s]", err)
		return
	}
	// está func imprime o resultado no terminal
	utilscielo.DisplayObjectFormatJSON(exp)

	// Output:
	// - typescielo.ResultGenerated
	// {
	// 	"CardToken": "HASH_FROM_CREDIT_CARD_TOKENIZADE",
	// 	"Links": []
	// }
}
```

## Exemplo 3

### Montando as informações do cartão

```go
func main() {
	dataPayment := &typescielo.DataToPayment{
		TypePayment:         typescielo.CC,
		BillingOrderID:      "122",
		NameCustomer:        "Name from your client",
		ValuePayment:        utilscielo.ConvertFloatToCents(1500.50),
		NumberInstallments:  typescielo.NumInstallments,
		InvoicesDescription: "Test payment",
		CardNumber:          "1234123412341231",
		// TokenCard:           "",
		NamePrintedOnCard: "JOÃO D O SILVA",
		ExpirationDate:    "12/2030",
		SaveCard:          typescielo.ConsSavedCard,
		Brand:             "Visa",
		CodeCVC:           "123",
	}
	
	card := execute.GetCreditCard(dataPayment)

	// está func imprime o resultado no terminal
	utilscielo.DisplayObjectFormatJSON(card)

	// OutPut:
	// - typescielo.CreditCard
	// {
	// 	"CardNumber":"1234123412341231"
	// 	"Holder":"JOÃO D O SILVA"
	// 	"ExpirationDate":"12/2030"
	// 	"SaveCard":true
	// 	"Brand":"Visa"
	// 	"SecurityCode":"123"
	// 	"CardToken":"token enviado pela cielo"
	// }
}
```

## Exemplo 4

### Executando pagamento

```go
func main() {
	dataPayment := &typescielo.DataToPayment{
		TypePayment:         typescielo.CC,
		BillingOrderID:      "122",
		NameCustomer:        "Name from your client",
		ValuePayment:        utilscielo.ConvertFloatToCents(1500.50),
		NumberInstallments:  typescielo.NumInstallments,
		InvoicesDescription: "Test payment",
		CardNumber:          "1234123412341231",
		// TokenCard:           "",
		NamePrintedOnCard: "JOÃO D O SILVA",
		ExpirationDate:    "12/2030",
		SaveCard:          typescielo.ConsSavedCard,
		Brand:             "Visa",
		CodeCVC:           "123",
	}
	pay := execute.CreatePayment(dataPayment)
	card := execute.GetCreditCard(dataPayment)
	order, err := execute.ExecPaymentCreditCard(dataPayment, card, pay)
	if err != nil {
		return
	}
	
	// está func imprime o resultado no terminal
	utilscielo.DisplayObjectFormatJSON(order)

	// OutPut:
	// - typescielo.Order
	// Este resultado é enorme...	
}
```



# TODO

- Pagamento por cartão de débito
- Pagamento por boleto


# License

MIT - Veja [LICENSE](https://github.com/DiegoSantosWS/gocielo/blob/master/LICENSE) file