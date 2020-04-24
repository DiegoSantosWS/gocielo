# GOCIELO

[![GoDoc](https://godoc.org/github.com/DiegoSantosWS/gocielo?status.svg)](https://godoc.org/github.com/DiegoSantosWS/gocielo) [![Build Status](https://travis-ci.org/DiegoSantosWS/gocielo.svg?branch=master)](https://travis-ci.org/DiegoSantosWS/gocielo)

Lib de integração com gateway de pagamento CIELO

## Como usar

```bash

$ mkdir your-project 
$ cd your-project
$ go mod github.com/[your-username]/your-project
$ go get github.com/DiegoSantosWS/gocielo

```

## Exemplo 1

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
}
```

## Exemplo 2

### Executando criando hash do cartão

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
}
```