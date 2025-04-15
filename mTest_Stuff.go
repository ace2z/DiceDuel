package main

import (
	// = = = = = Native Libraries

	. "local/CORE"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"

	. "github.com/ace2z/GOGO/Gadgets/MDC"

	"github.com/charmbracelet/lipgloss"
)

type EVTMP_OBJ struct {
	NAME     string
	PRIORITY int // Used for Prediction/Analysis. This is the PRIORITY of the event. Higher Priority will take precedence

	LAST_SEEN            int    // This is the number of hands/games BACK when we last saw this EXACT event
	LAST_SEEN_WINNER     string // When Last Seen, this is the WINNER of the hand
	LAST_SEEN_NEXTWINNER string // When Last Seen, this is the WINNER of the NEXT hand that follows this one

	// Optional. If there is a value associated with this event, we save it here
	// (as well as the PREVIOUS value) for this in relation to the HAND)
	VAL       int
	PREVVAL   int
	SIGNATURE string // If applicable, this is a SIGNATURE of this event. Something that makes it searchable (not always unique)

	// this is the HANDLER function that does the LOGIC for detecting this particular event
	SHOW_ME bool                                                       // If FALSE, this event is hidden from the display
	Handler func(GM *HAND_OBJ, event EVENT_OBJ, HIST *[]HAND_OBJ) bool `json:"-"` // Cant save this to JSON
	// Color style for displaying the event
	COLOR     *color.Color   `json:"-"` // DONT  save this to JSON
	LG_COLOR  lipgloss.Style `json:"-"` // DONT  save this to JSON
	hiddenGUY string

	TestInner inner_OBJ
	COLOR_ID  string // This is the LIPGLOSS, ID of the color in the COLOR_MATRIX
}

type inner_OBJ struct {
	BADNAME string
	HIDEME  LG_STYLE_TYPE
}

func Test_Stuff() {

	// commit_ver, _, _ := betaRUN_COMMAND("git rev-parse --short=6 HEAD", SILENT, "-env|LIONO=Thundercats", "-env|BillyJean=MIKEJACK")
	// C.Print("Commit Vaer: ")
	// Y.Println(commit_ver)

	//betaRUN_COMMAND("env")
	RUN_COMMAND("env", "-env|LIONO=Thundercats", "-env|BillyJean=MIKEJACK")

	DO_EXIT()

	EV := EVTMP_OBJ{}
	EV.hiddenGUY = "I just cant get enough"

	myMap1 := make(map[string]interface{})
	myMap1["name"] = "John Doe"
	myMap1["age"] = 30
	myMap1["city"] = "New York"

	SHOW_STRUCT(EV, "LG_COLOR", "HIDEME", "-hide")
	SHOW_STRUCT(myMap1)

	C.Println("Testing new MDC method")
	var num = 3.14159
	test := MakeRound(num)

	W.Println("\n", num)
	C.Print("Rounded: ")
	Y.Println(test)

	var allowed_DICE = []string{"1", "2", "3", "4", "5", "6"}
	rres := Read_UserInput("RED Dice: ", allowed_DICE, "-min|2", "-max|2", MAGENTA, NOEOL)
	r2 := Read_UserInput("   BLUE Dice: ", allowed_DICE, "-min|2", "-max|2", CYAN)
	Y.Print("\nRed Dice: ")
	M.Println(rres)
	Y.Print("\nBLUE Dice: ")
	C.Println(r2)

} // end of main
