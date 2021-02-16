package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
	"github.com/webview/webview"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.HideBanner = true
	routes.SetupRoutes(e)
	serverAddr := composeServerAddr()
	go func() {
		e.Logger.Info(e.Start(serverAddr))
	}()

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("AWS Secrets manager")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://" + serverAddr)
	w.Run()
}

func composeServerAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		portN, err := GetFreePort()
		if err != nil {
			panic("Fail to aquire free port to start service")
		}
		port = strconv.Itoa(portN)
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
