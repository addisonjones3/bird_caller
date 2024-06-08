package main

import (
	"fmt"
	"log"

	"github.com/addisonjones3/bird_caller/pkg/twilio/client"
	"github.com/joho/godotenv"
)

const (
	TWILIO_SID      = "TWILIO_ACCOUNT_SID"
	TWILIO_SECRET   = "TWILIO_AUTH_TOKEN"
	BIRD_CALLER_SID = "BIRD_CALLER_SID"
	BIRD_CALLER_KEY = "BIRD_CALLER_KEY"
)

var (
	logger           *log.Logger
	bcSID            string
	bcKey            string
	twilioAccountSID string
	found            bool
)

func init() {
	logger = log.Default()
	logger.Println("getting account vars")
	envMap, err := godotenv.Read("./.env")
	if err != nil {
		panic(err)
	}

	bcSID, found = envMap[BIRD_CALLER_SID]
	if !found {
		panic(fmt.Sprintf("did not find %s", BIRD_CALLER_SID))
	}
	bcKey, found = envMap[BIRD_CALLER_KEY]
	if !found {
		panic(fmt.Sprintf("did not find %s", BIRD_CALLER_KEY))
	}
	twilioAccountSID, found = envMap[TWILIO_SID]
	if !found {
		panic(fmt.Sprintf("did not find %s", TWILIO_SID))
	}
	logger.Println("account vars read successfully")
}

func main() {
	tc, err := client.NewTwilioClient(bcSID, bcKey, twilioAccountSID)
	logger.Println("twilio client created successfully")
	if err != nil {
		panic(err)
	}
	bal, err := tc.GetBalance()
	if err != nil {
		panic(err)
	}

	logger.Println(bal)

}
