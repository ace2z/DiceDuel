package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"

	//. "local/INGEST_ENGINE"

	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

// VALUE string
var LETTERS = []string{"e", "r", "d", "v", "h", "s", "l"}

func Extra_KEYS_Handle_Engine(input string) bool {
	var flg_remove_last = false
	var flg_show_history = false

	var flg_save = false
	var flg_load = false

	for _, tmpl := range LETTERS {

		if strings.Contains(input, tmpl) {

			// If they want to edit, remove, delete the LAST entry in history
			if CONTAINS(input, "e") || CONTAINS(input, "r") || CONTAINS(input, "d") {
				flg_remove_last = true
			} else if CONTAINS(input, "v") || CONTAINS(input, "h") {
				flg_show_history = true

			} else if CONTAINS(input, "s") {
				flg_save = true
			} else if CONTAINS(input, "l") {
				flg_load = true
			}
			break
		}
	}

	// Deletes last item in HISTORY
	if flg_remove_last {
		if len(HISTORY) > 0 {
			HISTORY = HISTORY[:len(HISTORY)-1]
			Y.Println("")
			Y.Println("")
			MW.Print(" *** LAST ITEM REMOVED ***")
			Y.Println("")
			Y.Println("")

			return true
		}

		// othberwise, show a message
		Y.Println("")
		Y.Println("")
		M.Print(" *** Nothing In History!! ***")
		Y.Println("")
		Y.Println("")

		return true
	}

	// Shows the FULL history from the first roll of the game
	if flg_show_history {
		ShowDice_HISTORY(true)
		DONT_SHOW_HIST = true // a little override hack to prevent showing history twice (due to the way the for loop works)
		return true
	}

	// Saves a particular Game HISTORY to a file
	if flg_save {
		if len(HISTORY) > 0 {
			// Save the last game to a file
			Save_Game_ToFile()
			return true
		} else {
			Y.Println("")
			Y.Println("")
			MW.Print(" *** Nothing In History!! ***")
			Y.Println("")
			Y.Println("")

		}

		return true
	}

	if flg_load {
		// Load a game from a file
		// Implement the logic to load the game object from a file
		// This function should deserialize the game object from a file
		// and return it
		Load_Game_FromFile()
		return true
	}

	// otherwise, Delete the LAST item from history
	return false

}
