package EVENTS

import (
	// "flag"
	//"fmt"
	. "local/CORE"
	"sort"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	. "github.com/ace2z/GOGO/Gadgets"
	//	"github.com/fatih/color"
	// "github.com/charmbracelet/lipgloss"
)

var EVENT_LIST []EVENT_OBJ

// Event Names.. to make them consistent across functions that show or consume them
const (
	// RED_by_ONE  string = "RE*1"
	// BLUE_by_ONE string = "BL*1"
	COMMON_NUMS string = "COMNUM"
	RED_FULL    string = "RED_360"
	BLUE_FULL   string = "BLU_360"
	RED_1       string = "RED_1"
	BLUE_1      string = "BLU_1"
	RED_INC     string = "Red_INC"
	BLUE_INC    string = "Blu_INC"

	RED_HINUM  string = "reHI"
	BLUE_HINUM string = "blHI"
)

func populate_EVENTS_LIST() {

	EVENT_LIST = []EVENT_OBJ{

		// {RED_INC, 1, NULV, "", "", NULV, NULV, "", true, detect_INC_DROP, nil, LG_JUST_RED},
		// {BLUE_INC, 1, NULV, "", "", NULV, NULV, "", true, detect_INC_DROP, nil, LG_JUST_BLUE},

		{RED_FULL, 5, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_RED_360, ""},
		{BLUE_FULL, 5, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_BLUE_360, ""},

		{RED_1, 4, NULV, "", "", NULV, NULV, "", true, detect_HAVE_1, nil, LG_WHITE_RED_UNDY, ""},
		{BLUE_1, 4, NULV, "", "", NULV, NULV, "", true, detect_HAVE_1, nil, LG_LIGHT_on_BLUE_UNDY, ""},

		{RED_HINUM, 1, NULV, "", "", NULV, NULV, "", true, detect_HI_NUM, nil, LG_WHITE_ORANGE, ""},
		{BLUE_HINUM, 1, NULV, "", "", NULV, NULV, "", true, detect_HI_NUM, nil, LG_WHITE_PURPLE, ""},

		{COMMON_NUMS, 1, NULV, "", "", NULV, NULV, "", true, detect_COMMON_NUMS, nil, LG_COMMON_NUMS, ""},

		// {RED_by_ONE, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_WHITE_ORANGE},
		// {BLUE_by_ONE, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_WHITE_PURPLE},

	}

	// Now SORT the event list by priority
	sort.Slice(EVENT_LIST, func(a, b int) bool {
		return EVENT_LIST[a].PRIORITY > EVENT_LIST[b].PRIORITY
	})

}
