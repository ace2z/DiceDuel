package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/INGEST_ENGINE"

	//"strings"
	"time"

	. "github.com/ace2z/GOGO/Gadgets"
	. "github.com/ace2z/GOGO/Gadgets/FileOPS"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

var SG_filename string

func Save_Game_ToFile() bool {

	// Generate a filename if we havent already. We do this once per game session
	if SG_filename == "" {
		var to = time.Now()

		year := INT_to_STRING(to.Year())
		month := INT_to_STRING(int(to.Month()))
		day := INT_to_STRING(to.Day())
		hour := INT_to_STRING(to.Hour())
		minute := INT_to_STRING(to.Minute())

		weekd := to.Weekday().String()

		weekd = weekd[0:3]

		SG_filename = "SAVEGAME_" + year + "_" + month + "_" + day + "_" + weekd + "--" + hour + "_" + minute + ".dat"
	}

	// Check if the file already exists
	if FILE_EXISTS(SG_filename) {
		// If file exists, we always overwrite it
		REMOVE_FILE(SG_filename)
	}

	W.Println("")
	W.Println("")
	G.Print(" *** SAVING game ***")
	SAVE_Struct_2_DISK(SG_filename, HISTORY)

	W.Println("")
	W.Println("")

	// Save the game to a file
	// This function should serialize the game object and save it to a file
	// Implement the logic to save the game object to a file
	return true
}

func Load_Game_FromFile() {
	G.Println("Loading game from FILE TBD...")
}
