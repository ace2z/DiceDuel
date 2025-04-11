package EVENTS

import (
	// "flag"
	//"fmt"
	. "local/CORE"
	// . "github.com/ace2z/HP_COMMON"
	//"encoding/json"
	//. "github.com/ace2z/GOGO/Gadgets"
	//	"github.com/fatih/color"
	// "github.com/charmbracelet/lipgloss"
)

var EVENT_LIST []EVENT_OBJ

// Event Names.. to make them consistent across functions that show or consume them
const (
	RED_by_ONE  string = " RED*by*ONE "
	BLUE_by_ONE string = " BLU*by*ONE "
	RED_FULL    string = " RED_FULL "
	BLUE_FULL   string = " BLU_FULL "
	RED_1       string = " RED_1 "
	BLUE_1      string = " BLU_1 "
	RED_INC     string = "Red_INC"
	BLUE_INC    string = "Blu_INC"
)

func populate_EVENTS_LIST() {

	EVENT_LIST = []EVENT_OBJ{

		{RED_INC, 1, NULV, "", "", NULV, NULV, "", true, detect_INC_DROP, nil, LG_JUST_RED},
		{BLUE_INC, 1, NULV, "", "", NULV, NULV, "", true, detect_INC_DROP, nil, LG_JUST_BLUE},

		{RED_by_ONE, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_WHITE_ORANGE},
		{BLUE_by_ONE, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_WHITE_PURPLE},

		{RED_FULL, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_RED_RED},
		{BLUE_FULL, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_BLUE_BLUE},

		{RED_1, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_WHITE_RED},
		{BLUE_1, 1, NULV, "", "", NULV, NULV, "", true, Detect_diceSIDE_Probabilities, nil, LG_LIGHT_on_BLUE},
	}

}
