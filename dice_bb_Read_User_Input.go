package main

import (
	// = = = = = Native Libraries

	//"fmt"
	"os"
	//"strings"
	//"os/exec"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
	"golang.org/x/term"
)

var BACKSPACE = 127
var DELETE = 126
var ENTER = 13
var D_char = 100
var E_char = 101
var F_char = 102
var CTRL_C = 3

// color.New(color.FgMagenta, color.Bold)

func Read_User_Input(ALL_PARAMS ...interface{}) string {
	RESET_TERMINAL()
	var MESSAGE = " Please provide input: "

	var COLOR = WHITE
	var USER_TEXT_COLOR = BOLD_YELLOW

	var first_set = false

	for _, param := range ALL_PARAMS {
		string_val, isString := param.(string)
		color_val, isCOLOR := param.(*color.Color)
		//int_val, is_int := param.(int)

		if isString {
			// if string_val == "-showtyped" {
			// 	showtyped = true
			// }

			// If anything else.. assume its the message to display
			MESSAGE = string_val
			continue
		}

		// If the param is a color
		// FIRST color passed will be the color of the message
		// SECOND time color is passed will be the color of the USERS text when they typed
		if isCOLOR {
			if first_set == false {
				COLOR = color_val
				first_set = true
			} else {
				USER_TEXT_COLOR = color_val
			}
			continue
		}
	} //end of for

	//1. Switch stdin to raw mode to read characters without pressing Enter
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		M.Println("Error:", err)
		return ""
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	//2. Print the MESAGE
	COLOR.Print(MESSAGE)

	//3. Now LOOP to read characters
	// we save everythign into final string
	var final_STRING = ""

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

		//4. If CTRL_C, exit the program
		if charNUM == byte(CTRL_C) {

			RESET_TERMINAL()

			DO_EXIT("-silent")
		}

		//5. If they press BACKSPACE
		if charNUM == byte(BACKSPACE) || charNUM == byte(DELETE) {
			// Removes the last character from the string
			final_STRING = final_STRING[:len(final_STRING)-1]

			// This is the backspace character
			os.Stdout.Write([]byte{'\b', ' ', '\b'})
			continue
		}

		//6. If they press ENTER
		if charNUM == byte(ENTER) {
			break
		}

		final_STRING = final_STRING + strVAL
		lastCHAR := final_STRING[len(final_STRING)-1:]
		USER_TEXT_COLOR.Print(lastCHAR)

	} //end of for

	RESET_CONSOLE()
	WHITE.Println("")

	return final_STRING
}
