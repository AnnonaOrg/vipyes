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
	menu.Inline(menu.Row(btnAddToChat))
	c.Send("vip Yes! 自动删除群内非会员消息.", menu)
	return nil
}
