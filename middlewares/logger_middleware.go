package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toufiq-austcse/logger"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		LogRequest(context)
		w := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = w
		// Do something here
		context.Next()
		LogResponse(context, w)
	}
}

func LogRequest(context *gin.Context) {
	logger.Log.WithFields(logrus.Fields{
		"header": context.Request.Header,
		"body":   context.Request.Body,
		"method": context.Request.Method,
		"url":    context.Request.URL.String(),
	}).Info("Request")
}

func LogResponse(context *gin.Context, writer *responseBodyWriter) {
	logger.Log.WithFields(logrus.Fields{
		"status": context.Writer.Status(),
		"header": context.Writer.Header(),
		"url":    context.Request.URL.String(),
		"body":   writer.body.String(),
	}).Info("Response")
}
