package utilscielo

import (
	"errors"
	"fmt"
	"strings"
)

var codeReturned map[int64]string

// ErrorStr represents data erro
type ErrorStr struct {
	Code int64  `json:"Code"`
	Msg  string `json:"Message"`
}

// ErrorString return the error
func ErrorString(errsCielo []ErrorStr) error {
	var str []string
	for _, err := range errsCielo {
		str = append(str, fmt.Sprintf("%v - %s\n", err.Code, err.Msg))
	}
	return errors.New(strings.Join(str, "\n"))
}

// CodeErrorCard return the error received from api
func (c ErrorStr) CodeErrorCard() error {
	code := returnCodeError()
	str, ok := code[c.Code]
	if ok {
		strErr := fmt.Sprintf("%d - %s.", c.Code, str)
		return errors.New(strErr)
	}
	return nil
}

// returnCodeError represents the errors of api translated
func returnCodeError() map[int64]string {
	codeReturned = make(map[int64]string)
	codeReturned = map[int64]string{
		5:  "Não Autorizada",
		57: "Cartão Expirado",
		78: "Cartão Bloqueado",
		99: "Operation Successful, porem tempo de resposta esgotou",
		77: "Cartão Cancelado",
		70: "Problemas com o Cartão de Crédito",
	}
	return codeReturned
}
