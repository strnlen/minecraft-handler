package main

import (
    "bufio"
    "strings"
    "fmt"
    "os"
    "os/exec"
    "time"
)

const (
	kUNKNOWN_EVENT = iota
	kPLAYER_CHAT_EVENT = iota
    kPLAYER_COMMAND_EVENT = iota
	kSERVER_EVENT = iota
	kAUTHENTICATION_EVENT = iota
	kPLAYER_EVENT_JOINED = iota
    kPLAYER_EVENT_WAS = iota
    kPLAYER_EVENT_LOST = iota
    kPLAYER_EVENT_LEFT = iota
    kPLAYER_EVENT_FELL = iota
    kPLAYER_EVENT_DROWNED = iota
)

const (
    kPEVENT_JOINED_STR = "joined"
    kPEVENT_WAS_STR = "was"
    kPEVENT_LOST_STR = "lost"
    kPEVENT_LEFT_STR = "left"
    kPEVENT_FELL_STR = "fell"
    kPEVENT_DROWNED_STR = "drowned"
)

const (
    kPCOMMAND_TP_STR = "+tp"
    kPCOMMAND_TP_WAYPOINT_STR = "+tpway"
    kPCOMMAND_WAYPOINT_SAVE_STR = "+waypoint_save"
    kPCOMMAND_WAYPOINT_RESET_STR = "+waypoint_reset"
    kPCOMMAND_WEATHER = "+weather"
    kPCOMMAND_CLOCK = "+clock"
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
        if lineType == kPLAYER_COMMAND_EVENT {
            LineResolveCommand(incomingLine)

        }
	}

	if scanner.Err() != nil {
	    // handle error.
	}

}

func LineResolveCommand(line string) int {
    fmt.Println("LineResolveCommand")
    tokenizedString := strings.Split(line, " ")
    fmt.Println(tokenizedString)
    const playerCommandBase = 4

    switch tokenizedString[playerCommandBase] {
        case kPCOMMAND_TP_STR:
            fmt.Println("Teleport Player")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_TP_WAYPOINT_STR:
            fmt.Println("Teleport Waypoint")
            //Command & Waypoint Validation
            //Authentication Check

        case kPCOMMAND_WAYPOINT_SAVE_STR:
            fmt.Println("Waypoint Save")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_WAYPOINT_RESET_STR:
            fmt.Println("Waypoint Reset")
            //Authentication Check

        case kPCOMMAND_WEATHER:
            fmt.Println("Weather")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_CLOCK:
            fmt.Println("Clock")
            currentTime := time.Now()
            var currentTimeFormatted = currentTime.Format("01-02-2006 15:04:05")
            ServerAnnounce(currentTimeFormatted)

        default:
            fmt.Println("Player Unknown Command")

    }
    return 0

}

func LineDetermineType(line string) int {
    fmt.Println("LineDetermineType")
    tokenizedString := strings.Split(line, " ")
    fmt.Println(tokenizedString)
    //In User Authentication Events, [User is the second token.
    //In Server Events, thread/WARN]: is the third token.
    //In Player Chat, the forth token begins with a < character.
    //In Player Events, the fifth token is either [joined, was, lost, left, fell, drowned].

    if tokenizedString[1] == "[User" {
    	fmt.Println("User Authentication Event")
    	return kAUTHENTICATION_EVENT

    } else if tokenizedString[2] == "thread/WARN]:" {
    	fmt.Println("Server Event")
    	return kSERVER_EVENT

    } else if string([]rune(tokenizedString[3])[0]) == "<"{
        if string([]rune(tokenizedString[4])[0]) == "+"{
            fmt.Println("Player Chat Event")
            return kPLAYER_COMMAND_EVENT

        } else {
            fmt.Println("Player Chat Event")
            return kPLAYER_CHAT_EVENT

        }
    	
    }

    for e, token := range tokenizedString {    
        fmt.Println("Token Index: ", e, token)
    }

    //Check the string against the known player event types, return the type of event
    switch tokenizedString[4] {
        case kPEVENT_JOINED_STR:
            fmt.Println("Player Joined")
            return kPLAYER_EVENT_JOINED

        case kPEVENT_WAS_STR:
            fmt.Println("Player Was")
            return kPLAYER_EVENT_WAS

        case kPEVENT_LOST_STR:
            fmt.Println("Player Lost")
            return kPLAYER_EVENT_LOST

        case kPEVENT_LEFT_STR:
            fmt.Println("Player Left")
            return kPLAYER_EVENT_LEFT

        case kPEVENT_FELL_STR:
            fmt.Println("Player Fell")
            return kPLAYER_EVENT_FELL

        case kPEVENT_DROWNED_STR:
            fmt.Println("Player Drowned")
            return kPLAYER_EVENT_DROWNED
        default:
            fmt.Println("Player Unknown Event")

    }

    return kUNKNOWN_EVENT
}

func ServerWhisper(player string, message string){
    //tell
    var commandString  = "/tell " + player + "\\\"" + message + "\\\""
    fmt.Println("Server tell:", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerAnnounce(message string){
    //say
    var commandString  = "/say " + message + " "
    fmt.Println("Server say: ", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerTp(playerOne string, playerTwo string){
    //tp
    var commandString  = "/tp"
    fmt.Println("Server tp: ", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerSaveAll(){
    ///save-all [flush]
    var commandString  = "/save-all"
    fmt.Println("Server save-all: ", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerSaveOff(){
    ///save-off
    var commandString  = "/save-off"
    fmt.Println("Server save-off: ", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerSaveOn(){
    ///save-on
    var commandString  = "/save-on"
    fmt.Println("Server save-on: ", commandString)
    ScreenExecuteCommand(commandString)
}

func ServerJingle(player string){
    ///playsound minecraft:entity.wither.ambient voice @a 

}

func ServerWeather(weatherState string){
    //weather
    var commandString  = "/save-on"
    fmt.Println("Server save-on: ", commandString)
    ScreenExecuteCommand(commandString)
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
    //var fullCommand = "/usr/bin/screen -S MinecraftServer -p 0 -X stuff \"" + command + "^M\""
    screenCmd := exec.Command("bash", "-c", "/usr/bin/screen -S MinecraftServer -p 0 -X stuff \"" + command + "^M\"")
    screenCmdOut, err := screenCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println(string(screenCmdOut))

}