package handler

import (
	"fmt"

	"net/http"
	"app/internal/broker"
	"github.com/labstack/echo/v4"
)


func StreamEvent(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := c.Response().Writer.(http.Flusher)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Streaming unsupported")
	}
	
	ch := make(chan string, 1)
	broker.Brokerk.AddClient(ch)
	defer broker.Brokerk.RemoveClient(ch)
	
	ctx := c.Request().Context()

	for {
		select {
		case <-ctx.Done():
			return nil
		case payload := <-ch:
			fmt.Fprintf(c.Response().Writer, "event: update\ndata: %s\n\n", payload)
			flusher.Flush()
		}
	}
}

