package env

import (
	. "testing"
)

func TestBasicFunctions(t *T) {
	if _, ok := Get("SOME_TEST_VAR"); ok {
		t.Errorf("$SOME_TEST_VAR should not have been present.")
	}

	Set("SOME_TEST_VAR", "the value")
	if val, ok := Get("SOME_TEST_VAR"); !ok {
		t.Errorf("$SOME_TEST_VAR should have been present.")
	} else if val != "the value" {
		t.Errorf("$SOME_TEST_VAR should have been \"the value\"")
	}

	Unset("SOME_TEST_VAR")
	if _, ok := Get("SOME_TEST_VAR"); ok {
		t.Errorf("$SOME_TEST_VAR should not have been present.")
	}
}

func TestGetInt(t *T) {
	val, ok, err := GetInt("TEST_GET_INT")
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Errorf("$TEST_GET_INT should not have been found")
	}

	Set("TEST_GET_INT", "42")
	val, ok, err = GetInt("TEST_GET_INT")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("$TEST_GET_INT should have been found")
	}
	if val != 42 {
		t.Errorf("$TEST_GET_INT should have been 42")
	}

	Set("TEST_GET_INT", "-1234")
	val, ok, err = GetInt("TEST_GET_INT")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("$TEST_GET_INT should have been found")
	}
	if val != -1234 {
		t.Errorf("$TEST_GET_INT should have been -1234")
	}

	Set("TEST_GET_INT", "whatever")
	val, ok, err = GetInt("TEST_GET_INT")
	if err == nil {
		t.Errorf("$TEST_GET_INT should not have been parsed")
	}
	if !ok {
		t.Errorf("$TEST_GET_INT should have been found")
	}
}

func TestGetIntDefault(t *T) {
	val, ok, err := GetIntDefault("TEST_GET_INT_DEFAULT", 23)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Errorf("$TEST_GET_INT_DEFAULT should not have been found")
	}
	if val != 23 {
		t.Errorf("$TEST_GET_INT_DEFAULT should have defaulted to 23")
	}

	Set("TEST_GET_INT_DEFAULT", "42")
	val, ok, err = GetIntDefault("TEST_GET_INT_DEFAULT", 24)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("$TEST_GET_INT_DEFAULT should have been found")
	}
	if val != 42 {
		t.Errorf("$TEST_GET_INT_DEFAULT should have been 42")
	}

	Set("TEST_GET_INT_DEFAULT", "-1234")
	val, ok, err = GetIntDefault("TEST_GET_INT_DEFAULT", 24)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("$TEST_GET_INT_DEFAULT should have been found")
	}
	if val != -1234 {
		t.Errorf("$TEST_GET_INT_DEFAULT should have been -1234")
	}

	Set("TEST_GET_INT_DEFAULT", "whatever")
	val, ok, err = GetIntDefault("TEST_GET_INT_DEFAULT", 4444)
	if err == nil {
		t.Errorf("$TEST_GET_INT_DEFAULT should not have been parsed")
	}
	if !ok {
		t.Errorf("$TEST_GET_INT_DEFAULT should have been found")
	}
	if val != 4444 {
		t.Errorf("$TEST_GET_INT_DEFAULT should have defaulted to 4444")
	}
}

func TestMustGet(t *T) {
	Set("TEST_MUST_GET", "testing")
	val := MustGet("TEST_MUST_GET")
	if val != "testing" {
		t.Errorf("$TEST_MUST_GET should have been \"testing\"")
	}

	Unset("TEST_MUST_GET")
	defer func() {
		recover()
	}()
	val = MustGet("TEST_MUST_GET")
	t.Errorf("Should not have reached here")
}

func TestMustGetInt(t *T) {
	expectPanic(t, func () {
		MustGetInt("TEST_MUST_GET_INT")
	})

	Set("TEST_MUST_GET_INT", "42")
	val := MustGetInt("TEST_MUST_GET_INT")
	if val != 42 {
		t.Errorf("$TEST_MUST_GET_INT should have been 42")
	}

	Set("TEST_MUST_GET_INT", "whatever")
	expectPanic(t, func () {
		MustGetInt("TEST_MUST_GET_INT")
	})
}

func TestMustGetIntDefault(t *T) {
	val := MustGetIntDefault("TEST_MUST_GET_INT_DEFAULT", 23)
	if val != 23 {
		t.Errorf("$TEST_MUST_GET_INT_DEFAULT should have defaulted to 23")
	}

	Set("TEST_MUST_GET_INT_DEFAULT", "42")
	val = MustGetIntDefault("TEST_MUST_GET_INT_DEFAULT", 24)
	if val != 42 {
		t.Errorf("$TEST_MUST_GET_INT_DEFAULT should have been 42")
	}

	Set("TEST_MUST_GET_INT_DEFAULT", "whatever")
	expectPanic(t, func() {
		MustGetIntDefault("TEST_MUST_GET_INT_DEFAULT", 4444)
	})
}

func expectPanic(t *T, f func()) {
	defer func() {
		recover()
	}()

	f()

	t.Errorf("Expected a panic, no panic was encountered.")
}
