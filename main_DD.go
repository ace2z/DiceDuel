package main

import (
	// = = = = = Native Libraries

	. "local/CORE"
	. "local/DICEY"

	"time"

	//"reflect"
	//"bytes"
	//"encoding/binary"
	//"encoding/gob"
	//"reflect"
	//"github.com/qdm12/reprint"
	//. "local/INGEST_ENGINE"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
	//"os"
)

var VERSION_NUM = ""

//  *************************** MAIN AREA ****************************

func SUPER_FREAK_Test(ALL_PARAMS ...interface{}) {

	pm := Ingest_FUNC_PARAMS(ALL_PARAMS...)

	// Will search the MAP and see if we have a match for -sheep
	if pm.HAVE_Param("-sheep") {
		W.Println(" Sheep are SPECTACULAR!!!!")
	}

	if pm.HAVE_Param("--maptest") {
		SHOW_STRUCT(pm.MAP)
	}
	if pm.HAVE_Param("--int") {
		W.Println("Int: ", pm.INT)
	}
	if pm.HAVE_Param("-float") {
		W.Println("Float: ", pm.FLOAT)
	}

	if pm.HAVE_Param("-i64") {
		W.Println("Int64: ", pm.INT64)
	}

	if pm.HAVE_Param("-newcolor") {
		pm.COLOR.Println("You asked to print in THIS COLOR")
	}

	if pm.IsString(2) {
		W.Println("String: ", pm.STRING)
	}
	if pm.IsMAP() {
		M.Print("MAP is: ")
		SHOW_STRUCT(pm.MAP)

		mval_name := pm.MAP["name"].(string)
		Y.Print("Mval NAME is: ")
		C.Println(mval_name)
	}

	// If there is a GENERIC type structure we need to get, use this fucntion
	// If you want have multiples of the same GENERIC, pass an INT (like 2) to get the 2nd one
	if Param_GENERIC[TOLSTOY](pm, 2) {
		W.Println("")
		W.Println("")
		W.Println("YO, Here is your GENERIC: ")
		SHOW_STRUCT(pm.GENERIC)

		//To access the actual VALUE in a particular generic, use assertion like this:
		Y.Print("ACTUAL - Generic NAME is: ")
		C.Println(pm.GENERIC.(TOLSTOY).Name)
		W.Println("")
	}

} //end of

type TOLSTOY struct {
	Name           string
	Age            int
	Sex            string
	unexport_Field string

	hiddenTIME time.Time
}

// func COPY_Arr(arr []any) []string {
// 	var newArr []string
// 	for _, val := range arr {
// 		newArr = append(newArr, val)
// 	}
// 	return newArr
// }

func main() {

	myMap := NEW_Map()
	myMap["name"] = "John Doe"
	myMap["bodycount"] = 30
	myMap["city"] = "New York"

	var terryINT = 8675309
	var test64 int64 = 1234567890123457899
	var testSTR = "New Terry Test String"
	var tbool = true

	var asl = TOLSTOY{
		Name:           "Terry",
		Age:            55,
		Sex:            "Non-Binary",
		unexport_Field: "Morgan Freeman",
	}

	var SEC = TEST_DeepCopyEngine(asl).(TOLSTOY)

	SEC.Name = "CharlieCox"
	SEC.Age = 40
	SEC.Sex = "MALE"
	SEC.unexport_Field = "La Fay"
	Y.Println("--")
	new_SHOW_STRUCT(asl)
	new_SHOW_STRUCT(SEC, CYAN)
	Y.Println("--")
	DO_EXIT("-silent")

	var STRLIST = []string{"Daredevil", "Tony Stark", "Thor", "Black Widow"}
	//var sobj2 = STRLIST
	//sobj2[0] = "IronMan"

	// err, sFIN := COPY_DeepCopyEngine(STRLIST, &s2).([]string)
	//var s2 []string
	//sFIN := COPY_DeepCopyEngine(STRLIST)
	// if err != nil {
	// }
	// var s2 = sFIN.([]string)

	// s2[0] = "WorldWar HULK"

	// SHOW_STRUCT(STRLIST, GREEN)
	// SHOW_STRUCT(s2, CYAN)

	// //SHOW_STRUCT(sFIN, MAGENTA)

	// Y.Println("--")
	// PressAny()

	// //destMAP := COPY_Map(myMap)
	// //tmpMAP := COPY_DeepCopyEngine(myMap)
	// destMAP := tmpMAP.(MAP_TYPE)
	// destMAP["Updated destMAP"] = "SellTheKid_MarkoMark"
	// destMAP["bodycount"] = 145

	// SHOW_STRUCT(myMap)
	// SHOW_STRUCT(destMAP, CYAN)
	PressAny()

	SUPER_FREAK_Test("-sheep", "-maptest", myMap, "--int", 4, "-float", 5.56, "Your String TERRY", "-i64", test64, "-newcolor", BW)
	//SUPER_FREAK_Test("Your String TERRY", myMap1, "Second Simple STRING", asl, mapTWO, SEC)

	// Error handling
	if terryINT != 0 {
	}
	if test64 != 0 {
	}
	if testSTR != "" {
	}
	if tbool {
	}
	if asl != (TOLSTOY{}) {
	}
	if myMap != nil {
	}
	if len(STRLIST) > 0 {
	}

	PressAny()

	DO_EXIT()

	//1. Init the program (and the command line parameters)
	CLI_PARAM_INIT()
	MASTER_INIT("DiceDuel", VERSION_NUM)

	//2. Main Dice Engine
	if INPUT_RED_DICE != "" && INPUT_BLUE_DICE != "" {
		Dice_Engine_INIT(INPUT_RED_DICE, INPUT_BLUE_DICE)

		DO_EXIT("-silent")
	}

	// otherwise go into a loop
	for {

		//1. Main Dice Engine
		Dice_Engine_INIT("", "")

		//2. Show the Dice History  (last 10 hands)
		ShowDice_HISTORY(false)
	}

} // end of main
