package check_vip

import (
	"sync"

	"github.com/dbidib/sendtome/features"
	tele "gopkg.in/telebot.v3"
)

var syncMap sync.Map

func init() {
	features.RegisterFeature(tele.OnText, OnText)
	features.RegisterFeature(tele.OnPhoto, OnText)
	features.RegisterFeature(tele.OnAudio, OnText)
	features.RegisterFeature(tele.OnAnimation, OnText)
	features.RegisterFeature(tele.OnDocument, OnText)
	features.RegisterFeature(tele.OnVideo, OnText)
	features.RegisterFeature(tele.OnVoice, OnText)
}

// OnText
func OnText(c tele.Context) error {

	if c.Message().FromGroup() {
		if !c.Sender().IsPremium {
			if err := c.Delete(); err != nil {
				return c.Reply("警告:本群仅限会员发言")
			} else {
				return nil
			}
		}
	}

	return nil
}
