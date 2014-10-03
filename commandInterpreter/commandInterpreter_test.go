package commandInterpreter

import (
	"testing"
)

func TestParseSetCommand(t *testing.T) {
	//** TDD Implemented **
	command, err := ParseCommand("SET hello world")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
		t.Error("SET command failure.")
	}

	command, err = ParseCommand("set hello world")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
		t.Error("SET command failure. Case sensitivity")
	}

	command, err = ParseCommand("set hello world  ")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
		t.Error("SET command failure. Space at the end")
	}

	command, err = ParseCommand("set hello 'world'")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
		t.Error("SET command failure. Single inverted comma")
	}

	command, err = ParseCommand("set hello 'foo bar'")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo bar" {
		t.Error("SET command failure. Single inverted comma, multiword")
	}

	command, err = ParseCommand("set hello 'foo    bar'")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo    bar" {
		t.Error("SET command failure. Single inverted command, multiword, large spacing")
	}

	command, err = ParseCommand("set hello 'foo    bar   '")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo    bar   " {
		t.Error("SET command failure. Single inverted comma, multiword, spacing at the end")
	}

	command, err = ParseCommand("set hello \"foo bar\"")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo bar" {
		t.Error("SET command failure. Double inverted comma, multiword")
	}

	command, err = ParseCommand("set hello \"foo    bar\"")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo    bar" {
		t.Error("SET command failure.Double inverted comma, multiword, spacing")
	}

	command, err = ParseCommand("set hello \"foo    bar   \"")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "foo    bar   " {
		t.Error("SET command failure.Double inverted comma,multiword, spacing at the end")
	}

	command, err = ParseCommand("set hello world EX 100")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" || command.Options[0].OptionType != "EX" || command.Options[0].OptionValue != "100" {
		t.Error("SET command failure. EX")
	}

	command, err = ParseCommand("set hello world PX 100")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" || command.Options[0].OptionType != "PX" || command.Options[0].OptionValue != "100" {
		t.Error("SET command failure. PX")
	}

	command, err = ParseCommand("set hello world NX")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" || command.Options[0].OptionType != "NX" {
		t.Error("SET command failure. NX")
	}

	command, err = ParseCommand("set hello world XX")
	if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" || command.Options[0].OptionType != "XX" {
		t.Error("SET command failure. XX")
	}

	

	_, err = ParseCommand("set hello")
	if err == nil {
		t.Error("SET command failure. Expecting error. value not set")
	}

	_, err = ParseCommand("set hello   ")
	if err == nil {
		t.Error("SET command failure. Expecting error. value not set, empty space")
	}

	//** TDD: Implement following **
	/*
		command, err = ParseCommand("set hello ' A '")
		if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != " A " {
			t.Error("SET command failure.")
		}

		command, err = ParseCommand("set hello world EX 100 NX")
		if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" || command.Options[0].OptionType != "EX" || command.Options[0].OptionValue != "100" || command.Options[1].OptionType == "NX" {
			t.Error("SET command failure. MultiOption EX and NX")
		}

		command, err = ParseCommand(" set hello world")
		if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
			t.Error("SET command failure.")
		}

		command, err = ParseCommand("set   hello   world")
		if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
			t.Error("SET command failure.")
		}

		command, err = ParseCommand("set hello   world")
		if err != nil || command.CommandType != "SET" || command.CommandValue["key"] != "hello" || command.CommandValue["value"] != "world" {
			t.Error("SET command failure.")
		}
	*/
}

func TestGetCommand(t *testing.T) {
	command, err := ParseCommand("GET hello")
	if err != nil || command.CommandType != "GET" || command.CommandValue["key"] != "hello" {
		t.Error("GET command failure.")
	}

	command, err = ParseCommand("get hello")
	if err != nil || command.CommandType != "GET" || command.CommandValue["key"] != "hello" {
		t.Error("GET command failure. Case sensitivity")
	}

	command, err = ParseCommand("GET hello  ")
	if err != nil || command.CommandType != "GET" || command.CommandValue["key"] != "hello" {
		t.Error("GET command failure.")
	}

	_, err = ParseCommand("GET")
	if err == nil {
		t.Error("GET command failure. No argument")
	}

	// ======= TDD Implement following =========
	/*
		command, err = ParseCommand(" GET hello  ")
		if err != nil || command.CommandType != "GET" || command.CommandValue["key"] != "hello" {
			t.Error("GET command failure. Prefix spacing")
		}
	*/
}

func TestCommonCommand(t *testing.T) {
	_, err := ParseCommand("")
	if err == nil {
		t.Error("Nothing sent, expecting error")
	}

	_, err = ParseCommand("    ")
	if err == nil {
		t.Error("Blank space sent, expecting error")
	}

	_, err = ParseCommand("Non command")
	if err == nil {
		t.Error("Non command sent, expecting error")
	}

	_, err = ParseCommand("\n")
	if err == nil {
		t.Error("Carriage return sent, expecting error")
	}
}
