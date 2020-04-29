package gocielo

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/DiegoSantosWS/gocielo/execute"
	"github.com/DiegoSantosWS/gocielo/typescielo"
	"github.com/DiegoSantosWS/gocielo/utilscielo"
)

func TestMain(t *testing.T) {
	t.Skip("PULANDO...")
}

func ExempleCreatePaymenet() {
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

func ExempleAddCreditCard() {
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

func ExempleGetCreditCard() {
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

func ExempleGetCreditCard() {
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
