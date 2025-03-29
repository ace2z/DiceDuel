package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	. "local/DLOGIC"

	//	"bufio"
	//"fmt"
	//"os"
	//"unicode"
	//"strings"
	//"os/exec"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/fatih/color"
	//"golang.org/x/term"
	//"github.com/inancgumus/screen"
	//tea "github.com/charmbracelet/bubbletea"
)

var HISTORY []GAME_OBJ

func Process_Dice_Value_INPUT(red_dice string, blue_dice string) {
	// Error handling. make sure both red and blue_dice have 2 digits
	if len(red_dice) != 2 || len(blue_dice) != 2 {
		Y.Println("")
		Y.Println("  *** ERROR: Not enough Dice values specified! ***")
		return
	}

	// 1. Extract the first and second numbers from red_dice...and blue dice
	red_b := red_dice[0:1]
	red_a := red_dice[1:2]

	blue_b := blue_dice[0:1]
	blue_a := blue_dice[1:2]

	// Convert to integers
	red_a_int := STRING_to_INT(red_a)
	red_b_int := STRING_to_INT(red_b)
	blue_a_int := STRING_to_INT(blue_a)
	blue_b_int := STRING_to_INT(blue_b)

	/*  //*** DEBUG
	M.Print("     ***  RED: ")
	W.Print("a:")
	M.Print(red_a_int)
	W.Print(" --> ")
	W.Print("b:")
	M.Println(red_b_int)


	C.Print("     *** BLUE: ")
	W.Print("a:")
	C.Print(blue_a_int)
	W.Print(" --> ")
	W.Print("b:")
	C.Println(blue_b_int)
	W.Println("")

	*/

	//3. Create a new GAME_OBJ
	var GM GAME_OBJ
	GM.RED_A = red_a_int
	GM.RED_B = red_b_int
	GM.BLUE_A = blue_a_int
	GM.BLUE_B = blue_b_int

	// = = = =
	//4. Process the Dice Logic EVENTS (that are enabled) based on the dice values we got
	//GM = DL_Events_Engine(GM)
	DL_Events_Engine(&GM)

	//5. finally add the GAME (with all the events that were detected) ...to the history
	HISTORY = append(HISTORY, GM)
}

var also_allowed = []string{"e", "r", "d", "v", "h", "s", "l"}

func Dice_Engine_INIT(INPUT_RED_DICE string, INPUT_BLUE_DICE string) {

	var red_dice = INPUT_RED_DICE
	var blue_dice = INPUT_BLUE_DICE

	var prompt_user = false

	if red_dice == "" || blue_dice == "" {
		prompt_user = true
	}

	if prompt_user {

		// If we were NOT provided values, then we need to prompt the user
		Y.Println("")
		Y.Print(" Enter numeric dice values")
		W.Print(" for LAST 2 Games ")
		Y.Println("...do RED first ")

		tmp_red := Read_User_Input("      RED DICE: ", BOLD_MAGENTA, BOLD_YELLOW, 2, also_allowed, "-digits", NOEOL)
		if tmp_red == "" || Extra_KEYS_Handle_Engine(tmp_red) {

			return
		}
		tmp_blue := Read_User_Input("     BLUE DICE: ", BOLD_CYAN, BOLD_WHITE, 2, also_allowed, "-digits")
		if tmp_blue == "" || Extra_KEYS_Handle_Engine(tmp_blue) {
			return
		}

		W.Println("")

		red_dice = tmp_red
		blue_dice = tmp_blue

	}

	//5. Process the Dice Input
	Process_Dice_Value_INPUT(red_dice, blue_dice)

} // end of main
