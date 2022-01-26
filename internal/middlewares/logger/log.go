package logger 

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net"
	"strings"
	"time"
)



func realIP(req *http.Request) string {
	ra := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		ra = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}
	return ra
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry := log.NewEntry(log.New())
		start := time.Now()

		if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
			entry = entry.WithField("requestId", reqID)
		}

		if remoteAddr := realIP(r); remoteAddr != "" {
			entry = entry.WithField("remoteAddr", remoteAddr)
		}

		entry.WithFields(log.Fields{
			"request": r.RequestURI,
			"method":  r.Method,
		}).Info("started handling request")

        next.ServeHTTP(w, r)

		latency := time.Since(start)

		entry.WithFields(log.Fields{
			"took":   latency,
		}).Info("completed handling request")
    })
}