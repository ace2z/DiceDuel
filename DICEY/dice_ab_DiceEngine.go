package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"
	. "local/EVENTS"
	. "local/META"

	. "github.com/ace2z/GOGO/Gadgets"
)

/*
Placeholder for Game Session ID ( for future use)
Every game (series of hands / rolls) is a 'session' and all the hands within it
Share the same Game Session ID
If you need to save this to a database, you can use this to ensure all games and hands are UNIQUE
*/
var TMP_GAME_SESS = "8675309"

var also_allowed = []string{"e", "r", "d", "v", "h", "s", "l"}

func Dice_Engine(INPUT_RED_DICE string, INPUT_BLUE_DICE string) {

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
		Y.Print("...do ")
		M.Print("RED ")
		Y.Println("first ")

		tmp_red := Read_User_Input("      RED DICE: ", BOLD_MAGENTA, BOLD_YELLOW, 2, also_allowed, "-digits", NOEOL)
		if tmp_red == "" || Extra_KEYS_Handle_Engine(tmp_red) {
			W.Println("")
			return
		}
		tmp_blue := Read_User_Input("     BLUE DICE: ", BOLD_CYAN, BOLD_WHITE, 2, also_allowed, "-digits")
		if tmp_blue == "" || Extra_KEYS_Handle_Engine(tmp_blue) {
			W.Println("")
			return
		}

		red_dice = tmp_red
		blue_dice = tmp_blue

		W.Println("")
	}

	//5. Process the Dice Input. Get the Hand OBJECT
	var HND = Process_Dice_Value_INPUT(red_dice, blue_dice)

	// error handling.. If they didnt enture enough digits
	if HND.HID != "pending" {
		return
	}
	HND.GAME_SESSION = TMP_GAME_SESS // This is the ISESSION ID
	hand_num := len(HISTORY) + 1

	//5b Generate a unique ID for this hand (within this Game Session ID)
	tmpid := INT_to_STRING(hand_num) + "_Ra:" + INT_to_STRING(HND.RED_A) + "-Rb:" + INT_to_STRING(HND.RED_B)
	tmpid += "_Ba:" + INT_to_STRING(HND.BLUE_A) + "-Bb:" + INT_to_STRING(HND.BLUE_B)
	tmpid += "_SESS:" + TMP_GAME_SESS
	HND.HID = tmpid

	// = = =
	// 6. Process all the META DATA for this HAND
	MetaData_Generation_ENGINE(&HND, &HISTORY)

	// = = =
	// 7.  Evaluate for EVENTS. We have a seriues of custom EVENTS that we want to check for
	DL_Events_Engine(&HND, &HISTORY)

	// = = =
	//8. Now with everything checked generated and events checked for, we SAVE TO HISTORY
	HISTORY = append(HISTORY, HND)

	// = = =
	//9. ALSO!!! Update the "FUTURE_WININER" of the PREVIOUS hand in history
	hlen := len(HISTORY)
	if hlen >= 2 {
		pind := hlen - 2 // PREVIOUS item... before last one we just added
		prev := HISTORY[pind]

		// Update this hands NEXT_WINNER ... this is based on whoever won the most recent hand we JUST hadded
		prev.NEXT_WINNER = HND.WINNER
		HISTORY[pind] = prev
	}

} // end of DiceEngine
