package tray

import (
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/pkg/browser"
)

var url string

func Start(u string) {
	url = u
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
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
