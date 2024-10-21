package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetupLogger(printPretty bool, disableHtmlEscape bool) {
	Log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:       printPretty,
		DisableHTMLEscape: disableHtmlEscape,
	})
}
