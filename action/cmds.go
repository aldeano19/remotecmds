package action

import (
	"bytes"
	"os/exec"
	"errors"
	"io"
	"os"
	"log"
	"strings"
	"github.com/b4b4r07/go-pipe"
)


const (
	UTCTime              Action = "utc-time"
	CPUUsage             Action = "cpu-usage"
	RAMAvailable         Action = "ram-available"
	DownloadUrlInto      Action = "download-url-into"
	Speak                Action = "speak"
	sendScreenshot       Action = "send-screenshot"
	triggerWebhook       Action = "trigger-webhook"
)


type Action string

//type CmdList []exec.Cmd

// RawCmd can be loaded as-is into exec.Cmd.
type RawCmd []string

// VarCmd implies that preprocessing needs to happen. Eg, string interpolation, etc.
// Any %s string within the command needs to be replaced by a value.
type VarCmd []string

//type BuildableCmd interface {
//	Build() exec.Cmd
//}

//var (
//	utcTimePipe  		CmdList = []exec.Cmd{*exec.Command("date")}
//	cpuUsagePipe 		CmdList = []exec.Cmd{
//							*exec.Command("top", "-l 1", "-n 0"),
//							*exec.Command("grep", "\"CPU usage\""),
//						}
//	memoryUsagePipe 	CmdList = []exec.Cmd{
//							*exec.Command("top", "-l 1", "-n 0"),
//							*exec.Command("grep", "\"PhysMem\""),
//						}
//	downloadUrlIntoPipe CmdList = []exec.Cmd{
//							*exec.Command("cd", "%s"),
//							*exec.Command("curl", "-O %s"),
//						}
//	speakPipe 			CmdList = []exec.Cmd{
//							*exec.Command("say", "%s"),
//						}
//)

var (
	dateRawCmd         RawCmd = []string{"date"}
	topRawCmd          RawCmd = []string{"top", "-l 1", "-n 0"}
	grepCPUUsageRawCmd RawCmd = []string{"grep", "\"CPU usage\""}
	grepMEMUsageRawCmd RawCmd = []string{"grep", "\"PhysMem\""}

	cdVarCmd   VarCmd = []string{"cd", "{}"}
	curlVarCmd VarCmd = []string{"curl", "-O", "{}"}
	sayVarCmd  VarCmd = []string{"say", "{}"}

)

func (cmd RawCmd) String() string {
	return strings.Join(cmd, " ")
}

func (cmd VarCmd) String() string {
	return strings.Join(cmd, " ")
}

func (cmd RawCmd) Build() *exec.Cmd{

	// TODO: Validation checks

	var execCmd *exec.Cmd
	if len(cmd) > 1{
		execCmd = exec.Command(cmd[0], cmd[1:]...)
	} else {
		execCmd = exec.Command(cmd[0])
	}

	return execCmd
}

func (cmd VarCmd) Build(args ...string) exec.Cmd {

	/*
		TODO:
			Validation checks:
			Mismatch len(args) and {} variables in cmd
	*/

	aIndex := 0
	for i, str := range cmd {
		if str == "{}" {
			cmd[i] = args[aIndex]
			aIndex++
		}
	}

	return *exec.Command(cmd[0], cmd[1:]...)
}
//
//var commandMap = map[Action]CmdList{
//	UTCTime: utcTimePipe, // date
//	CPUUsage:  cpuUsagePipe, // top -l 1 -n 0
//	RAMAvailable: memoryUsagePipe, // top -l 1 -s 0 | grep PhysMem
//	DownloadUrlInto: downloadUrlIntoPipe, // cd /some/path && curl -O SOME_URL
//	Speak: speakPipe, // say "something"
//}
//
//func (pipe CmdList) rawCmd(separator string) string  {
//	var buffer bytes.Buffer
//
//	for i, cmd := range pipe {
//		buffer.WriteString(strings.Join(cmd.Args, " "))
//		if i != len(pipe) - 1 {
//			buffer.WriteString(separator)
//		}
//	}
//
//	return buffer.String()
//}
//
//func (pipe CmdList) execOne() string {
//
//	bytes, err := pipe[0].CombinedOutput()
//
//	panicOnErr(err)
//
//	return string(bytes)
//}
//
//func (pipe CmdList) execMany() string {
//	r, w := io.Pipe()
//
//	pipe[0].Stdout = w
//	pipe[1].Stdin = r
//
//	pipe[0].Start()
//
//
//	pipe[1].Start()
//
//	pipe[0].Wait()
//	w.Close()
//
//
//	i := 2
//	for i < len(pipe) {
//		r, w = io.Pipe()
//		pipe[i-1].Stdout = w
//		pipe[i].Stdin = r
//
//		pipe[i].Start()
//
//		pipe[i-1].Wait()
//		w.Close()
//	}
//
//	var tmpBuf bytes.Buffer
//
//	pipe[i-1].Stdout = &tmpBuf
//
//	return tmpBuf.String()
//}
//
//func (pipe CmdList) execRaw() string {
//
//	commandString := pipe.rawCmd("|")
//
//	log.Printf("Executing rawCmd command = %s", commandString)
//	out, err := exec.Command("bash", "-c", commandString).Output()
//
//	panicOnErr(err)
//	return string(out)
//}
//
//func (pipe CmdList) execRawWithParams(action Action, args ...string ) string {
//	commandString := fmt.Sprintf(pipe.rawCmd("&"), args)
//
//	log.Printf("Executing rawCmd command = %s", commandString)
//	out, err := exec.Command("bash", "-c", commandString).Output()
//
//	panicOnErr(err)
//	return string(out)
//}


//func (pipe CmdList) Execute1() (string, error){
//	if pipe == nil || len(pipe) < 1 {
//		message := "Nothing to execute."
//
//		return "", errors.New(message)
//	}
//
//	if len(pipe) == 1 {
//		//log.Println("Only one cmd on pipe.")
//		return pipe.execOne(), nil
//	} else {
//		//log.Printf("%d cmds on pipe.", len(pipe))
//		return pipe.execRaw(), nil
//	}
//}


func (act Action) Process() (string, error) {

	switch act {
	case UTCTime:
		cmd := dateRawCmd.Build()

		output, err := cmd.CombinedOutput()
		panicOnErr(err)

		return string(output), nil

	case CPUUsage:
		var b bytes.Buffer

		log.Println(topRawCmd.String())
		log.Println(grepCPUUsageRawCmd.String())


		err := pipe.Command(&b, topRawCmd.Build(), grepCPUUsageRawCmd.Build())

		//err := pipe.Command(&b, exec.Command(topRawCmd[0], topRawCmd[1:]...), exec.Command(grepMEMUsageRawCmd[0], grepMEMUsageRawCmd[1:]...))

		log.Println("HEREE")
		panicOnErr(err)


		_, err = io.Copy(os.Stdout, &b)
		panicOnErr(err)



		return b.String(), nil

	case RAMAvailable:
		var b bytes.Buffer
		err := pipe.Command(&b, topRawCmd.Build(), grepMEMUsageRawCmd.Build())
		panicOnErr(err)

		return b.String(), nil
	default:
		return "", errors.New("bad action=" + string(act))
	}

}


func panicOnErr(err error)  {
	if err != nil {
		//log.Print("Error: ")
		//log.Fatal(err)
		panic(err)
	}
}
