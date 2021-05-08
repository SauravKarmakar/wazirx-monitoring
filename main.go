package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jasonlvhit/gocron"
	log "github.com/kataras/golog"
	"github.com/sfreiberg/gotwilio"
)

const (
	marketStatusURL = "https://api.wazirx.com/api/v2/market-status"
	bttThreshold    = 0.65
)

func Monitoring(lastValue float64) {
	// Condition to check the threshold
	if lastValue >= bttThreshold {
		fmt.Println("BTT crossed threshold !!!, sending notification.")
		sendWhatsAppNotification(lastValue)
	} else {
		fmt.Println("No new notification, value less than threshold.")
	}
}

func main() {
	log.Info("-- Wazirx Monitoring --")

	// Reach to the Wazirx API to get the market value
	lastValue := getBTTLastValue()

	// CRON JOB to check in every Minute
	gocron.Every(5).Minute().Do(Monitoring, lastValue)
	<-gocron.Start()

	// WITHOUT CRON, for debugging
	// Monitoring(lastValue)
}

func getBTTLastValue() float64 {

	lastValue := 0.0

	// make an http call and get the value
	resp, err := http.Get(marketStatusURL)
	if err != nil {
		log.Error("Error while getting the market status from WAZIRX. URL: " + marketStatusURL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Response for market-status is not ok, status code: ", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error while reading the response body")
	}

	data := WazirxMarkets{}

	marshalerr := json.Unmarshal(body, &data)
	if marshalerr != nil {
		log.Error("Error while marshalling the markey status data. ERROR: ", marshalerr)
	}

	for _, entity := range data.Markets {
		if entity.BaseMarket == "btt" && entity.QuoteMarket == "inr" {
			// fmt.Println(entity.Last)
			lastValue, _ = strconv.ParseFloat(entity.Last, 32)
		}
	}
	return lastValue
}

func sendWhatsAppNotification(lastValue float64) {
	// WHAT'S UP INTEGRATION
	accountSid := "ACcb47220b40a626be5c1585d7c24bad03"
	authToken := "644415d08a82a21f9fa447bbffb5434c"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+14155238886"
	to := "+917709431782"
	message := fmt.Sprintf("Wazirx: BTT crossed Threshold !!! \n\n\nLastValue: %f,\nThresholdValue: %f", lastValue, bttThreshold)
	// twilio.SendSMS(from, to, message, "", "")
	_, b, c := twilio.SendWhatsApp(from, to, message, "", "")
	if c != nil {
		log.Error("Error while sending the what's app message. Error: ", c)
	}
	if b != nil {
		log.Error("Exception while sending the what'sapp message, Error: ", b)
	}
	//fmt.Println(a)
}

// --- will be required for rounding of values --- //
// --- KEEP IT --- //
// func round(num float64) int {
// 	return int(num + math.Copysign(0.5, num))
//}

// func toFixed(num float64, precision int) float64 {
// 	output := math.Pow(10, float64(precision))
// 	return float64(round(num*output)) / output
// }
