package edd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type Signal struct {
	EmitterID           string                 `json:"emitterId"`
	StepNodeID          string                 `json:"stepNodeId"`
	Data                map[string]interface{} `json:"data"`
	EventFlowIdentifier string                 `json:"eventflowIdentifier"`
}

type EventFlowSignalRequest struct {
	RequestIdentifier uuid.UUID `json:"requestIdentifier"`
	Signals           []Signal  `json:"signals"`
}

type EngageDataDrive struct {
	apiKey  string
	baseUrl string
}

func NewEngageDataDrive(apiKey string) *EngageDataDrive {
	return &EngageDataDrive{
		apiKey:  apiKey,
		baseUrl: "https://edd.kursaha.com",
	}
}

func (e *EngageDataDrive) SendSignal(signals []Signal) error {
	url := e.baseUrl + "/api/event-flows/signal"

	payload := EventFlowSignalRequest{
		RequestIdentifier: uuid.New(),
		Signals:           signals,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode != 200 {
		return fmt.Errorf("signal request failed with status code: %d", resp.StatusCode)
	}

	return nil
}
