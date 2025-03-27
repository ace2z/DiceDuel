package main

import (
	// = = = = = Native Libraries

	"fmt"
	"os"
	//"strings"
	"os/exec"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
	"golang.org/x/term"
)

func Read_User_Input(ALL_PARAMS ...interface{}) {
	var MESSAGE = " Please provide input: "
	var COLOR = WHITE

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

		if isCOLOR {
			COLOR = color_val
			continue
		}
	} //end of for

	//1. Switch stdin to raw mode to read characters without pressing Enter
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		M.Println("Error:", err)
		return
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
			return
		}
		var charNUM = b[0]
		var strVAL = string(b[0])

		//4. If CTRL_C, exit the program
		if charNUM == byte(CTRL_C) {

			RESET_TERMINAL()
			// rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
			// rawModeOff.Stdin = os.Stdin
			// _ = rawModeOff.Run()
			// rawModeOff.Wait()

			//os.Exit(0)
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
		C.Print(lastCHAR)

	} //end of for

	// W.Println("")
	// W.Print("You Entered: ")
	// G.Println(full_string)
}
