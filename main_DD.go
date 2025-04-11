package main

import (
	// = = = = = Native Libraries

	. "local/CORE"
	. "local/DICEY"

	//	"time"

	//"fmt"
	//"strings"
	//"text/template"

	//"reflect"
	//"bytes"
	//"encoding/binary"
	//"encoding/gob"
	//"reflect"
	//"github.com/qdm12/reprint"
	//. "local/INGEST_ENGINE"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/jinzhu/copier"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
	//"os"
)

var VERSION_NUM = ""

//  *************************** MAIN AREA ****************************

func main() {

	//1. Init the program (and the command line parameters)
	CLI_PARAM_INIT()
	MASTER_INIT("DiceDuel", VERSION_NUM)

	//2. Main Dice Engine
	if INPUT_RED_DICE != "" && INPUT_BLUE_DICE != "" {
		Dice_Engine_INIT(INPUT_RED_DICE, INPUT_BLUE_DICE)

		DO_EXIT("-silent")
	}

	// otherwise go into a loop
	for {

		//1. Main Dice Engine
		Dice_Engine_INIT("", "")

		//2. Show the Dice History  (last 10 hands)
		ShowDice_HISTORY(false)
	}

} // end of main
