package main_test

import (
	"testing"

	"github.com/DiegoSantosWS/gocielo/execute"
	"github.com/DiegoSantosWS/gocielo/typescielo"
	"github.com/DiegoSantosWS/gocielo/utilscielo"
	"github.com/gofrs/uuid"
)

func TestPayment(t *testing.T) {
	orderID, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}
	dPay := &typescielo.DataToPayment{
		BillingOrderID:      orderID.String(),
		ValuePayment:        150050,
		NumberInstallments:  1,
		InvoicesDescription: "Test pay lib",
		CardNumber:          "1234123412341231",
		// TokenCard:           "146afae7-76c2-4124-96fc-8902cccf4f5f",
		NamePrintedOnCard: "JOAO DA SILVA",
		ExpirationDate:    "12/2030",
		SaveCard:          true,
		Brand:             "Visa",
		CodeCVC:           "123",
		NameCustomer:      "João da Silva Mendes",
		TypePayment:       "CreditCard",
	}
	card := execute.GetCreditCard(dPay)
	payment := execute.CreatePayment(dPay)
	orderResult, err := execute.ExecPayment(dPay, card, payment)
	if err != nil {
		t.Errorf("[ tests ] Error to register payment. status [%d], Error [%s]", orderResult.Payment.Status, err)
	}

	if orderResult.Payment.Status != int64(2) {
		t.Errorf("[ tests ] Payment not registed status %d", orderResult.Payment.Status)
	}

	// log.Println(fmt.Sprintf("%s", orderResult.Payment.CreditCard.CardToken))
	// log.Println(fmt.Sprintf("%s", orderResult.Payment.ReturnMessage))
	utilscielo.DisplayObjectFormatJSON(orderResult)
}
