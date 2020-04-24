package main

import (
	"fmt"
	"log"

	"github.com/DiegoSantosWS/gocielo/execute"
	"github.com/DiegoSantosWS/gocielo/typescielo"
	"github.com/DiegoSantosWS/gocielo/utilscielo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(fmt.Sprintf("[main gocielo.init] Thet error to read the file from .env Error [%s]", err))
	}
}

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
