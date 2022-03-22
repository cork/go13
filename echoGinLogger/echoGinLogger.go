package echoGinLogger

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

type logData struct {
	StatusCode   int
	TotalLatency time.Duration
	ProxyLatency time.Duration
	SiteSource   string
	LocalIP      string
	RemoteIP     string
	TargetHost   string
	Method       string
	RequestURI   url.URL
}

// EchoLogger instances a Logger middleware that will write to Loggers
func EchoLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		start := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		timeStamp := time.Now()
		totalLatency := timeStamp.Sub(start)

		if totalLatency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			totalLatency = totalLatency - totalLatency%time.Second
		}

		ipAddr, _ := splitHostPort(req.RemoteAddr)
		statusColor := StatusCodeColor(res.Status)
		methodColor := MethodColor(req.Method)

		fmt.Printf("[echo] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n",
			start.Format("2006/01/02 - 15:04:05"),
			statusColor, res.Status, reset,
			totalLatency,
			ipAddr,
			methodColor, req.Method, reset,
			req.URL.String(),
		)

		return err
	}
}

func MethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

func StatusCodeColor(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// Copied from https://golang.org/src/net/url/url.go
// validOptionalPort reports whether port is either an empty string
// or matches /^:\d*$/
func validOptionalPort(port string) bool {
	if port == "" {
		return true
	}
	if port[0] != ':' {
		return false
	}
	for _, b := range port[1:] {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}

// splitHostPort separates host and port. If the port is not valid, it returns
// the entire input as host, and it doesn't check the validity of the host.
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
func splitHostPort(hostport string) (host, port string) {
	host = hostport

	colon := strings.LastIndexByte(host, ':')
	if colon != -1 && validOptionalPort(host[colon:]) {
		host, port = host[:colon], host[colon+1:]
	}

	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}

	return
}
