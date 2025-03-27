package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/DLOGIC"

	//	"bufio"
	"fmt"
	"os"
	"unicode"
	//"strings"
	"os/exec"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/fatih/color"
	"golang.org/x/term"
	//"github.com/inancgumus/screen"
	//tea "github.com/charmbracelet/bubbletea"
)

type DLOGIC_OBJ struct {
	RED_INC    bool
	RED_DROP   bool
	HAVE_RED_6 bool
	HAVE_RED_1 bool
	RED_by_1   bool

	BLUE_INC    bool
	BLUE_DROP   bool
	HAVE_BLUE_6 bool
	HAVE_BLUE_1 bool
	BLUE_by_1   bool

	RED_B  int
	RED_A  int
	BLUE_B int
	BLUE_A int

	RB_DIFF int

	ALL_EVEN bool // if both red and blue dice are EVEN
	ALL_ODD  bool // if both red and blue dice are ODD

	WINNER    string
	WIN_COLOR *color.Color
}

var HISTORY []DLOGIC_OBJ

func Process_Dice_Value_INPUT(red_dice string, blue_dice string) {

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

	/*
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

	//3. Perform dice-logic
	var DL DLOGIC_OBJ
	DL.RED_A = red_a_int
	DL.RED_B = red_b_int
	DL.BLUE_A = blue_a_int
	DL.BLUE_B = blue_b_int
	main_Dice_Rules_Logic(red_b_int, red_a_int, "RED", &DL)
	main_Dice_Rules_Logic(blue_b_int, blue_a_int, "BLUE", &DL)

	//4. Get the Diff betwen Red and Blue for last two games
	DL.RB_DIFF = INT_GetDiff(red_b_int, blue_b_int)

	//5. Now determine the ODD vs EVEN
	if IS_EVEN(red_b_int) && IS_EVEN(blue_b_int) {
		DL.ALL_EVEN = true

	} else if IS_ODD(red_b_int) && IS_ODD(blue_b_int) {
		DL.ALL_ODD = true
	}

	HISTORY = append(HISTORY, DL)
}

// func Read_User_Input(ALL_PARAMS ...interface{}) string {

// 	var showtyped = false
// 	var useColor = WHITE

// 	for _, param := range ALL_PARAMS {
// 		string_val, is_string := param.(string)
// 		color_val, isCOLOR := param.(*color.Color)
// 		//int_val, is_int := param.(int)

// 		if is_string {
// 			if string_val == "-showtyped" {
// 				showtyped = true
// 			}
// 			continue
// 		}

// 		if isCOLOR {
// 			useColor = color_val
// 			continue
// 		}
// 	} //end

// 	if useColor == nil {
// 		useColor = WHITE
// 		tmp := "dummy"
// 		if strings.Contains(tmp, "RED") {
// 		}
// 	}

// 	reader := bufio.NewReader(os.Stdin)
// 	userTEMP, _ := reader.ReadString('\n')
// 	//userTEMP = strings.TrimSuffix(userTEMP, "\n")

// 	if showtyped {
// 		//useColor.Print(userTEMP)
// 	}

// 	return userTEMP

// } //end of

func allowed_Char(b []byte) bool {

	return true
}

var BACKSPACE = 127
var DELETE = 126
var ENTER = 13
var D_char = 100
var E_char = 101
var F_char = 102
var CTRL_C = 3

func Char_IS_Allowed(input []byte) bool {

	var firstCHAR = input[0]

	isDigit := unicode.IsDigit(rune(firstCHAR))
	if isDigit {
		return true
	}

	// if b == byte(BACKSPACE) || b == byte(ENTER) {
	// 	return true
	// }
	return false
}

// func showCursor() {
// 	if term.IsTerminal(int(os.Stdout.Fd())) {
// 		fmt.Print("\033[?25h")
// 	}
// }

func Read_USER_INPUT_RealTime() {
	// Switch stdin to raw mode to read characters without pressing Enter
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	Y.Print("Enter text (press Ctrl+C to exit): ")

	var full_string = ""
	for {
		// Read a single byte (character) from stdin
		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		var charNUM = b[0]
		var strVAL = string(b[0])

		//1. If CTRL_C, exit the program
		if charNUM == byte(CTRL_C) {

			rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
			rawModeOff.Stdin = os.Stdin
			_ = rawModeOff.Run()
			rawModeOff.Wait()

			os.Exit(0)
			//DO_EXIT("-silent")
		}

		//2. If they press BACKSPACE
		if charNUM == byte(BACKSPACE) || charNUM == byte(DELETE) {
			// Removes the last character from the string
			full_string = full_string[:len(full_string)-1]

			// This is the backspace character
			os.Stdout.Write([]byte{'\b', ' ', '\b'})
			continue
		}

		//3. If they press ENTER
		if charNUM == byte(ENTER) {
			break
		}

		full_string = full_string + strVAL
		lastCHAR := full_string[len(full_string)-1:]
		C.Print(lastCHAR)

		// Print the character read
		//fmt.Printf("You entered: %q\n", string(b))
	} //end of for

	W.Println("")
	W.Print("You Entered: ")
	G.Println(full_string)
}

func Dice_Engine_INIT(INPUT_RED_DICE string, INPUT_BLUE_DICE string) {

	// PROMPT("Welcome to the Dice Duel Game!")
	// PROMPT("The game is simple: You and the computer each roll a die. The highest number wins.")

	var red_dice = ""
	var blue_dice = ""
	if INPUT_RED_DICE != "" {
		red_dice = INPUT_RED_DICE
	}
	if INPUT_BLUE_DICE != "" {
		blue_dice = INPUT_BLUE_DICE
	}
	Y.Println("")
	Y.Print(" Enter ")
	W.Print("last 2 Dice Games ")
	Y.Println("RED first ")
	M.Print("   RED DICE: ")

	if red_dice == "" {
		red_dice = "" //Read_User_Input()

		// if user typed e or d or something, we remove last item and start over again
		if Need_to_FIX_PREVIOUS(red_dice) {
			return
		}
	}

	C.Print("  BLUE DICE: ")
	if blue_dice == "" {

		blue_dice = "" //Read_User_Input()
		// if user typed e or d or something, we remove last item and start over again
		if Need_to_FIX_PREVIOUS(blue_dice) {
			return
		}
	}
	W.Println("")

	//2. Perform the dice logic
	Process_Dice_Value_INPUT(red_dice, blue_dice)

} // end of main
