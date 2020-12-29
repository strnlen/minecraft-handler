package main

import (
    "bufio"
    "strings"
    "fmt"
    "os"
    "os/exec"
)

const (
	kUNKNOWN_EVENT = iota
	kPLAYER_CHAT_EVENT = iota
	kSERVER_EVENT = iota
	kAUTHENTICATION_EVENT = iota
	kPLAYER_EVENT = iota
)

func main() {
	fmt.Println("< Minecraft Handler >")
	commandLineArguments := os.Args
	fmt.Println("Command Line Arguments: ", commandLineArguments)
	//whiteListPath := os.Args[1]
	//fmt.Println("Path: " + whiteListPath)

	//waypointsPath := os.Args[2]
	//fmt.Println("Path: " + waypointsPath)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var incomingLine = scanner.Text()
	    fmt.Println("Incoming Line: ", incomingLine)
	    var lineType = LineDetermineType(incomingLine)
	    fmt.Println("This line is of type: ", lineType)
	}

	if scanner.Err() != nil {
	    // handle error.
	}

}

func LineDetermineType(line string) int {
    s := strings.Split(line, " ")
    fmt.Println(s)
    //In User Authentication Events, [User is the second token.
    //In Server Events, thread/WARN]: is the third token.
    //In Player Chat, the forth token begins with a < character.
    //In Player Events, the fifth token is either [joined, was, lost, left, fell, drowned].

    if s[1] == "[User" {
    	fmt.Println("User Authentication Event")
    	return kAUTHENTICATION_EVENT

    } else if s[2] == "thread/WARN]:" {
    	fmt.Println("Server Event")
    	return kSERVER_EVENT

    } else if string([]rune(s[3])[0]) == "<"{
    	fmt.Println("Player Chat Event")
    	return kPLAYER_CHAT_EVENT
    	
    }
    return kUNKNOWN_EVENT
}

func ScreenListSessions(){
	var fullCommand = "screen -ls"
    screenCommand := exec.Command(fullCommand)
    screenSessionsRaw, err := screenCommand.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> screen -ls")
    var screenSessionsString = string(screenSessionsRaw)
    screenSessionsLines := strings.Split(screenSessionsString, "/n")
	fmt.Println(screenSessionsLines)

}

func ScreenExecuteCommand(command string){
	var fullCommand = "screen -S MinecraftServer -p 0 -X stuff \"" + command + "^M\""
    stuffCommand := exec.Command(fullCommand)
    stuffOut, err := stuffCommand.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> screen stuff")
    fmt.Println(string(stuffOut))

}