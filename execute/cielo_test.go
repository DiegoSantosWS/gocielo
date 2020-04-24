package execute_test

import (
	"encoding/json"
	"testing"

	"github.com/DiegoSantosWS/gocielo/execute"
	"github.com/DiegoSantosWS/gocielo/typescielo"
	"github.com/DiegoSantosWS/gocielo/utilscielo"
	"github.com/gofrs/uuid"
)

func TestCreatePayment(t *testing.T) {
	casesPay := []struct {
		Name        string
		DataPayment *typescielo.DataToPayment
		Exp         *typescielo.Payment
	}{
		{
			Name: "JOAO DA SILVA",
			DataPayment: &typescielo.DataToPayment{
				BillingOrderID:      "",
				ValuePayment:        1500,
				NumberInstallments:  1,
				InvoicesDescription: "Test pay",
				CardNumber:          "1234123412341231",
				NamePrintedOnCard:   "JOAO D SILVA",
				ExpirationDate:      "12/2030",
				SaveCard:            true,
				Brand:               "Visa",
				CodeCVC:             "123",
				NameCustomer:        "JOAO DA SILVA",
			},
			Exp: &typescielo.Payment{
				Amount:         1500,
				Installments:   1,
				SoftDescriptor: "Test pay",
				Provider:       typescielo.Provider,
			},
		},
	}
	for _, pay := range casesPay {
		orderID, _ := uuid.NewV4()
		pay.DataPayment.BillingOrderID = orderID.String()
		t.Run(pay.Name, func(t *testing.T) {
			testcreatPayment(t, pay.Name, pay.DataPayment, pay.Exp)
		})
	}
}

func testcreatPayment(t *testing.T, name string, cc *typescielo.DataToPayment, exp *typescielo.Payment) {
	got := execute.CreatePayment(cc)
	validatePayment(t, name, got, exp)
}

func TestGetCreditCard(t *testing.T) {
	casesPay := []struct {
		Name        string
		DataPayment *typescielo.DataToPayment
		Exp         *typescielo.CreditCard
	}{
		{
			Name: "JOAO DA SILVA",
			DataPayment: &typescielo.DataToPayment{
				BillingOrderID:      "",
				ValuePayment:        1500,
				NumberInstallments:  1,
				InvoicesDescription: "Test pay",
				CardNumber:          "1234123412341231",
				NamePrintedOnCard:   "JOAO D SILVA",
				ExpirationDate:      "12/2030",
				SaveCard:            true,
				Brand:               "Visa",
				CodeCVC:             "123",
				NameCustomer:        "JOAO DA SILVA",
			},
			Exp: &typescielo.CreditCard{
				CardNumber:     "1234123412341231",
				CardToken:      "",
				Holder:         "JOAO D SILVA",
				ExpirationDate: "12/2030",
				SaveCard:       true,
				Brand:          "Visa",
				SecurityCode:   "123",
			},
		},
	}
	for _, pay := range casesPay {
		orderID, _ := uuid.NewV4()
		pay.DataPayment.BillingOrderID = orderID.String()
		t.Run(pay.Name, func(t *testing.T) {
			testGetCreditCard(t, pay.Name, pay.DataPayment, pay.Exp)
		})
	}
}

func testGetCreditCard(t *testing.T, name string, cc *typescielo.DataToPayment, exp *typescielo.CreditCard) {
	got := execute.GetCreditCard(cc)
	validateGetCreditCard(t, name, *got, *exp)
}

func TestAddCreditCard(t *testing.T) {
	casesReq := []struct {
		Name       string
		CreditCard typescielo.ReqDataCard
	}{
		{
			Name: "JOAO DA SILVA",
			CreditCard: typescielo.ReqDataCard{
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
			},
		},
	}
	for _, card := range casesReq {
		t.Run(card.Name, func(t *testing.T) {
			testAddCreditCard(t, card.Name, card.CreditCard)
		})
	}
}

func testAddCreditCard(t *testing.T, name string, req typescielo.ReqDataCard) {
	ok, body, err := execute.AddCreditCard(req)
	if !ok || err != nil {
		t.Errorf("alguma mensage [%s]", err)
	}
	exp := typescielo.ResultGenerated{}
	err = json.Unmarshal(body, &exp)
	if err != nil {
		t.Error(err)
	}
	utilscielo.DisplayObjectFormatJSON(exp)
}
