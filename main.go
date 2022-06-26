package main

import (
	"flag"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// Define global variables.
var (
	log     *logrus.Logger
	textLog *logrus.Logger
)

func init() {
	// Initialize the Loggers.
	log = logrus.New()
	textLog = logrus.New()

	// Configure the Loggers.
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	textLog.SetFormatter(&logrus.TextFormatter{})
	textLog.SetOutput(os.Stdout)
}

func expandPath(destination *string) string {
	// Check if path is absolute and return if it is.
	if path.IsAbs(*destination) {
		return *destination
	}

	// Get the working directory.
	wd, err := os.Getwd()
	if err != nil {
		textLog.Fatal(err)
	}

	return path.Join(wd, *destination)
}

func main() {
	// Define the command-line flags.
	listenPort := flag.Int("listenport", 8888, "Specify the port the proxy will listen on")
	listenAddress := flag.String("listenaddress", "0.0.0.0", "Specify the address the proxy will listen on")
	connectPort := flag.Int("connectport", 80, "Specify the port the proxy will connect to")
	connectAddress := flag.String("connectaddress", "127.0.0.1", "Specify the address the proxy will connect to")
	flagDirectory := flag.String("logdir", "log", "Specify the directory to store logs in")

	// Parse the command-line flags.
	flag.Parse()

	// Make sure the log directory exists.
	logDirectory := expandPath(flagDirectory)
	err := os.MkdirAll(logDirectory, 0755)
	if err != nil {
		textLog.Fatal(err)
	}

	// Welcome banner.
	textLog.Info("Starting TCP Proxy...")
	textLog.Infof("Listening on %s:%d", *listenAddress, *listenPort)
	textLog.Infof("Connecting to %s:%d", *connectAddress, *connectPort)
	textLog.Infof("Logging to %s", logDirectory)

}
