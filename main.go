/*
 * traPCollection API
 *
 * traPCollectionのAPI
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/comail/colog"
	"github.com/labstack/echo-contrib/session"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"sample/model"
	"sample/openapi"
	"sample/router"
	sess "sample/session"
)

func main() {
	log.Printf("Server started")
	env := os.Getenv("COLLECTION_ENV")

	db, err := model.EstablishDB()
	if err != nil {
		panic(err)
	}

	sess, err := sess.NewSession(db)
	if err != nil {
		panic(fmt.Errorf("Failed In Session Constructor: %w", err))
	}

	e := echo.New()
	e.Use(session.Middleware(sess.Store()))
	e.Use(middleware.Recover())

	if env == "development" || env == "mock" {
		colog.SetMinLevel(colog.LDebug)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: true,
			Flag:   log.Ldate | log.Ltime | log.Lshortfile,
		})

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}` + "\n",
		}))
		e.Use(middleware.Logger())
	} else {
		colog.SetMinLevel(colog.LError)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: false,
			Flag:   log.Ldate | log.Ltime | log.Lshortfile,
		})
	}

	colog.Register()

	clientID := os.Getenv("CLIENT_ID")
	if len(clientID) == 0 {
		panic(errors.New("ENV CLIENT_ID IS NULL"))
	}
	clientSecret := os.Getenv("CLIENT_SECRET")
	if len(clientSecret) == 0 {
    panic(errors.New("ENV CLIENT_SECRET IS NULL"))
	}

	api, err := router.NewAPI(sess)
	if err != nil {
		panic(err)
	}

	openapi.SetupRouting(e, api)

	e.Start(os.Getenv("PORT"))
}
