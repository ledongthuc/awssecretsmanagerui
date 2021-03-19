package main

import (
	"embed"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/browser"

	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
	"github.com/ledongthuc/awssecretsmanagerui/server/tray"
)

//go:embed static/*
var staticResources embed.FS

type params struct {
	Host     string
	Port     string
	Headless bool
}

func main() {
	p := parseParams()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.HideBanner = true
	routes.SetupRoutes(e, staticResources)

	serverAddr := fmt.Sprintf("%s:%s", p.Host, p.Port)
	if p.Headless {
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
		tray.Start(url, staticResources)
	}
}

func parseParams() params {
	p := params{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	if p.Host == "" {
		p.Host = "localhost"
	}
	if p.Port == "" {
		if !isPortOpen(3000) {
			p.Port = "3000"
		} else {
			port, err := getFreePort()
			if err != nil {
				log.Warn("Can't generate dynamic port", err)
				p.Port = "3000"
			} else {
				p.Port = strconv.FormatInt(int64(port), 10)
			}
		}
	}
	if os.Getenv("HEADLESS") != "" {
		p.Headless, _ = strconv.ParseBool(os.Getenv("HEADLESS"))
	}
	return p
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

func isPortOpen(port int) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", strconv.FormatInt(int64(port), 10)), 1*time.Second)
	if conn != nil {
		conn.Close()
		return true
	}
	return false
}
