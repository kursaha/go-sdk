package mailkeets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequestDto struct {
	FromName          string `json:"fromName"`
	FromAddress       string `json:"fromAddress"`
	To                string `json:"to"`
	Subject           string `json:"subject"`
	ContentType       string `json:"contentType"`
	Body              string `json:"body"`
	RequestIdentifier string `json:"requestIdentifier"`
	UnsubscribedList  string `json:"unsubscribedList"`
}

type Mailkeets struct {
	apiKey  string
	baseUrl string
}

func NewMailkeets(apiKey string) *Mailkeets {
	return &Mailkeets{
		apiKey:  apiKey,
		baseUrl: "https://mailkeets.kursaha.com",
	}
}

func (m *Mailkeets) SendEmail(mail MailRequestDto) error {
	jsonBody, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", m.baseUrl+"/api/mail-send", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("mail send failed with status code %d", resp.StatusCode)
	}

	return nil
}
