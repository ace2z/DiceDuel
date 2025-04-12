package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"
	// . "local/EVENTS"
	// . "local/META"

	. "github.com/ace2z/GOGO/Gadgets"
)

/*
Processes the red and blue dice input. This can come from:
* the CLI by way of --red and --blue
* Realtime user input
* or from a saved game file (if you load one)
*/
func Process_Dice_Value_INPUT(red_dice string, blue_dice string) HAND_OBJ {
	// Error handling. make sure both red and blue_dice have 2 digits
	if len(red_dice) != 2 || len(blue_dice) != 2 {
		Y.Println("")
		Y.Println("  *** ERROR: Not enough Dice values specified! ***")
		return HAND_OBJ{}
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
	var HND HAND_OBJ
	HND.RED_A = red_a_int
	HND.RED_B = red_b_int
	HND.BLUE_A = blue_a_int
	HND.BLUE_B = blue_b_int
	HND.HID = "pending"

	return HND
}
