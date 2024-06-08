package client

import (
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioClient struct {
	trc        *twilio.RestClient
	accountSID string
}

func NewTwilioClient(sid string, key string, accountSID string) (*TwilioClient, error) {
	tcParams := twilio.ClientParams{Username: sid, Password: key, AccountSid: sid}
	trc := twilio.NewRestClientWithParams(tcParams)
	tc := &TwilioClient{trc: trc, accountSID: accountSID}

	return tc, nil
}

func (tc *TwilioClient) GetBalance() (string, error) {
	fetchParams := &api.FetchBalanceParams{
		PathAccountSid: &tc.accountSID,
	}
	bal, err := tc.trc.Api.FetchBalance(fetchParams)
	if err != nil {
		return "", err
	}

	return *bal.Balance, nil
}
