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

	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
	"github.com/ledongthuc/awssecretsmanagerui/server/wrapper"
)

//go:embed static/*
var staticResources embed.FS
var DefaultPort = "30301"

type params struct {
	Host string
	Port string
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
	go func() {
		if err := e.Start(serverAddr); err != nil {
			panic(err)
		}
	}()

	wrapper.RunWrapperUI(serverAddr)
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
		if !isPortOpen(DefaultPort) {
			p.Port = DefaultPort
		} else {
			port, err := getFreePort()
			if err != nil {
				log.Warn("Can't generate dynamic port", err)
				p.Port = DefaultPort
			} else {
				p.Port = strconv.FormatInt(int64(port), 10)
			}
		}
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

func isPortOpen(port string) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), 1*time.Second)
	if conn != nil {
		conn.Close()
		return true
	}
	return false
}
