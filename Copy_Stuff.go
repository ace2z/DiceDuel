package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"
	"bytes"
	"encoding/gob"
	. "github.com/ace2z/GOGO/Gadgets"
	"reflect"
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
func COPY_DeepCopyEngine(SOURCE any) any {
	FNAME := "DeepCopyEngine"
	PLACEHOLDER()
	//var FNAME = "COPY_DeepCopy_Engine: "
	origTYPE := reflect.TypeOf(SOURCE)
	origKIND := origTYPE.Kind()
	C.Println("ORIG_TYPE: ", origTYPE)
	Y.Println(" ORIGKIND: ", origKIND)
	W.Println("")

	gob.Register(origTYPE)
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

	//1. Create the enc and DEC onbjects we need. Attach them to a buffer
	var buf bytes.Buffer
	encOBJ := gob.NewEncoder(&buf)
	decOBJ := gob.NewDecoder(&buf)

	//2. Encode the source value, and make sure we dont have any errors
	if err := encOBJ.Encode(SOURCE); err != nil {
		M.Print(FNAME, "ENCODE Error: ")
		Y.Println(err)
		DO_EXIT()
		//return err, nil
	}

	//3. Create the target object we are copying TO
	//var TARGET = reflect.New(origTYPE)
	var TARGET = reflect.New(reflect.TypeOf(SOURCE))

	//4. Now DECODE the buffer into the target object
	if err := decOBJ.Decode(TARGET.Interface()); err != nil {
		M.Print(FNAME)
		W.Print("DECODE, ")
		M.Print(" ERROR: ")
		Y.Println(err)
		DO_EXIT()
		//return err, nil
	}

	//4b. For STRUCTS, we need to do extra handling to deal with Unexported fields
	// This is because the gob package only encodes exported fields by default
	if is_STRUCT {
		origVAL := reflect.ValueOf(SOURCE).Elem()
		TARGval := reflect.ValueOf(TARGET).Elem()

		for i := 0; i < origVAL.NumField(); i++ {
			field := origVAL.Type().Field(i)
			if !field.IsExported() {
				TARGval.Field(i).Set(origVAL.Field(i))
			}
		}
	}

	//5. Return the target object
	return TARGET.Elem().Interface()
}

func TEST_DeepCopyEngine(SOURCE any) any {
	FNAME := "DeepCopyEngine"
	PLACEHOLDER()
	//var FNAME = "COPY_DeepCopy_Engine: "
	origTYPE := reflect.TypeOf(SOURCE)
	origKIND := origTYPE.Kind()
	C.Println("ORIG_TYPE: ", origTYPE)
	Y.Println(" ORIGKIND: ", origKIND)
	W.Println("")

	gob.Register(origTYPE)
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
		TARGET := reflect.New(origTYPE).Elem()
		origVALS := reflect.ValueOf(SOURCE)

		for i := 0; i < origTYPE.NumField(); i++ {
			field := origVALS.Field(i)
			var not_exported = field.CanSet() == false
			var regular = field.CanInterface()

			// For regular fields
			if regular {
				destField := TARGET.Field(i)
				if destField.CanSet() {
					destField.Set(field)
				// Otherwise if this is an EXPORTABLE field
				} else {
					M.Print(FNAME, " ERROR: ")
					Y.Println("Field cannot be set")
					DO_EXIT()
				}
			}
			
			var not_exported = !field.IsExported()
			if not_exported {
			 	TARGval.Field(i).Set(origVAL.Field(i))
			}
			TARGET.Field(i).Set(origVALS.Field(i))
		}

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
