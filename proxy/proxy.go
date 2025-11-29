package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/helpers"
)

// ReverseProxy creates a reverse proxy to forward requests to target services
func ReverseProxy(target string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Parse target URL
		targetURL, err := url.Parse(target)
		if err != nil {
			helpers.SendResponse(ctx, http.StatusInternalServerError, "URL not parsed", "invalid target url")
			return
		}

		//create reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// customize director to modify the request
		proxy.Director = func(r *http.Request) {
			r.URL.Scheme = targetURL.Scheme
			r.URL.Host = targetURL.Host
			r.URL.Path = ctx.Param("proxyPath")
			r.Host = targetURL.Host

			// Forward headers
			r.Header = ctx.Request.Header
		}

		//Error handler
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("Proxy error: %v\n", err)
			helpers.SendError(ctx, http.StatusBadGateway, "Service error", "The is not available")
		}

		// serve the proxy request
		proxy.ServeHTTP(ctx.Writer, ctx.Request)

	}
}
