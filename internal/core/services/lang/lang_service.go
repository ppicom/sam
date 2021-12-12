package lang

import (
	"log"
	"time"
)

type LangService interface {
	WorkingDir(month time.Time) string
	MonthName(month time.Time) string
}

func NewLangService(language string) LangService {
	switch language {
	case "cat":
		return catLangService{}
	default:
		log.Fatalln("Not implemented")
	}
	return nil
}
