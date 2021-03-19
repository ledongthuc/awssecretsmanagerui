package tray

import (
	"embed"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/pkg/browser"
)

var url string
var icon []byte

func Start(u string, staticResources embed.FS) {
	url = u
	icon, _ = staticResources.ReadFile("static/icons/favicon-16x16.png")

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTooltip("AWS Secrets Mananger UI")

	mOpen := systray.AddMenuItem("New window", "New App window")
	mCopyURL := systray.AddMenuItem("Copy URL", "Copy URL")
	mQuit := systray.AddMenuItem("Close", "Close")

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		case <-mOpen.ClickedCh:
			browser.OpenURL(url)
			return
		case <-mCopyURL.ClickedCh:
			clipboard.WriteAll(url)
			return
		}
	}
}

func onExit() {
	// clean up here
}
