package start

import (
	"github.com/dbidib/sendtome/features"

	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func Onstart(c tele.Context) error {
	return nil
}

func init() {
	features.RegisterFeature("/start", Onstart)
}
