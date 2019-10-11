package structs

import (
	"errors"
	"gopkg.in/qamarian-dtp/err.v0" // v0.4.0
	"reflect"
	"regexp"
)

// Group () groups a set of structs, based on their values of a particular field.
//
// Example:
//
// Imagine having the following types:
//
// 	type type0 struct {
//		Name string
//		Location string
//	}
//
// 	type type1 struct {
//		CompanyType string
//		Name string
//	}
//
//	type type2 struct {
//		Name string
//		Age int
//	}
//
// We could then create the following elements and group them, based on the values of
// their "Name" field:
//
//	e0 := struct {Name string; BotanicalName string}{"Apple", "Apple de Compino"}
//	e1 := type0 {"Ibrahim Oladipupo Qamardeen", "Nigeria"}
//	e2 := type0 {"Ibrahim Oladipupo Qamardeen", "Lagos"}
//	e3 := type1 {"Inc", "Apple"}
//	e4 := type2 {"Sammy", 2}
//
// 	grouping, err := structs.Group ("Name", &e0, &e1, &e2, &e3, &e4)
//
// 	fmt.Println (grouping, err) // map[Apple:[&e0 &e3] Ibrahim Oladipupo Qamardeen:[&e1 &e2] Sammy:[&e4]] <nil>
//
// Notes:
//
// 	1. Elements can only be group based on exported fields. E.g.: "Name" not "name",
// 		"Age" not "age",
// 	2. The grouping field must be common to all elements.
func Group (field string, elements ... interface {}) (o map[interface {}][]interface {},
	e error) {

	if _Group_buggyFunc == true {
		return nil, errors.New ("An error occured. Function might be buggy.")
	}

	// All inputs validation. ..1.. {

	// Function definition. ..2.. {
	validate := func (elements []interface{}) (err error) {
		for _, someElement := range elements {
			pointer := reflect.ValueOf (someElement)
			if pointer.Kind () != reflect.Ptr {
				return errors.New ("An element is not a pointer.")
			}
			if pointer.Elem ().Kind () != reflect.Struct {
				return errors.New ("A pointer is not a struct " +
					"pointer.")
			}
			if pointer.Elem ().FieldByName (field).IsValid () == false {
				return errors.New ("An element does not have " +
					"that field.")
			}
		}
		return nil
	}
	// ..1.. }

	errX := validate (elements)
	if errX != nil {
		return nil, err.New ("An invalid element supplied.", nil, nil, errX)
	}

	if _Group_fieldPattern.MatchString (field) == false {
		return nil, errors.New ("Invalid field name.")
	}
	// ..2.. }

	o = map[interface {}][]interface {} {}

	for _, anElement := range elements {
		fieldValue := reflect.ValueOf (anElement).Elem ().FieldByName (field)

		fieldElements, okX := o [fieldValue.Interface ()]
		if okX == false {
			fieldElements = []interface{} {}
		}
		fieldElements = append (fieldElements, anElement)
		o [fieldValue.Interface ()] = fieldElements
	}

	return o, nil
}

var (
	_Group_fieldPattern *regexp.Regexp
	_Group_buggyFunc bool = false
)

func init () {
	// | --
	var errX error
	// -- |

	_Group_fieldPattern, errX = regexp.Compile ("^[A-Z][A-Za-z0-9_]*$")
	if errX != nil {
		_Group_buggyFunc = true
	}
}
