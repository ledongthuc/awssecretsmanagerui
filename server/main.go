package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/browser"

	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
	"github.com/ledongthuc/awssecretsmanagerui/server/tray"
)

type arguments struct {
	Host     string
	Port     int
	Headless bool
}

func parseArguments() (a arguments) {
	flag.StringVar(&a.Host, "host", "localhost", "Host of service. Default is localhost")
	flag.IntVar(&a.Port, "port", 0, "Port of service. Default is random free port")
	flag.BoolVar(&a.Headless, "headless", false, "Disable UI app. Good to deploy/run from server")
	flag.Parse()

	if a.Port == 0 {
		var err error
		if a.Port, err = getFreePort(); err != nil {
			panic(err)
		}
	}

	return a
}

func main() {
	a := parseArguments()
	serverAddr := fmt.Sprintf("%s:%d", a.Host, a.Port)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.HideBanner = true
	routes.SetupRoutes(e)

	if a.Headless {
		if err := e.Start(serverAddr); err != nil {
			panic(err)
		}
	} else {
		go func() {
			if err := e.Start(serverAddr); err != nil {
				panic(err)
			}
		}()

		url := fmt.Sprintf("http://%s", serverAddr)
		browser.OpenURL(url)
		tray.Start(url)
	}
}

func composeServerAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func getFreePort() (int, error) {
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
