package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type SMSClient struct {
	hc         *http.Client
	accountSid string
	authToken  string
	from       string
}

func NewSMSClient(c Config) *SMSClient {
	return &SMSClient{
		hc: &http.Client{
			Timeout: time.Second * 10,
		},
		accountSid: c.TwilioAccountSid,
		authToken:  c.TwilioAuthToken,
		from:       c.TwilioFrom,
	}
}

func (s *SMSClient) Send(to string, body string) error {
	var (
		urlStr = fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", s.accountSid)
		req    *http.Request
		resp   *http.Response
		bs     []byte
		err    error
	)

	mobile := strings.ReplaceAll(to, " ", "")

	msgData := url.Values{}
	msgData.Set("To", mobile)
	msgData.Set("From", s.from)
	msgData.Set("Body", body)

	if req, err = http.NewRequest(http.MethodPost, urlStr, strings.NewReader(msgData.Encode())); err != nil {
		return err
	}

	req.SetBasicAuth(s.accountSid, s.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if resp, err = s.hc.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if bs, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		logrus.Debugf("sms sent to: %s, body: %s", mobile, body)

		return nil
	}

	return fmt.Errorf("sms gateway returned invalid status code: %d, body: %s", resp.StatusCode, bs)
}
