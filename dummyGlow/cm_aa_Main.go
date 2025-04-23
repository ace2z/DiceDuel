package main

import (
	// = = = = = Native Libraries
	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	"github.com/fatih/color"
	//"sort"
	//"strconv"
	//"strings"
	"fmt"
	//"strings"
)

// *************************** MAIN AREA ****************************

type CLR_OBJ struct {
	DESC  string   // Friendly name for this color
	FG    int      // The numeric 256 value of the color
	BG    int      // The background number for this color (if applicable)
	ATTRS []string // These are the Attribute codes for Terminal Functions (blink, bold, underline etc)
	ID    string   // A unique ID for this color,
}

var COLOR_MATRIX = []CLR_OBJ{
	{DESC: C_BLACK, FG: 0},
	{DESC: C_BROWN, FG: 1},
	{DESC: C_drk_GREEN, FG: 2},
	{DESC: C_OLIVE, FG: 3},
	{DESC: C_drk_BLUE, FG: 4},
	{DESC: C_PURPLE, FG: 5},
	{DESC: C_drk_CYAN, FG: 6},
	{DESC: C_WHITE, FG: 7},
	{DESC: C_GREY, FG: 8},
	{DESC: C_RED, FG: 9, BG: 0, ATTRS: []string{C_BOLD}},
	{DESC: C_GREEN, FG: 10},
	{DESC: C_YELLOW, FG: 11, BG: 0, ATTRS: []string{C_BOLD}},
	{DESC: C_BLUE, FG: 12},
	{DESC: C_MAGENTA, FG: 13, BG: 0, ATTRS: []string{C_BOLD}},
	{DESC: C_CYAN, FG: 14},
	{DESC: C_Hi_WHITE, FG: 15},
}

func COLOR_Matrix_Engine(show_TEXT string, ALL_PARAMS ...any) {
	Y.Println(" Started COLOR_Matrix_Engine")

	color_formatted, notes := ESC_Code_Generator(90, 3, C_BOLD, C_FG_BRIGHT) //, BG_BRIGHT)
	Y.Print(" Notes: ", notes)
	SHOW_STRUCT(color_formatted, GREEN)

	var FG_BOLD_RED = color.New(color.FgRed).Add(color.Bold)
	var BG_BOLD_RED = color.New(color.BgRed)
	fmt.Println(color_formatted + " ■■■■■■ - Custom Freedom      " + ec_RESET)
	fmt.Println("")

	FG_BOLD_RED.Println(" ■■■■■■ -=-=- FG FATIH the RED Goat")
	BG_BOLD_RED.Println(" ■■■■■■ -=-=- BackGround FATIH BACKGROUND       ")
	BOLD_YELLOW.Println(" ■■■■■■ -=-=- BOLD YELLOW -=- GOATED")
	BOLD_GREEN.Println(" ■■■■■■ -=-=- BOLD GREEN -=- GOATED")
	fmt.Println("")

	PressAny()

	// Populate the COLOR_MATRIX with the 256 colors DESC
	for i := 0; i < 256; i++ {

		var is_STD = i < 16
		if is_STD {
			continue
		}

		// Otherwise, we geneerate a new item for one of the higher colors
		var obj = CLR_OBJ{}
		//convert ind to a string
		var ind_text = fmt.Sprintf("%d", i)
		obj.DESC = "C_" + ind_text
		obj.FG = i
		obj.BG = 0
		COLOR_MATRIX = append(COLOR_MATRIX, obj)
	}

	// tmp := param_PREF + "miso"
	// //tmp := "miso"
	// W.Println(tmp)
	// C.Print(" Is FuncParam?: ")
	// Y.Println(is_FuncParam(tmp, "miso"))

	// purp := NewColor(C_BROWN)
	// mage := NewColor(C_RED)

	// fgnum := 9
	// bgnum := 0
	// e_ESC := "\x1b[3"
	// e_fg := e_ESC + "38;5;"
	// e_bg := e_ESC + "48;5;"
	// e_BOLD := e_ESC + "1m"
	// e_CLOSE := "m"
	// //e_reset := e_ESC + "0m"
	// e_message := "■■■■■■ | heaven"

	// reset: \x1b[0m
	//redBoldGreenBackground := "\x1b[1;31;42m"
	//redBoldGreenBackground := "\x1b[101m"

	CC := 1

	for_bg := true
	for_bright := true
	// If less than 15, we need a different hex modifier
	if CC < 15 {
		if for_bg {
			CC += 40
		} else {
			CC += 30
		}

		// If we are turning on BRIGHT

		if for_bright {
			CC += 60
		}

		// Otehrwise we are needing to use 256 color mode
	}

	CCtext := fmt.Sprintf("%d", CC)
	formatted_Color_Line := hex_esc + CCtext + "m"
	//redBoldGreenBackground := "\x1b[1;33;118m"
	fmt.Println(formatted_Color_Line + " ■■■■■■ This is red text with a green background " + "\x1b[0m")

	PressAny()

	//CCtext += ";5;"

	e_BOLD := ";1"
	e_BOLD = ""
	e_CLOSE := "m"
	//fgnum := "130"

	//fmt.Println("\x1b[32mThis text is green.\x1b[0m")
	fmt.Println(hex_esc + CCtext + e_BOLD + e_CLOSE + "■■■■■■ | he2aven " + ec_RESET)
	//fmt.Printf("%s%d%s %s%d%s %s", e_fg, fgnum, e_CLOSE, e_bg, bgnum, e_CLOSE, e_BOLD)
	//fmt.Println(e_message)
	fmt.Println("")

	fmt.Println("\033[33;1m ■■■■■■ This is Bright Yellow\033[0m")
	fmt.Println("\033[33m ■■■■■■ This is Yellow\033[0m")

	PressAny()

	green := NewColor(C_GREEN) //, C_RED, C_drk_CYAN) //, "reverse", "bold")
	//SHOW_STRUCT(TJ, MAGENTA)
	//Print_ENGINE("Everything Everywhere", TJ)
	// fmt.Print("There is a lady thats sure", 5559191, "all that glitters is gold")
	// fmt.Println(" **end")
	// fmt.Println("You know you Want it", 8675309, "carry on my wayward song")

	green.Print("Bying a stairway to heaven", 919181, "Men who star at goats")
	fmt.Println(" **end")
	green.Println("You know you Want it", 8675309, "carry on my wayward song")
	fmt.Println(" **end")

	W.Println("")
	W.Println("")

	Print_ENGINE("Terry Carry On? ", C_MAGENTA)
	Print_ENGINE("Purple? ", C_PURPLE)
	fmt.Println(" **end")
	PressAny()

	// //2. Now we have all 256 colors int he matrix. Lets setup the REVERSE colors (compounds)
	// // // Npw print all colors in the matrix

	// var bg_colors []CLR_OBJ
	// for _, clr := range COLOR_MATRIX {
	// 	if strings.Contains(clr.DESC, "C_") {
	// 		clr.DESC = strings.Replace(clr.DESC, "C_", "BG_", -1)
	// 	} else {
	// 		clr.DESC = "BG_" + clr.DESC
	// 	}
	// 	clr.BG = clr.FG
	// 	clr.FG = 0 //0 for black. Override if another color looks better for the foreground
	// 	bg_colors = append(bg_colors, clr)
	// }
	// COLOR_MATRIX = append(COLOR_MATRIX, bg_colors...)

	// cnt := 0
	// for _, clr := range COLOR_MATRIX {
	// 	nlflag := false
	// 	Ansi_Print(" "+clr.DESC+" ", clr.FG, clr.BG, nlflag)
	// 	// fmt.Print(" ")
	// 	// Ansi_Print(" "+clr.DESC+" ", 0, clr.FG, nlflag)

	// 	cnt++
	// 	if cnt%12 == 0 {
	// 		fmt.Println("")
	// 	}
	// }

	// PressAny()
	// var msg = " Terry#"

	// // DEbug
	// var fgSEQ = "\x1b[38;5;"
	// //var bgSEQ = "\033[48;5;"
	// var RESET = "\x1b[0m"

	// // fmt.Printf(fgSEQ + "\x1b[38;5;%dm%3d\x1b[0m ", i, i)
	// for i := 0; i < 256; i++ {
	// 	fmt.Printf(fgSEQ+"%dm"+msg+"%d"+RESET, i, i)
	// 	if (i+1)%16 == 0 {
	// 		fmt.Println()
	// 	}
	// }

	// PressAny()
	// if i >= 8 {
	// 	fmt.Printf("\033[1;%dmColor %d\033[0m\n", colorCode, i) // Bright colors
	// } else {
	// 	fmt.Printf("\033[0;%dmColor %d\033[0m\n", colorCode, i) // Standard colors
	// }
	//}
	// var start_seq = "\033["
	// var STD_SEQ = "\033[0;"
	// var BOLD_SEQ = "\033[1;"
	// var RESET = "\033[0m"
	// var end_seq = "\033[0m"

	// No point in doing the standard colors if we are using 256 colors terminal
	// for i := 0; i < 16; i++ {
	// 	var cc_ALGO = 30 + i%8

	// 	// The first 8 colors are the standard colors
	// 	// The next 8 colors are bright/bold versions of the first 8
	// 	if i <= 6 {
	// 		fmt.Printf(STD_SEQ+`%dm`+msg+`%d `+RESET, cc_ALGO, i)
	// 	}
	// 	if i== 6 {
	// 		fmt.Println("")
	// 	}
	// 	// Then the bright colors
	// 	fmt.Printf(BOLD_SEQ+`%dm`+msg+`%d `+RESET, cc_ALGO, i)

	// 	if i == 8 {
	// 		fmt.Println("")
	// 	} else if i == 15 {
	// 		fmt.Println("")
	// 	}
	// }

	// ==== OLD STUFF ====
	// var style = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("3")).Foreground(lipgloss.Color("1"))

	// fmt.Print(style.Render(" TEST HOT DOG "))
	// PressAny()
	// //1. Lets populate with our Preload colors
	// COLOR_MATRIX = append(COLOR_MATRIX, preload_Colors...)

	// //2. Generate COLOR ID's when we get to COmpounding colors, we will need them
	// GEN_Color_IDs()

	// //3. Now lets generate compound / overlay colors
	// Generate_Compound_Colors()

	// Print(" Dummy ", "red", "hi_green")
	// W.Println(" Its Just Me")

	// // Print(" Dummy ", "red", "hi_green")
	// // W.Println(" Its Just Me")

	// PressAny()

	// for _, item := range COLOR_MATRIX {
	// 	W.Print(item.DESC, " ===>  ")
	// 	C.Print(item.ID, " ")
	// 	Y.Print(item.comp_test, " ")
	// 	Println(" TEST ", item.ID)
	// }

}
