# GOCIELO

[![GoDoc](https://godoc.org/github.com/DiegoSantosWS/gocielo?status.svg)](https://godoc.org/github.com/DiegoSantosWS/gocielo) [![Build Status](https://travis-ci.org/DiegoSantosWS/gocielo.svg?branch=master)](https://travis-ci.org/DiegoSantosWS/gocielo)

Lib de integração com gateway de pagamento CIELO

## Como usar

Para usar a lib siga o exemplo abaixo.

```go
func main() {
	dataPayment := &typescielo.DataToPayment{
		TypePayment:         typescielo.CC,
		BillingOrderID:      "122",
		NameCustomer:        "Name from your client",
		ValuePayment:        utilscielo.ConvertFloatToCents(1500.50),
		NumberInstallments:  typescielo.NumInstallments,
		InvoicesDescription: "Test payment",
		CardNumber:          "1234123412341231", //
		// TokenCard:           "",
		NamePrintedOnCard: "JOÃO D O SILVA",
		ExpirationDate:    "12/2030",
		SaveCard:          typescielo.ConsSavedCard,
		Brand:             "Visa",
		CodeCVC:           "123",
	}
	pay := execute.CreatePayment(dataPayment)
	card := execute.GetCreditCard(dataPayment)
	order, err := execute.ExecPayment(dataPayment, card, pay)
	if err != nil {
		return
	}
	utilscielo.DisplayObjectFormatJSON(order)
}
```