package typescielo

var (
	// URLCieloProd return url to payment on production metoth used POST|PUT
	URLCieloProd string = "https://api.cieloecommerce.cielo.com.br"
	// URLClieloSandBox return url to payment on dev metoth used POST|PUT
	URLClieloSandBox string = "https://apisandbox.cieloecommerce.cielo.com.br"
	// URLCieloProdGET return url from query on production
	URLCieloProdGET string = "https://apiquery.cieloecommerce.cielo.com.br"
	// URLClieloSandBoxGET return url from query on sandbox clielo
	URLClieloSandBoxGET string = "https://apiquerysandbox.cieloecommerce.cielo.com.br"

	Domain string

	Provider string
)

const (
	// MerchantID code of access
	MerchantID string = "f25e01c5-d5ef-46d7-9f7d-6ade9beaaa82"
	// MerchantKey code of key access
	MerchantKey string = "KRWIRWBTVGWXVEHIGLWXJYCHNBJQHUIXALJTDHOP"
	// URISale respresents the uri to send one sales
	URISale string = "1/sales/"
	// URICadCreditCard represents the uri to send data of credit card to register
	URICadCreditCard string = "1/card/"
)

var (
	// StatPayNotFinish Aguardando atualização de status
	StatPayNotFinish int64
	// StatPayAuthorized Pagamento apto a ser capturado ou definido como pago
	StatPayAuthorized int64 = 1
	// StatPayConfirmed Pagamento confirmado e finalizado
	StatPayConfirmed int64 = 2
	// StatPayDenied Pagamento negado por Autorizador
	StatPayDenied int64 = 3
	// StatPayVoided Pagamento cancelado
	StatPayVoided int64 = 10
	// StatPayRefunded Pagamento cancelado após 23:59 do dia de autorização
	StatPayRefunded int64 = 11
	// StatPayPening Aguardando Status de instituição financeira
	StatPayPening int64 = 12
	// StatPayAborted Pagamento cancelado por falha no processamento ou por ação do AF
	StatPayAborted int64 = 13
	// StatPayScheduled Recorrência agendada
	StatPayScheduled int64 = 20
)

const (
	// CD type to debit card
	CD string = "DebitCard"
	// CC refferer that credit card
	CC string = "CreditCard"
	// BOL refferer that boleto
	BOL string = "Boleto"
)

const (
	// NumInstallments represents the number of installments
	NumInstallments int64 = 1 //numero de parcelas
	// ConsSavedCard represents if are from save the data of credit | debit card
	ConsSavedCard bool = true
)

// Payment represents the struct of payment
type Payment struct {
	ServiceTaxAmount    int64         `json:"ServiceTaxAmount"`
	Installments        int64         `json:"Installments"`
	Interest            int64         `json:"Interest"`
	Capture             bool          `json:"Capture"`
	Authenticate        bool          `json:"Authenticate"`
	Recurrent           bool          `json:"Recurrent"`
	CreditCard          *CreditCard   `json:"CreditCard"`
	ProofOfSale         string        `json:"ProofOfSale"`
	Tid                 string        `json:"Tid"`
	AuthorizationCode   string        `json:"AuthorizationCode"`
	PaymentID           string        `json:"PaymentId"`
	Type                string        `json:"Type"`
	Amount              int64         `json:"Amount"`
	SoftDescriptor      string        `json:"SoftDescriptor"`
	Provider            string        `json:"Provider"`
	Currency            string        `json:"Currency"`
	Country             string        `json:"Country"`
	ExtraDataCollection []interface{} `json:"ExtraDataCollection"`
	Status              int64         `json:"Status"`
	ReturnCode          string        `json:"ReturnCode"`
	ReturnMessage       string        `json:"ReturnMessage"`
	Links               *[]Link       `json:"Links"`
	ReceivedDate        string        `json:"ReceivedDate"`
	IsSplitted          bool          `json:"IsSplitted"`
}

// Link represents the struct of links returned
type Link struct {
	Method string `json:"Method"`
	Rel    string `json:"Rel"`
	Href   string `json:"Href"`
}

// CreditCard represents the struct to register/received credit card to paymenet
type CreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SaveCard       bool   `json:"SaveCard"`
	Brand          string `json:"Brand"`
	SecurityCode   string `json:"SecurityCode"`
	CardToken      string `json:"CardToken"`
}

// Customer represents struct of customer to payment
type Customer struct {
	Name string `json:"Name"`
}

// Order represnets the return of payment
type Order struct {
	MerchantOrderID string    `json:"MerchantOrderId"`
	Customer        *Customer `json:"Customer"`
	Payment         *Payment  `json:"Payment"`
}

// DataToPayment represent the data to generate the register of payment
type DataToPayment struct {
	TypePayment         string `json:",omitempty"`
	BillingOrderID      string `json:",omitempty"`
	NameCustomer        string `json:",omitempty"`
	ValuePayment        int64  `json:",omitempty"`
	NumberInstallments  int64  `json:",omitempty"`
	InvoicesDescription string `json:",omitempty"`
	CardNumber          string `json:",omitempty"`
	TokenCard           string `json:",omitempty"`
	NamePrintedOnCard   string `json:",omitempty"`
	ExpirationDate      string `json:",omitempty"`
	SaveCard            bool   `json:",omitempty"`
	Brand               string `json:",omitempty"`
	CodeCVC             string `json:",omitempty"`
}

// ReqDataCard ...
type ReqDataCard struct {
	Card          RequestCreditCard `json:"card_data"`
	CodeCVC       string            `json:"cvc"`
	TypePayment   string            `json:"type_payment"`
	ChangeDefault bool              `json:"change_default"`
}

// RequestCreditCard represents the request to register of credit card
type RequestCreditCard struct {
	CustomerName   string `json:"CustomerName"`
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	Brand          string `json:"Brand"`
}

// ResultGenerated receive of result retuned of api cielo
type ResultGenerated struct {
	CardToken string      `json:"CardToken"`
	Links     interface{} `json:"Links"`
}
