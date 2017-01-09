package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var logger *log.Logger

func init() {
	// Ignore timestamps.
	logger = log.New(os.Stdout, "", 0)
}

func usage(msg string) {
	logger.Printf("%s\n\n", msg)
	fmt.Fprintf(os.Stderr, "usage: %s version | create", path.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	var port = flag.String("port", "8880", "The port of the application. (string)")
	var host = flag.String("host", "127.0.0.1", "The ip / host of the application. (string)")
	var inside = flag.Bool("inside", true, "Binary runs inside a Kubernetes cluster. (bool)")
	flag.Parse() // parse the flags
	srv := newServer(*host, *port, newHandler(*inside))
	err := srv.start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't start web server")
		os.Exit(0)
	}

	//args := os.Args[1:]
	//if len(args) < 1 {
	//	usage("insufficient number of parameters")
	//}

	//opName := args[0]
	//var op operation

	//switch opName {
	//case "create":
	//default:
	//	usage(fmt.Sprintf("unknown operation: %s", opName))
	//}

}
