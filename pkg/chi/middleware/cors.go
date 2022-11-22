package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range r.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("content-type", "application/json")
		}
		//
		next.ServeHTTP(w, r)
	})
}
