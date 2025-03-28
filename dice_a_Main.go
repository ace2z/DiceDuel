package main

import (
	// = = = = = Native Libraries

	. "local/CORE"

	//. "local/INGEST_ENGINE"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

var VERSION_NUM = ""

//  *************************** MAIN AREA ****************************

func main() {

	//1. Init the program (and the command line parameters)
	CLI_PARAM_INIT()
	MASTER_INIT("DiceDuel", VERSION_NUM)

	//2. Main Dice Engine
	if INPUT_RED_DICE != "" || INPUT_BLUE_DICE != "" {
		Dice_Engine_INIT(INPUT_RED_DICE, INPUT_BLUE_DICE)

		DO_EXIT("-silent")
	}

	input := Read_User_Input("Yo whats your name? ", BOLD_CYAN, BOLD_MAGENTA)

	C.Print("You Typed: **")
	M.Print(input)
	C.Println("** ")

	PressAny()

	DO_EXIT()
	// otherwise go into a loop
	for {
		ShowDice_HISTORY()
		//Predict_Engine() // Call this after history

		Dice_Engine_INIT("", "")
	}

} // end of main
