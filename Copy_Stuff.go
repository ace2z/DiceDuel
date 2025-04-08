package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"
	"bytes"

	"encoding/gob"
	"reflect"
	"unsafe"

	. "github.com/ace2z/GOGO/Gadgets"

	//"encoding/json"
	"github.com/fatih/color"
)

func helper_Show(mtype string) {
	PREFIX("DeepCopyEngine, copying: ", CYAN, INS_NEWLINE, NOEOL)

	Y.Println(mtype)

}

/*
Master OBJ, Struct,Map Slice copy function
GO Does 'shallow copying' of objects by default.
Since most of the time we cant to copy object to a COMPLETELY NEW INSTANCE
We need to DEEP COPY objects.

This performs a DEEP copy on Slices (ie  []string  ), Maps, and Structs
This ensures the SOURCE object and DEST object are always DIFFRERNT
YOU MUST USE assertion to get the ACTION value of the returned object!!!
*/
//func COPY_DeepCopy_Engine(SOURCE any, DEST any) (error, any) {
// func COPY_DeepCopyEngine(SOURCE any) any {
// 	FNAME := "DeepCopyEngine"
// 	PLACEHOLDER()
// 	//var FNAME = "COPY_DeepCopy_Engine: "
// 	origTYPE := reflect.TypeOf(SOURCE)
// 	origKIND := origTYPE.Kind()
// 	C.Println("ORIG_TYPE: ", origTYPE)
// 	Y.Println(" ORIGKIND: ", origKIND)
// 	W.Println("")

// 	gob.Register(origTYPE)
// 	var is_MAP = origKIND == reflect.Map
// 	var is_SLICE = origKIND == reflect.Slice
// 	var is_STRUCT = origKIND == reflect.Struct
// 	var is_ARRAY = origKIND == reflect.Array
// 	//var is_INTERFACE = origKIND == reflect.Interface

// 	if is_MAP == false && is_SLICE == false && is_STRUCT == false && is_ARRAY == false {
// 		M.Println(FNAME, "  ERROR: Invalid TYPE")
// 		Y.Println("  SOURCE is NOT a valid MAP, SLICE, or STRUCT")
// 		DO_EXIT()
// 		//return nil, nil
// 	}

// 	//1. Create the enc and DEC onbjects we need. Attach them to a buffer
// 	var buf bytes.Buffer
// 	encOBJ := gob.NewEncoder(&buf)
// 	decOBJ := gob.NewDecoder(&buf)

// 	//2. Encode the source value, and make sure we dont have any errors
// 	if err := encOBJ.Encode(SOURCE); err != nil {
// 		M.Print(FNAME, "ENCODE Error: ")
// 		Y.Println(err)
// 		DO_EXIT()
// 		//return err, nil
// 	}

// 	//3. Create the target object we are copying TO
// 	//var TARGET = reflect.New(origTYPE)
// 	var TARGET = reflect.New(reflect.TypeOf(SOURCE))

// 	//4. Now DECODE the buffer into the target object
// 	if err := decOBJ.Decode(TARGET.Interface()); err != nil {
// 		M.Print(FNAME)
// 		W.Print("DECODE, ")
// 		M.Print(" ERROR: ")
// 		Y.Println(err)
// 		DO_EXIT()
// 		//return err, nil
// 	}

// 	//4b. For STRUCTS, we need to do extra handling to deal with Unexported fields
// 	// This is because the gob package only encodes exported fields by default
// 	if is_STRUCT {
// 		origVAL := reflect.ValueOf(SOURCE).Elem()
// 		TARGval := reflect.ValueOf(TARGET).Elem()

// 		for i := 0; i < origVAL.NumField(); i++ {
// 			field := origVAL.Type().Field(i)
// 			if !field.IsExported() {
// 				TARGval.Field(i).Set(origVAL.Field(i))
// 			}
// 		}
// 	}

//		//5. Return the target object
//		return TARGET.Elem().Interface()
//	}

func helper_StructField_META(field reflect.StructField) {
	/*
		var stype = reflect.TypeOf(tmpSTRUCT)
		var total_fields = stype.NumField()

		field := stype.Field(i)
		ftype := field.Type
		fname := field.Name

		var not_Exportable = field.IsExported() == false
		//field := srcElem.Field(i)

		FVAL := reflect.ValueOf(field)

		fvtemp := reflect.ValueOf(&ftype).Elem()
		FADDY := fvtemp.Addr().Interface()
	*/

	//	var ftmp = reflect.TypeOf(fOBJ(ind))
	//var field = stype.Field(ind)
	//C.Println("field: ", field)
	//var ftype = reflect.TypeOf(fTMP).Key().Elem()
	var fname = field.Name
	var ftype = field.Type
	var fval = reflect.ValueOf(field)

	// var fvtemp = reflect.ValueOf(&ftype).Elem()
	// var ftmp2 = fvtemp.Type()
	//var faddy = ftmp2.Addr().Interface()

	//var ftmp = reflect.ValueOf(fval).Elem()
	var faddy = fval.Addr().Interface()

	//var not_Exportable = ofield.IsExported() == false
	var not_Exportable = false

	//var field = fTMP.Field

	//var fname = oftype.Name

	//var fTMP = reflect.TypeOf(fOBJ)

	// var ftype = field.Type
	// var fname = field.Name

	// var fval = reflect.ValueOf(field)

	//var not_Exportable = false

	// var ftype = reflect.TypeOf(fOBJ)
	// var fname = ftype.Name()
	// var fval = reflect.ValueOf(ftype)

	W.Print("Fname: ")
	Y.Println(fname)

	W.Print("Ftype: ")
	Y.Println(ftype)
	W.Print("FVal: ")
	Y.Println(fval)

	W.Print("FADDY: ")
	Y.Println(faddy)

	W.Print("Not Exported: ")
	Y.Println(not_Exportable)

	PressAny()
}

func new_SHOW_STRUCT(tmpSTRUCT any, ALL_PARAMS ...any) {

	var useCOLOR = YELLOW

	// Collects the input params specified... supports INT and FLOAT dynamically
	for _, param := range ALL_PARAMS {
		//str_val, isSTR := param.(string)
		color_val, isCOLOR := param.(*color.Color)

		// if a COLOR is specified, this is how the struct will be displayed (defaults to yellow)
		if isCOLOR {
			useCOLOR = color_val
			continue
		}
	}
	if useCOLOR == nil {

	}

	var stype = reflect.TypeOf(tmpSTRUCT)
	M.Println("STYPE: ", stype)
	var total_fields = stype.NumField()
	//var srcValue = reflect.ValueOf(&tmpSTRUCT)
	//var srcElem = srcValue.Elem()

	//valueOfS := reflect.ValueOf(tmpSTRUCT)
	//typeOfS := valueOfS.Type()

	for i := 0; i < total_fields; i++ {

		field := stype.Field(i)
		ftype := field.Type
		fname := field.Name

		var not_Exportable = field.IsExported() == false
		//field := srcElem.Field(i)

		FVAL := reflect.ValueOf(field)

		fvtemp := reflect.ValueOf(&ftype).Elem()
		FADDY := fvtemp.Addr().Interface()
		W.Print("")
		G.Print(i)
		C.Print(" - Not Exported: ")
		Y.Println(not_Exportable)
		C.Println("     fname: ", fname)
		C.Println("     ftype: ", ftype)
		C.Println("      FVAL: ", FVAL)
		C.Println("     FADDY: ", FADDY)
		W.Println("")
		helper_StructField_META(field)

		continue

		//field := stype.Field(i)

		// Get the field's value
		//fieldValue := FVAL.Field(i)

		// fieldName := typeOfS.Field(i).Name
		// // Accessing hidden fields
		// fieldValue := unsafe.Pointer(FVAL.Addr().UnsafeAddr())
		// useCOLOR.Printf("Field Name: %s, Value: %v\n", fieldName, *(*interface{})(fieldValue))
		// //fmt.Printf("Field Name: %s, Value: %v\n", fieldName, *(*interface{})(fieldValue))
	}

	// byte_json, err := json.MarshalIndent(tmpSTRUCT, "", "\t") // Marshall takes a struct and makes it into JSON
	// if err != nil {
	// 	BOLD_MAGENTA.Print(" *** ")
	// 	BOLD_WHITE.Print(" SHOW_STRUCT NOTE: ")
	// 	BOLD_YELLOW.Print(" You CANNOT pass a struct with EXPORTED fields of type: ")
	// 	BOLD_CYAN.Println("func()")
	// 	BOLD_YELLOW.Println(" MUST make field names start with LOWERCASE or underscore ")
	// 	DO_EXIT("SHOW_STRUCT", "invalid input", true, err)
	// }
	// pretty := string(byte_json)
	// useCOLOR.Println(pretty)

}

//go:linkname memmove runtime.memmove
func memmove(dst unsafe.Pointer, src unsafe.Pointer, length int)

func TEST_DeepCopyEngine(SOURCE any) any {
	FNAME := "DeepCopyEngine"
	PLACEHOLDER()
	//var FNAME = "COPY_DeepCopy_Engine: "
	origTYPE := reflect.TypeOf(SOURCE)
	origKIND := origTYPE.Kind()

	//C.Println("ORIG_TYPE: ", origTYPE)
	Y.Println(" ORIGKIND: ", origKIND)
	W.Println("")

	var is_MAP = origKIND == reflect.Map
	var is_SLICE = origKIND == reflect.Slice
	var is_STRUCT = origKIND == reflect.Struct
	var is_ARRAY = origKIND == reflect.Array
	//var is_INTERFACE = origKIND == reflect.Interface

	if is_MAP == false && is_SLICE == false && is_STRUCT == false && is_ARRAY == false {
		M.Println(FNAME, "  ERROR: Invalid TYPE")
		Y.Println("  SOURCE is NOT a valid MAP, SLICE, or STRUCT")
		DO_EXIT()
		//return nil, nil
	}

	//2. For STRUCTS, we need to do extra handling to deal with Unexported fields
	// This is because the gob package only encodes exported fields by default
	if is_STRUCT {
		//SRCVALS := reflect.ValueOf(SOURCE).Elem()
		SRCVALS := reflect.ValueOf(SOURCE)
		TARGET := reflect.New(origTYPE).Elem()
		TARGVALS := reflect.ValueOf(&TARGET)

		var total_fields = origTYPE.NumField()

		// C.Print("SRCVALS: ")
		// Y.Println(SRCVALS)
		// C.Print("total_fields: ")
		// Y.Println(total_fields)

		//origVALS := reflect.ValueOf(SOURCE)
		// targ_NUM_FIELDS := TARGET.NumField()
		// C.Print("Target TYPE: ")
		// Y.Println(TARGET.Type())
		// Y.Println("Target NUM_FIELDS: ", targ_NUM_FIELDS)
		// SHOW_STRUCT(SOURCE, MAGENTA)

		//var srcElem = SRCVALS.Elem()
		var destElem = TARGVALS.Elem()
		for i := 0; i < total_fields; i++ {
			var src_FTYPE = origTYPE.Field(i)
			var not_Exportable = src_FTYPE.IsExported() == false

			//var dest_FTYPE = TARGET.Type().Field(i)

			// var src_NAME = src_FTYPE.Name
			// var dest_NAME = dest_FTYPE.Name
			// W.Println("")
			// M.Print("not_Exportable: ")
			// Y.Print(not_Exportable)
			// Y.Print(" | ")

			// Y.Print("src_NAME: ")
			// G.Print(src_NAME)
			// Y.Print(" - dest_NAME: ")
			// G.Println(dest_NAME)

			// C.Print(" src_FTYPE: ")
			// Y.Println(src_FTYPE)
			// C.Print("dest_FTYPE: ")
			// Y.Println(dest_FTYPE)

			//var src_FIELD_VAL = reflect.ValueOf(src_FIELD_TYPE)
			//destField := destValue.FieldByName(srcValue.Type().Field(i).Name
			//var src_by_NAME = src_FIELD_VAL.FieldByName(src_FIELD_TYPE.Field(i).Name)
			// dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.Addr().Pointer())).Elem()

			// 3. For NON Exportable fields (lowercase or _ underscore)
			var srcField = SRCVALS.Field(i)
			var dstField = destElem.Field(i)

			//C.Println("srcField: ", srcField)
			if not_Exportable {

				// Use unsafe to copy unexported fields
				dstFieldPtr := unsafe.Pointer(dstField.Addr().Pointer())
				srcFieldPtr := unsafe.Pointer(srcField.Addr().Pointer())
				size := srcField.Type().Size()

				// Copy the memory
				memmove(dstFieldPtr, srcFieldPtr, int(size))

				//3b. For all other fields
			} else {
				TARGET.Field(i).Set(srcField)
				//dstField.Set(srcField)
			}
			return TARGET.Interface()

		} //end of for

		//5. For all the OTHER types. We use gob
		gob.Register(origTYPE)

		// now return the TARGET struct (user will have to use Assertion to get the values)
		return TARGET.Interface()
	}

	//1. Create the enc and DEC onbjects we need. Attach them to a buffer
	var buf bytes.Buffer
	encOBJ := gob.NewEncoder(&buf)
	//decOBJ := gob.NewDecoder(&buf)

	//2. Encode the source value, and make sure we dont have any errors
	if err := encOBJ.Encode(SOURCE); err != nil {
		M.Print(FNAME, "ENCODE Error: ")
		Y.Println(err)
		DO_EXIT()
		//return err, nil
	}

	//3. Create the target object we are copying TO
	//var TARGET = reflect.New(origTYPE)
	//var TARGET = reflect.New(reflect.TypeOf(SOURCE))

	// //4. Now DECODE the buffer into the target object
	// if err := decOBJ.Decode(DST); err != nil {
	// 	M.Print(FNAME)
	// 	W.Print("DECODE, ")
	// 	M.Print(" ERROR: ")
	// 	Y.Println(err)
	// 	DO_EXIT()
	// 	//return err, nil
	// }

	//5. Return the target object
	//return TARGET.Elem().Interface()
	return nil
}
