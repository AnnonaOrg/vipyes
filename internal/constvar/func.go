package constvar

const (
	APP_ABOUT = "TG Vip Yes BOT."
)

func Version() string {
	return APP_VERSION
}

func About() string {
	text := APP_ABOUT + " v" + APP_VERSION
	return text
}
