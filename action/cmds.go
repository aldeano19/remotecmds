package action

import (
	"bytes"
	"os/exec"
	"errors"
	"github.com/aldeano19/go-pipe"
	"fmt"
	"strings"
)


type Action struct {
	Name *string
	Params *[]string
}



var (
	UTCTime          = "utc-time"
	CPUUsage         = "cpu-usage"
	RAMAvailable     = "ram-available"

	date	    	     = exec.Command("date")
	top     	         = exec.Command("top", "-l 1", "-n 0")
	grepCPUUsage  		 = exec.Command("grep", "CPU usage")
	grepMEMUsage  		 = exec.Command("grep", "PhysMem")

	cd 					 = exec.Command("cd")
	curl 				 = exec.Command("curl")
	say					 = exec.Command("say")
)


func (act Action) Process() (string, error) {
	var b bytes.Buffer

	fmt.Printf("%d:::THE NAME::: %s\n", len(*act.Name), act.Name )
	fmt.Printf("%d :::UTCTime::: %s\n", len(UTCTime)  , UTCTime)


	//TODO: why are this 2 strings not EQUAL ?????
	switch strings.Trim(*act.Name, " ") {
	case UTCTime:
		fmt.Println("HIII")
		output, err := date.CombinedOutput()
		panicOnErr(err)

		return string(output), nil

	case CPUUsage:
		err := pipe.Command(&b, exec.Command("top", "-l 1", "-n 0"), exec.Command("grep", "PhysMem"))
		panicOnErr(err)
		return b.String(), nil

	case RAMAvailable:
		err := pipe.Command(&b, top, grepMEMUsage)

		panicOnErr(err)
		return b.String(), nil

	default:
		return "", errors.New("bad action=" + string(*act.Name))
	}
}

func panicOnErr(err error)  {
	if err != nil {
		//log.Print("Error: ")
		//log.Fatal(err)
		panic(err)
	}
}
