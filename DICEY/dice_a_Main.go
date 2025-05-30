package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"
	// . "local/EVENTS"
	// . "local/META"

	. "github.com/ace2z/GOGO/Gadgets"
)

// func Process_Dice_Value_INPUT(red_dice string, blue_dice string) HAND_OBJ {
// 	// Error handling. make sure both red and blue_dice have 2 digits
// 	if len(red_dice) != 2 || len(blue_dice) != 2 {
// 		Y.Println("")
// 		Y.Println("  *** ERROR: Not enough Dice values specified! ***")
// 		return HAND_OBJ{}
// 	}

// 	// 1. Extract the first and second numbers from red_dice...and blue dice
// 	red_b := red_dice[0:1]
// 	red_a := red_dice[1:2]

// 	blue_b := blue_dice[0:1]
// 	blue_a := blue_dice[1:2]

// 	// Convert to integers
// 	red_a_int := STRING_to_INT(red_a)
// 	red_b_int := STRING_to_INT(red_b)
// 	blue_a_int := STRING_to_INT(blue_a)
// 	blue_b_int := STRING_to_INT(blue_b)

// 	/*  //*** DEBUG
// 	M.Print("     ***  RED: ")
// 	W.Print("a:")
// 	M.Print(red_a_int)
// 	W.Print(" --> ")
// 	W.Print("b:")
// 	M.Println(red_b_int)

// 	C.Print("     *** BLUE: ")
// 	W.Print("a:")
// 	C.Print(blue_a_int)
// 	W.Print(" --> ")
// 	W.Print("b:")
// 	C.Println(blue_b_int)
// 	W.Println("")

// 	*/

// 	//3. Create a new GAME_OBJ
// 	var HND HAND_OBJ
// 	HND.RED_A = red_a_int
// 	HND.RED_B = red_b_int
// 	HND.BLUE_A = blue_a_int
// 	HND.BLUE_B = blue_b_int
// 	HND.HID = "pending"

// 	return HND
// }

// var also_allowed = []string{"e", "r", "d", "v", "h", "s", "l"}

// // Eevery game Session will have a unique ID
// // if you for some reason need to save the games to a database
// // You can use this to ensure each Game and HANDS within it...are unique
// var TMP_GAME_SESS = "8675309"

// func Dice_Engine(INPUT_RED_DICE string, INPUT_BLUE_DICE string) {

// 	var red_dice = INPUT_RED_DICE
// 	var blue_dice = INPUT_BLUE_DICE

// 	var prompt_user = false

// 	if red_dice == "" || blue_dice == "" {
// 		prompt_user = true
// 	}

// 	if prompt_user {

// 		// If we were NOT provided values, then we need to prompt the user
// 		Y.Println("")
// 		Y.Print(" Enter numeric dice values")
// 		W.Print(" for LAST 2 Games ")
// 		Y.Print("...do ")
// 		M.Print("RED ")
// 		Y.Println("first ")

// 		tmp_red := Read_User_Input("      RED DICE: ", BOLD_MAGENTA, BOLD_YELLOW, 2, also_allowed, "-digits", NOEOL)
// 		if tmp_red == "" || Extra_KEYS_Handle_Engine(tmp_red) {
// 			W.Println("")
// 			return
// 		}
// 		tmp_blue := Read_User_Input("     BLUE DICE: ", BOLD_CYAN, BOLD_WHITE, 2, also_allowed, "-digits")
// 		if tmp_blue == "" || Extra_KEYS_Handle_Engine(tmp_blue) {
// 			W.Println("")
// 			return
// 		}

// 		red_dice = tmp_red
// 		blue_dice = tmp_blue

// 		W.Println("")
// 	}

// 	//5. Process the Dice Input. Get the Hand OBJECT
// 	var HND = Process_Dice_Value_INPUT(red_dice, blue_dice)
// 	// error handling.. If they didnt enture enough digits
// 	if HND.HID != "pending" {
// 		return
// 	}
// 	HND.GAME_SESSION = TMP_GAME_SESS // This is the ISESSION ID
// 	hand_num := len(HISTORY) + 1

// 	//5b Generate a unique ID for this hand (within this Game Session ID)
// 	tmpid := INT_to_STRING(hand_num) + "_Ra:" + INT_to_STRING(HND.RED_A) + "-Rb:" + INT_to_STRING(HND.RED_B)
// 	tmpid += "_Ba:" + INT_to_STRING(HND.BLUE_A) + "-Bb:" + INT_to_STRING(HND.BLUE_B)
// 	tmpid += "_SESS:" + TMP_GAME_SESS
// 	HND.HID = tmpid

// 	// = = =
// 	// 6. Process all the META DATA for this HAND
// 	MetaData_Generation_ENGINE(&HND, &HISTORY)

// 	// = = =
// 	// 7.  Evaluate for EVENTS. We have a seriues of custom EVENTS that we want to check for
// 	DL_Events_Engine(&HND, &HISTORY)

// 	// = = =
// 	//8. Now with everything checked generated and events checked for, we SAVE TO HISTORY
// 	HISTORY = append(HISTORY, HND)

// 	// = = =
// 	//9. ALSO!!! Update the "FUTURE_WININER" of the PREVIOUS hand in history
// 	hlen := len(HISTORY)
// 	if hlen >= 2 {
// 		pind := hlen - 2 // PREVIOUS item... before last one we just added
// 		prev := HISTORY[pind]

// 		// Update this hands NEXT_WINNER ... this is based on whoever won the most recent hand we JUST hadded
// 		prev.NEXT_WINNER = HND.WINNER
// 		HISTORY[pind] = prev
// 	}

// } // end of DICEY

func INIT_DiceDuel() {
	PLACEHOLDER()

	//1. Check to see if user provided CLI input
	//   or, If the user specified a file to load, then we load it and show the history
	// If so, we exit after its been processed
	var hard_exit = false
	var show_all = false
	if INPUT_RED_DICE != "" || INPUT_BLUE_DICE != "" {
		hard_exit = true
	}
	//1b If the user specified a file number, then we load it and show the history
	if LOAD_FILE_NUM > 0 {
		show_all = true
		hard_exit = true
	}

	//2. starts the main Loop
	for {
		//3. Main Dice Engine
		Dice_Engine(INPUT_RED_DICE, INPUT_BLUE_DICE)

		//4. Shows the dice history (so far) .. just last 10 hands
		ShowDice_HISTORY(show_all)

		if hard_exit {
			DO_EXIT(SILENT)
		}
	}

}
