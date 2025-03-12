package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/izya4ka/notes-web/gateway/gateerrors"
	"github.com/izya4ka/notes-web/gateway/util"
	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context, target string) error {
	route := target + c.Request().RequestURI
	route = strings.Replace(route, "*", c.Param("*"), 1)

	log.SetPrefix("[INFO] ")
	log.Printf("http://gateway:%s%s -> %s", os.Getenv("GATEWAY_PORT"), c.Request().RequestURI, route)
	log.SetPrefix("")

	req, err := http.NewRequest(c.Request().Method, route, c.Request().Body)
	if err != nil {
		return util.SendErrorResponse(c, gateerrors.ErrInternal)
	}

	req.Header = c.Request().Header

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return util.SendErrorResponse(c, gateerrors.ErrInternal)
	}

	for k, v := range resp.Header {
		c.Response().Header().Set(k, v[0])
	}
	c.Response().WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}
