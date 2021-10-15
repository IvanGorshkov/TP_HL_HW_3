package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/liderman/text-generator"
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

	return c.String(http.StatusOK, id.String() + " " + tg.Generate(template))
}
