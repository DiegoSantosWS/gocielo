package utilscielo

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var loc *time.Location

//GetLocation returns the location of america / são paulo
func getLocation() (*time.Location, error) {
	if loc == nil {
		loc = time.FixedZone("America/Sao_Paulo", -3*60*60)
	}
	return loc, nil
}

// ConvertFloatToCents convert float64 to int64 used to add value of Amount of cielo
func ConvertFloatToCents(fl float64) int64 {
	return int64(fl * 100)
}

// ExtractLastNumbers extract last charachters of an string
func ExtractLastNumbers(s string, n int) string {
	lengthStr := len(s)
	if lengthStr == 0 || n >= lengthStr {
		return ""
	}
	return s[lengthStr-n:]
}

// ConvertExpirationDateToTime convert date to time.time
func ConvertExpirationDateToTime(dt string) time.Time {
	if len(dt) == 0 {
		return time.Now()
	}
	exDte := strings.Split(dt, "/")
	m, _ := strconv.Atoi(exDte[0])
	y, _ := strconv.Atoi(exDte[1])
	l, err := getLocation()
	if err != nil {
		log.Println("[gocielo utilscielo.ConvertExpirationDateToTime] Has an error on hour...", err)
		return time.Date(y, time.Month(m), 30, 12+3, 0, 0, 0, time.UTC)
	}
	date := time.Date(y, time.Month(m), 30, 12, 0, 0, 0, l)
	return endOfMonth(date)
}

// DisplayObjectFormatJSON so para mostrar o json no terminal
func DisplayObjectFormatJSON(obj interface{}) {
	prettyJSON, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	fmt.Printf("The result\n %s\n", string(prettyJSON))
}

//EndOfMonth returns the end of the month ot t time
func endOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	ini := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	end := ini.AddDate(0, 1, 0).Add(-time.Nanosecond)
	endY, endM, endD := end.Date()
	return time.Date(endY, endM, endD, 12, 0, 0, 0, time.UTC)
}
