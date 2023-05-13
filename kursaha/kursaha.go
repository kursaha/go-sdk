package kursaha

import (
	"github.com/kursaha/go-sdk/kursaha/edd"
	"github.com/kursaha/go-sdk/kursaha/mailkeets"
)

type Kursaha struct {
	Mk  *mailkeets.Mailkeets
	Edd *edd.EngageDataDrive
}

func NewKursaha(apiKey string) *Kursaha {
	return &Kursaha{
		Mk:  mailkeets.NewMailkeets(apiKey),
		Edd: edd.NewEngageDataDrive(apiKey),
	}
}
