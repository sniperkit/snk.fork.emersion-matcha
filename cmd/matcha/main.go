/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"flag"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/sniperkit/snk.fork.emersion-matcha"
)

func main() {
	addr := os.Getenv("ADDRESS")
	if port := os.Getenv("PORT"); port != "" {
		addr = "127.0.0.1:" + port
	}
	if addr == "" {
		addr = "127.0.0.1:8088"
	}

	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := matcha.New(e, dir); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(addr))
}
