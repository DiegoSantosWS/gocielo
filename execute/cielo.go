package execute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/DiegoSantosWS/gocielo/typescielo"
	"github.com/DiegoSantosWS/gocielo/utilscielo"
)

var merchantID, merchantKEY string

func init() {
	prod := os.Getenv("ENV")
	if prod == "prod" {
		typescielo.Domain = fmt.Sprintf("%s/", typescielo.URLCieloProd)
		typescielo.Provider = "Cielo"
		merchantID = os.Getenv("CIELO_MERCHANT_ID")
		merchantKEY = os.Getenv("CIELO_MERCHANT_KEY")
	} else {
		typescielo.Domain = fmt.Sprintf("%s/", typescielo.URLClieloSandBox)
		typescielo.Provider = "Simulado"
		merchantID = typescielo.MerchantID
		merchantKEY = typescielo.MerchantKey
	}
}

// CreatePayment return of data to payment
func CreatePayment(cc *typescielo.DataToPayment) *typescielo.Payment {
	pay := &typescielo.Payment{
		Amount:         cc.ValuePayment,
		Installments:   cc.NumberInstallments,
		SoftDescriptor: cc.InvoicesDescription,
		Provider:       typescielo.Provider,
	}
	return pay
}

// GetCreditCard return the datas of credit card ", number, holder, expirationDate, brand, securityCode string, saveCard bool"
func GetCreditCard(cc *typescielo.DataToPayment) *typescielo.CreditCard {
	credC := &typescielo.CreditCard{
		CardNumber:     cc.CardNumber,
		CardToken:      cc.TokenCard,
		Holder:         cc.NamePrintedOnCard,
		ExpirationDate: cc.ExpirationDate,
		SaveCard:       cc.SaveCard,
		Brand:          cc.Brand,
		SecurityCode:   cc.CodeCVC,
	}
	return credC
}

// ExecPaymentCreditCard register the payment on cielo and return the result of requisition
func ExecPaymentCreditCard(dPay *typescielo.DataToPayment, card *typescielo.CreditCard, payment *typescielo.Payment) (*typescielo.Order, error) {
	dataPay, err := json.Marshal(
		&typescielo.Order{
			MerchantOrderID: dPay.BillingOrderID,
			Customer: &typescielo.Customer{
				Name: dPay.NameCustomer,
			},
			Payment: &typescielo.Payment{
				Type:           dPay.TypePayment,
				Amount:         payment.Amount,
				Installments:   payment.Installments,
				SoftDescriptor: payment.SoftDescriptor,
				Capture:        true,
				CreditCard: &typescielo.CreditCard{
					CardNumber:     card.CardNumber,
					Holder:         card.Holder,
					ExpirationDate: card.ExpirationDate,
					SecurityCode:   card.SecurityCode,
					Brand:          card.Brand,
					SaveCard:       true,
					CardToken:      card.CardToken,
				},
				Country:   "BRA",
				Currency:  "BRL",
				Provider:  typescielo.Provider,
				Recurrent: true,
			}})
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] couldnt the marshal of data received, transaction_id [%s]. Error [%s]", dPay.BillingOrderID, err))
		return nil, err
	}

	url := fmt.Sprintf("%s%s", typescielo.Domain, typescielo.URISale)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataPay))
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] Faild to send of request to url [%s], transaction_id [%s]. Error [%s]", url, dPay.BillingOrderID, err))
		return nil, err
	}
	req.Header.Add("MerchantId", merchantID)
	req.Header.Add("MerchantKey", merchantKEY)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] Faild to execute the request, transaction_id [%s]. Error [%s]", dPay.BillingOrderID, err))
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] couldn't read the data of return, transaction_id [%s]. Error [%s]", dPay.BillingOrderID, err))
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		cErr := []utilscielo.ErrorStr{}
		err := json.Unmarshal(body, &cErr)
		if err != nil {
			return nil, err
		}
		err = utilscielo.ErrorString(cErr)
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] Faild to execute the request, url [%s] status [%s] transaction_id [%s]. Error [%s]", url, resp.Status, dPay.BillingOrderID, err))
		return nil, err
	}

	var order *typescielo.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.ExecPayment ] couldn't unmarshal to body, transaction_id [%s]. Error [%s]", dPay.BillingOrderID, err))
		return nil, err
	}
	return order, nil
}

// AddCreditCard register the data from card join gateway from payment
func AddCreditCard(r typescielo.ReqDataCard) (ok bool, body []byte, err error) {
	dataCC, err := json.Marshal(r.Card)
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s%s", typescielo.Domain, typescielo.URICadCreditCard)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataCC))
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo execute.addCreditCard ] Faild to send of request to url [%s]. Error [%s]", url, err))
		return
	}
	req.Header.Add("MerchantId", merchantID)
	req.Header.Add("MerchantKey", merchantKEY)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo cmd.addCreditCard ] Faild to execute the request. Error [%s]", err))
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("[ gocielo cmd.addCreditCard ] couldn't read the data of return. Error [%s]", err))
		return
	}

	if resp.StatusCode >= http.StatusBadRequest {
		cErr := []utilscielo.ErrorStr{}
		err = json.Unmarshal(body, &cErr)
		if err != nil {
			return
		}
		err = utilscielo.ErrorString(cErr)
		log.Println(fmt.Sprintf("[ gocielo cmd.addCreditCard ] Faild to execute the request, url [%s] status [%s]. Error [%s]", url, resp.Status, err))
		return
	}

	ok = true
	return ok, body, err
}
