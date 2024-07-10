package start

import (
	"github.com/dbidib/sendtome/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/start", Onstart)
}

// Command: /start <PAYLOAD>
func Onstart(c tele.Context) error {

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnAddToChat := menu.URL("点击加我入群", "http://t.me/vipYesBot?startgroup=botstart")
	menu.Reply(menu.Row(btnAddToChat))
	c.Send("Vip Yes!", menu)
	return nil
}
