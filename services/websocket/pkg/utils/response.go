package utils

import (
	"github.com/gorilla/websocket"
)

// RespondWithError - response with error and close the connection
func RespondWithError(c *websocket.Conn, msg string, code int) {
	c.WriteMessage(code, []byte(msg))
	c.Close()
}

// RespondWithSuccess - response with success
func RespondWithSuccess(c *websocket.Conn, msg string, code int) {
	c.WriteMessage(code, []byte(msg))
}
