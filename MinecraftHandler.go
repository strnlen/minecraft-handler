package main

import (
    "bufio"
    "strings"
    "fmt"
    "os"
    "os/exec"
    "time"
)

/* Mongodb integration for saving Waypoints and other data. Break Server and Player command
/ parsers into seperate Go files for organization sake. Online time monitoring.
*/

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
	    fmt.Println("Incoming Line:", incomingLine)
	    var lineType = LineDetermineType(incomingLine)
	    fmt.Println("This line is of type:", lineType)
        if lineType == kPLAYER_COMMAND_EVENT {
            LineResolveCommand(incomingLine)

        }
	}

	if scanner.Err() != nil {
	    // handle error.
	}

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
        fmt.Println("Token Index:", e, token)
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

func PlayerNameClean(playerName string) string {
    var playerNameReplacer = strings.NewReplacer("<", "", ">", "", " ", "")
    playerName = playerNameReplacer.Replace(playerName)
    return playerName

}

func LineResolveCommand(line string) int {
    fmt.Println("LineResolveCommand")
    tokenizedString := strings.Split(line, " ")
    fmt.Println(tokenizedString)
    
    const playerCommandIndex = 4
    const playerNameIndex = 3
    
    var playerName = PlayerNameClean(tokenizedString[playerNameIndex])
    fmt.Println("Command Originating Player:", playerName)
    
    var tokensCount = len(tokenizedString)
    fmt.Println("Command Tokens:", tokensCount)

    switch tokenizedString[playerCommandIndex] {
        case kPCOMMAND_TP_STR:
            fmt.Println("Teleport Player")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_TP_WAYPOINT_STR:
            fmt.Println("Teleport Waypoint")
            //Command & Waypoint Validation
            //if tokensCount != 
            //Authentication Check

        case kPCOMMAND_WAYPOINT_SAVE_STR:
            fmt.Println("Waypoint Save")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_WAYPOINT_SHOW_STR:
            fmt.Println("Waypoint Show")
            //Authentication Check

        case kPCOMMAND_WAYPOINT_RESET_STR:
            fmt.Println("Waypoint Reset")
            //Authentication Check

        case kPCOMMAND_WEATHER:
            fmt.Println("Weather")
            //Command Validation
            //Authentication Check

        case kPCOMMAND_HELP:
            fmt.Println("Help")
            //No validation or authentication required for this command.
            CommandHelpHandler(playerName)

        case kPCOMMAND_CLOCK:
            fmt.Println("Clock")
            //No validation or authentication required for this command.
            CommandClockHandler()

        default:
            fmt.Println("Player Unknown Command")

    }
    return 0

}

func CommandWeatherHandler(){

}

func CommandHelpHandler(player string){
    //ServerAnnounce(player)
}

func CommandClockHandler(){
    currentTime := time.Now()
    var currentTimeFormatted = currentTime.Format("01-02-2006 15:04:05")
    ServerAnnounce(currentTimeFormatted)
}

func ServerWhisper(player string, message string){
    //tell
    var commandString  = "/tell " + player + "\\\"" + message + "\\\""
    ScreenExecuteCommand(commandString)
}

func ServerAnnounce(message string){
    //say
    var commandString  = "/say " + message + " "
    ScreenExecuteCommand(commandString)
}

func ServerTp(playerOne string, playerTwo string){
    //tp
    var commandString  = "/tp"
    ScreenExecuteCommand(commandString)
}

func ServerSaveAll(){
    ///save-all [flush]
    var commandString  = "/save-all"
    ScreenExecuteCommand(commandString)
}

func ServerSaveOff(){
    ///save-off
    var commandString  = "/save-off"
    ScreenExecuteCommand(commandString)
}

func ServerSaveOn(){
    ///save-on
    var commandString  = "/save-on"
    ScreenExecuteCommand(commandString)
}

func ServerJingle(player string){
    ///playsound minecraft:entity.wither.ambient voice @a 

}

func ServerWeather(weatherState string){
    //weather
    var commandString  = "/save-on"
    ScreenExecuteCommand(commandString)
}

func ScreenListSessions(){
    screenSessions := exec.Command("bash", "-c", "/usr/bin/screen -ls")
    screenSessionsOut, err := screenSessions.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println(string(screenSessionsOut))
    var screenSessionsString = string(screenSessionsOut)
    screenSessionsLines := strings.Split(screenSessionsString, "/n")
    fmt.Println(screenSessionsLines)

}

func ScreenExecuteCommand(command string){
    screenCmd := exec.Command("bash", "-c", "/usr/bin/screen -S MinecraftServer -p 0 -X stuff \"" + command + "^M\"")
    screenCmdOut, err := screenCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println(string(screenCmdOut))

}