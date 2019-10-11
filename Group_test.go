package structs

import (
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

// TestGroup () tests function Group ().
func TestGroup (t *testing.T) {
	str.PrintEtr ("Test started.", "nte", "TestGroup ()")

	// Ensuring function Group () fails, when trying to group non-struct kinds. ..1..
	// {
	gA, errA := Group ("someFiled", false, "1", &struct {name string}{"Foo Bar"},
		1024105)
	if gA != nil || errA == nil {
		str.PrintEtr ("Test failed! Ref: 0", "err", "TestGroup ()")
		t.FailNow ()
	}
	// ..1.. }

	// Ensuring function Group () fails, when trying to group with an invalid field
	// name. ..1.. {
	e0 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e1 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e2 := testType3 {"Birth", "22/04", "Lag"}

	gC, errC := Group ("_location", &e0, &e1, &e2)
	gD, errD := Group ("location", &e0, &e1, &e2)
	gE, errE := Group ("$d", &e0, &e1, &e2)
	gF, errF := Group ("%%", &e0, &e1, &e2)
	gG, errG := Group ("", &e0, &e1, &e2)

	if gC != nil || gD != nil || gE != nil || gF != nil || gG != nil || errC == nil ||
		errD == nil || errE == nil || errF == nil || errG == nil {

		str.PrintEtr ("Test failed! Ref: 2", "err", "TestGroup ()")
		t.FailNow ()
	}
	// ..1.. }

	// Ensuring function Group () fails, when trying to group with an uncommon field.
	// ..1.. {
	e3 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e4 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e5 := testType3 {"Birth", "22/04", "Lag"}

	resultI, errI := Group ("Name", &e3, &e4, &e5)
	if resultI != nil {
		str.PrintEtr ("Test failed! Ref: 4", "err", "TestGroup ()")
		t.FailNow ()
	}
	if errI == nil {
		str.PrintEtr ("Test failed! Ref: 5", "err", "TestGroup ()")
		t.FailNow ()
	}

	// Ensuring function Group () makes accurate grouping. ..1.. {
	e6 := testType1 {"Ibrahim Qamardeen", 22, "Oyo"}
	e7 := testType2 {"Qamarian", "IBZ", "Oyo"}
	e8 := testType3 {"Birth", "22/04", "Lag"}

	resultK, errK := Group ("Location", &e6, &e7, &e8)
	if errK != nil {
		str.PrintEtr ("Test failed! Ref: 7", "err", "TestGroup ()")
		t.FailNow ()
	}

	if !((resultK ["Oyo"][0] == &e6 && resultK ["Oyo"][1] == &e7 &&
		resultK ["Lag"][0] == &e8) ||

		(resultK ["Oyo"][0] == &e7 && resultK ["Oyo"][1] == &e6 &&
		resultK ["Lag"][0] == &e8)) {

		str.PrintEtr ("Test failed! Ref: 8", "err", "TestGroup ()")
		t.FailNow ()
	}
	// ..1.. }

	str.PrintEtr ("Test passed.", "std", "TestGroup ()")
}

type testType1 struct {
	Name     string
	Age      int
	Location string
}

type testType2 struct {
	Name     string
	Ceo      string
	Location string
}

type testType3 struct {
	Theme    string
	Account  string
	Location string
}
