package main

import (
	"flag"
	"local/remotecmds/action"
	"fmt"
)

var act string

func main() {
	//flag.BoolVar(&utcTime, "utctime", false, "Print UTC time")
	flag.StringVar(&act, "action", "", "Print UTC time")
	flag.Parse()

	// TODO: If action requires params, cpture params and use ProcessWithParamas() instead
	out, err := action.Action(act).Process()

	if err != nil { panic(err) }


	fmt.Println(out)

	//conn, err := net.Dial("tcp", "127.0.0.1:60080")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Sendind message=" + *msg)
	//fmt.Fprintf(conn, *msg)
	//
	//fmt.Print("Waiting for response...")
	//resp, err := bufio.NewReader(conn).ReadString('\n')
	//
	//if err != nil && err != io.EOF {
	//	panic(err)
	//}
	//
	//fmt.Print("Response: " + resp)
}
