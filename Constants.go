package main

//Event Types
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

//Player Event Sentinels
const (
    kPEVENT_JOINED_STR = "joined"
    kPEVENT_WAS_STR = "was"
    kPEVENT_LOST_STR = "lost"
    kPEVENT_LEFT_STR = "left"
    kPEVENT_FELL_STR = "fell"
    kPEVENT_DROWNED_STR = "drowned"
)

//Command String Constants
const (
    kPCOMMAND_TP_STR = "+tp"
    kPCOMMAND_TP_WAYPOINT_STR = "+tpway"
    kPCOMMAND_WAYPOINT_SAVE_STR = "+waypoint_save"
    kPCOMMAND_WAYPOINT_SHOW_STR = "+waypoint_show"
    kPCOMMAND_WAYPOINT_RESET_STR = "+waypoint_reset"
    kPCOMMAND_WEATHER = "+weather"
    kPCOMMAND_HELP = "+help"
    kPCOMMAND_CLOCK = "+clock"
)