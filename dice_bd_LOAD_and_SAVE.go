package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	. "local/DLOGIC"

	//"strings"
	. "github.com/ace2z/GOGO/Gadgets"
	. "github.com/ace2z/GOGO/Gadgets/FileOPS"
	"path/filepath"
	"time"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

var FILE_PREFIX = "SAVEGAME_"
var FILE_EXT = ".dat"

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

		SG_filename = FILE_PREFIX + year + "_" + month + "_" + day + "_" + weekd + "--" + hour + "_" + minute + FILE_EXT
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

type LOADFILE_OBJ struct {
	FULLPATH string
	NUM      int
}

func Load_Game_FromFile() {
	W.Println("")
	W.Println("")
	G.Println("Load from Saved Games: ")
	W.Println("")

	// Find all the files with SAVEGAME_ prefix
	files, err := filepath.Glob(FILE_PREFIX + "*" + FILE_EXT)

	// Iterate over the files and print their names
	if err != nil {
		M.Println("Error finding files:", err)
		return
	}
	if len(files) == 0 {
		M.Println("No files found with prefix", FILE_PREFIX)
		return
	}
	var ALL_FILES []LOADFILE_OBJ

	// Print the list of files
	for i, file := range files {
		W.Print(i, ": ")
		Y.Println(file)

		var fo LOADFILE_OBJ
		fo.FULLPATH = file
		fo.NUM = i
		ALL_FILES = append(ALL_FILES, fo)
	}

	// Ask the user to select a file
	C.Println("")
	menu_num := Read_User_Input("Select a file to load: ", BOLD_MAGENTA, BOLD_YELLOW, 2, nil, "-digits", NOEOL)
	mnum := STRING_to_INT(menu_num)

	// Now Load the selection they specified
	for _, x := range ALL_FILES {
		if x.NUM == mnum {
			// Load the game from the selected file
			// This function should deserialize the game object from a file
			// and return it
			// Y.Print("Loading from: ")
			// G.Println(x.FULLPATH)
			HISTORY = []GAME_OBJ{} // Clear the current history
			W.Println("")
			W.Println("")

			err := LOAD_Struct_from_FILE(x.FULLPATH, &HISTORY, false)
			if err != nil {
				M.Println(err)
			}
			W.Println("")
			W.Println("")
			//ShowDice_HISTORY(true)
			return
		}
	}

	//3. Ifw e get this far, means we DIDNT find the users selection
	M.Println("")
	M.Println("")
	MW.Println("ERROR: Invalid Selection")
	M.Println("")

}
