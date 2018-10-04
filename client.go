package main

import (
	"flag"
	"fmt"
	"os"
	"net"
	"bufio"
	"io"
)

var act string
var help bool

func main() {
	//flag.BoolVar(&utcTime, "utctime", false, "Print UTC time")
	flag.StringVar(&act, "a", "", "Print UTC time")
	flag.BoolVar(&help, "hh", false, "Print help")
	flag.Parse()

	helpStr := ` 
	UTCTime              aName = "utc-time"
	CPUUsage             aName = "cpu-usage"
	RAMAvailable         aName = "ram-available"`

	if help{
		fmt.Println(helpStr)
		os.Exit(0)
	}


	//out, err := action.Action{act, []string{}}.Process()
	//
	//if err != nil { panic(err) }
	//
	//
	//fmt.Println(out)

	conn, err := net.Dial("tcp", "127.0.0.1:60080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Sendind message=" + act)
	fmt.Fprintf(conn, act)

	fmt.Print("Waiting for response...")
	resp, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Print("Response: " + resp)
}
