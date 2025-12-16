package automation

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Run() {
	log.Println("[INFO] Launching browser with leakless DISABLED")

	// IMPORTANT: Leakless(false) disables the helper binary
	l := launcher.New().
		Leakless(false).
		Headless(false)

	// Optional but recommended: use installed Chrome
	// l = l.Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe")

	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://www.google.com")
	page.MustWaitLoad()

	log.Println("[INFO] Browser launched successfully without leakless")
}
