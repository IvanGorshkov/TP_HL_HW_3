package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/liderman/text-generator"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var hitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "hits",
})

func main() {

	if err := prometheus.Register(hitsTotal); err != nil {
		fmt.Println(err)
	}


	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler)
	e.GET("/metrics",  echo.WrapHandler(promhttp.Handler()))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}

func handler(c echo.Context) error {
	id := uuid.New()
	hitsTotal.Inc()
	tg := text_generator.New()
	template := "{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother}! {How are you|What's new with you}?"
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	return c.String(http.StatusOK, id.String() + " " + tg.Generate(template))
}
