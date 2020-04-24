package typescielo

const (
	// StatSuccess represents status of success
	StatSuccess string = "4"
)

var msgByStatus map[int64]string

// GetStatusMessage return status message
func GetStatusMessage(stat int64) (msg string) {
	msgMap := returnStatusMessage()
	msg, ok := msgMap[stat]
	if ok {
		return msg
	}
	return msg
}

func returnStatusMessage() map[int64]string {
	msg := make(map[int64]string)
	msg = map[int64]string{
		StatPayNotFinish:  "Aguardando atualização de status",
		StatPayAuthorized: "Pagamento apto a ser capturado ou definido como pago",
		StatPayConfirmed:  "Pagamento confirmado e finalizado",
		StatPayDenied:     "Pagamento negado por Autorizador",
		StatPayVoided:     "Pagamento cancelado",
		StatPayRefunded:   "Pagamento cancelado após 23:59 do dia de autorização",
		StatPayPening:     "Aguardando Status de instituição financeira",
		StatPayAborted:    "Pagamento cancelado por falha no processamento ou por ação do AF",
		StatPayScheduled:  "Recorrência agendada",
	}

	return msg
}
