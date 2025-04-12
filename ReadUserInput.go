package main

import (
	// = = = = = Native Libraries

	//"fmt"
	"os"
	//"strings"
	//"os/exec"

	//"github.com/fatih/color"
	. "github.com/ace2z/GOGO/Gadgets"
	"golang.org/x/term"
)

var kc_BACKSPACE = 127
var kc_DELETE = 126
var kc_ENTER = 13

// var D_char = 100
// var E_char = 101
// var F_char = 102
var kc_CTRL_C = 3

var tmp_NUMS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var tmp_LOWER = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var tmp_UPPER = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// Ultimate READ USER INPUT... Pass any of a the following:
// * a STRING which is the message/question displayed to the user
// * a *color.Color (color fatih) which is the color we display the messsage with
// * a 2nd color, which is the color of the USER input as they type
// // Use either of the following to limit the characters the user can type
// * ...either an ARRAY of stringS (containing single letters/chars) to RESTRICT user input to these chars
// * ...or "-allow|a|b|c" etc... to RESTRICT user input to THOSE characters
//
// The following can be COMBINED together (NOT mutually exclusive)...and OVERRIDE -allow or an []string input
// * "-digits" which restricts input to ONLY digits (0 - 9)
// * "-alpha" which restricts input to ONLY letters (a - z) and (A - Z)
//
// the following augment behavior
// * "-min|2" which REQUIRES user to type at least 2 chars
// * "-max|10" which restricts MAX chars user can type to 10
// * "-hide|-hidden|-pass" to make the input a PASSWORD (will show * )
// * "-noeol" to NOT show a newline after the user presses kc_ENTER
//
// The following allow for looped "retry" user input
// * "-valid_entry" to require the user to type a SPECIFIC string to return success
// * "-retry|3", if -valid_entry was passed, user has 3 attempts to type the correct value
// * "-exit_on_fail" If valid_entry is provided, Forces program to EXIT if the users fails to type the proper parameter
func beta_Read_User_Input(ALL_PARAMS ...interface{}) string {

	var MESSAGE = " Please provide input: " // Default message we display

	var COLOR = WHITE
	var USER_TEXT_COLOR = BOLD_YELLOW

	var ALLOWED_CHARS []string

	var MAX_CHARS = 0
	var min_chars = 0 // This is the minimum number of characters we REQUIRE to be typed

	var PASSWORD_MODE = false
	var SHOW_NEWLINE = true

	var need_valid_entry = ""
	var retry_max = 1
	var exit_on_fail = false

	// Ingest the parameters
	var pm = Ingest_FUNC_PARAMS(ALL_PARAMS...)
	if pm.IsString() {
		MESSAGE = pm.STRING
	}
	if pm.IsColor() {
		COLOR = pm.COLOR
	}
	if pm.IsColor(2) {
		USER_TEXT_COLOR = pm.COLOR
	}

	// If passed, overrides everything else
	if pm.IsARR() {
		ALLOWED_CHARS = []string{}
		ALLOWED_CHARS = pm.ARR
	}
	if pm.HAVE_Param("-allow") {
		ALLOWED_CHARS = []string{}
		ALLOWED_CHARS = append(ALLOWED_CHARS, pm.PIPE[:1]...)
	}

	// the Following can be COMBINED
	if pm.HAVE_Param("-digits") {
		//ALLOWED_CHARS = []string{}
		ALLOWED_CHARS = append(ALLOWED_CHARS, tmp_NUMS...)
	}
	if pm.HAVE_Param("-alpha") {
		//ALLOWED_CHARS = []string{}
		ALLOWED_CHARS = append(ALLOWED_CHARS, tmp_UPPER...)
		ALLOWED_CHARS = append(ALLOWED_CHARS, tmp_LOWER...)
	}
	if pm.HAVE_Param("-min") {
		min_chars = STRING_to_INT(pm.PIPE[1])
	}
	if pm.HAVE_Param("-max") {
		MAX_CHARS = STRING_to_INT(pm.PIPE[1])
	}
	if pm.HAVE_Param("-pass", "-hid") {
		PASSWORD_MODE = true
	}
	if pm.HAVE_Param("-noeol") {
		SHOW_NEWLINE = false
	}

	// = = = = = VALID User Entry verification
	// If passed, overrides everything else
	if pm.HAVE_Param("-valid_entry") {
		need_valid_entry = pm.PIPE[1]
	}
	if pm.HAVE_Param("-retry") {
		retry_max = STRING_to_INT(pm.PIPE[1])
	}
	if pm.HAVE_Param("-exit_on_fail") {
		exit_on_fail = true
	}

	//var first_set = false
	// for _, param := range ALL_PARAMS {
	// 	string_val, isString := param.(string)
	// 	color_val, isCOLOR := param.(*color.Color)
	// 	int_val, isINT := param.(int)
	// 	arr_val, isARR := param.([]string)

	// 	if isString {

	// 		// If they pass -allow|a|b|c  ... we only allow the user to type those characters
	// 		if HAS_PREFIX(string_val, "-allow") {
	// 			arr, _ := PIPE_SPLIT(string_val)
	// 			for n, x := range arr {
	// 				// skip the first item which is the named parameter itself
	// 				if n == 0 {
	// 					continue
	// 				}
	// 				ALLOWED_CHARS = append(ALLOWED_CHARS, x)
	// 			}
	// 			continue
	// 		}
	// 		// only allow digits.
	// 		if HAS_PREFIX(string_val, "-digits") {
	// 			ALLOWED_CHARS = append(ALLOWED_CHARS, tmp_NUMS...)
	// 			continue
	// 		}
	// 		// Makes the users input hidden (shows * instead of the actual character)
	// 		if HAS_PREFIX(string_val, "-pass") || HAS_PREFIX(string_val, "-hid") {
	// 			PASSWORD_MODE = true
	// 			continue
	// 		}

	// 		// if they pass min delimited by pipe... we know the user input should have AT LEAST this many characters
	// 		if HAS_PREFIX(string_val, "-min") {
	// 			_, mintext := PIPE_SPLIT(string_val, 1)

	// 			min_chars = STRING_to_INT(mintext)
	// 			continue
	// 		}

	// 		if HAS_PREFIX(string_val, NOEOL) {
	// 			SHOW_NEWLINE = false
	// 			continue
	// 		}

	// 		// If anything else.. assume its the message to display
	// 		MESSAGE = string_val
	// 		continue
	// 	}

	// 	// If the param is a color
	// 	if isCOLOR {
	// 		if first_set == false {
	// 			COLOR = color_val
	// 			first_set = true
	// 		} else {
	// 			USER_TEXT_COLOR = color_val
	// 		}
	// 		continue
	// 	}

	// 	// If number was passed, assume it is the LIMIT to the characters we allow user to type
	// 	if isINT {
	// 		MAX_CHARS = int_val
	// 		continue
	// 	}

	// 	// Alternate means of limiting what characters the user can type
	// 	if isARR {
	// 		ALLOWED_CHARS = arr_val
	// 		continue
	// 	}
	// } //end of for

	//1b. Error handling for -need_valid_entry (if specified)
	if need_valid_entry != "" {
		ALLOWED_CHARS = []string{}
		//MAX_CHARS = len(valid_entry)
		min_chars = 0 // reset if not already
		SHOW_NEWLINE = true
	}

	// = = =
	// 3. This OUTER loop facilities "valid entry" and "retry" logic
	var final_STRING = ""
	//var already_made_raw = false
	for i := 0; i < retry_max; i++ {
		final_STRING = ""

		//3b. Switch stdin to raw mode to read characters without pressing Enter

		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			M.Println("Error:", err)
			return ""
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		//3c. Print the MESAGE
		COLOR.Print(MESSAGE)

		//3d. Now LOOP to read characters from user
		for {
			// Read a single byte (character) from stdin
			b := make([]byte, 1)
			_, err := os.Stdin.Read(b)
			if err != nil {
				M.Println(err)
				return ""
			}
			var charNUM = b[0]
			var strVAL = string(b[0])

			//4. If kc_CTRL_C, exit the program
			if charNUM == byte(kc_CTRL_C) {

				RESET_TERMINAL()
				DO_EXIT("-silent")
			}

			//5. If they press kc_BACKSPACE
			if charNUM == byte(kc_BACKSPACE) || charNUM == byte(kc_DELETE) {
				// Removes the last character from the string
				if len(final_STRING) > 0 {
					final_STRING = final_STRING[:len(final_STRING)-1]

					// This is the backspace character
					os.Stdout.Write([]byte{'\b', ' ', '\b'})
				}
				continue
			}

			//6. If they press kc_ENTER,  Leave User Entry Inner Loop
			if charNUM == byte(kc_ENTER) {

				// Make sure we have at least min_chars (if it is not 0)
				if min_chars > 0 {
					if len(final_STRING) < min_chars {
						continue
					}
				}

				// Leave User Entry Inner Loop
				break
			}

			//7. If we are LIMITING characters, we DONT add anymore to final_STRING
			if MAX_CHARS > 0 {
				if len(final_STRING) == MAX_CHARS {
					continue
				}
			}

			// = = = = =
			//8. See if we have any of the allowed chars
			if len(ALLOWED_CHARS) > 0 {
				found := false
				for _, x := range ALLOWED_CHARS {
					if x == strVAL {
						found = true
						break
					}
					continue
				}

				if found == false {
					continue
				}
			}

			//10. This actually prints the character and saves it to the final string
			final_STRING = final_STRING + strVAL

			char_2show := "*" // Defaults to a star as if it were a password
			if PASSWORD_MODE == false {
				lastCHAR := final_STRING[len(final_STRING)-1:]
				char_2show = lastCHAR
			}

			//3. This shows the character the user typed
			USER_TEXT_COLOR.Print(char_2show)

		} //end of inner user Entry For

		// = = = =
		//8. Now if they are using the -need_valid_entry, we need to check if they typed the right thing
		if need_valid_entry != "" {
			RESET_TERM()

			W.Println("")
			if final_STRING == need_valid_entry {
				WG.Println(" Entry Accepted! ")
				W.Println("")
				RESET_CONSOLE()
				return final_STRING
			} else {
				WM.Println(" !!!! Invalid Response !!!! ")
				W.Println("")
				r := i + 1
				r = retry_max - r
				W.Print(r)
				// Y.Print(" of ")
				// Y.Print(retry_max)
				Y.Println(" attempts remaining...")
				W.Println("")

				continue
			}

		}
	} //end of outer retry loop

	// If they got this far, (and we are using -valid_entry). User has FAILED AT LIFE!
	// See what we need it do
	if need_valid_entry != "" {
		YM.Println(" Maximum attempts reached... ")
		if exit_on_fail {
			YM.Println(" Goodbye!")
			RESET_CONSOLE()
			DO_EXIT("-silent")
		}
	}

	// Otherwise re return the final string to the user (they can process further if needed)
	if SHOW_NEWLINE == true {
		WHITE.Println("")
	}
	RESET_CONSOLE()

	return final_STRING
}
