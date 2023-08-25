package basic

import (
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

// 代码来源：https://www.honeybadger.io/blog/golang-logging/

func TestLog(t *testing.T) {
	// 添加了时间戳
	log.Println("Hello World")
}

func TestLogFile(t *testing.T) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Hello World")
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	WarningLogger = log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestLogFormat(t *testing.T) {
	WarningLogger.Println("Warning")
	InfoLogger.Println("Info")
	ErrorLogger.Println("Error")
}

func TestLogrus(t *testing.T) {
	logrus.Println("Hello World")
}

func TestLogrusJson(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"foo": "foo",
		"bar": "bar",
	},
	).Info("Something happened")
}

func TestLevel(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 当设置为Debug级别的时候，才会答应该条消息
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Useful debugging information.")
	logrus.Info("Info information.")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
}
