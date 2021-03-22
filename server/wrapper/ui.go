package wrapper

import (
	"github.com/webview/webview"
)

func RunWrapperUI(address string) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("AWS Secrets Manager UI")
	w.SetSize(1250, 1000, webview.HintNone)
	w.Navigate("http://" + address)
	w.Run()
}
