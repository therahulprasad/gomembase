package storage

import (
	"github.com/therahulprasad/gomembase/commandInterpreter"
	"testing"
)

func setup() {
	Datastore = make(map[string]Node)
}

func TestSet(t *testing.T) {
	setup()

	commandOptions := []commandInterpreter.CommandOption{commandInterpreter.CommandOption{}}
	status, err := Set("hello", "world", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set hello world failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "EX", OptionValue: "100"}}
	status, err = Set("hello", "world", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set hell world EX 100 failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "PX", OptionValue: "100"}}
	status, err = Set("hello", "world", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set hell world PX 100 failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "NX"}}
	status, err = Set("helloNew", "worldNew", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set helloNew worldNew NX failed")
	}

	// hello world is already set NX should cause error
	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "NX"}}
	status, err = Set("hello", "world", commandOptions)
	if err == nil || status != -1 {
		t.Error("Set hello world NX failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "XX"}}
	status, err = Set("hello", "world", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set hell world XX failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "XX"}}
	status, err = Set("foo", "bar", commandOptions)
	if err == nil || status != -1 {
		t.Error("Expecting error, foo=>bar is not already set.Set foo bar XX")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "XX"}, commandInterpreter.CommandOption{OptionType: "EX", OptionValue: "100"}}
	status, err = Set("hello", "world", commandOptions)
	if err != nil || status != 0 {
		t.Error("Set hell world XX EX 100 failed")
	}

	commandOptions = []commandInterpreter.CommandOption{commandInterpreter.CommandOption{OptionType: "EX", OptionValue: "what"}}
	status, err = Set("hello", "world", commandOptions)
	if err == nil || status != -1 {
		t.Error("Expecting error, Non numeric EX provided")
	}
}

func TestGet(t *testing.T) {
	// Previous function is already run (hello must have world set)
	_, err := Get("hello")
	if err != nil {
		t.Error("Get hello failed")
	}

	_, err = Get("unknown")
	if err == nil {
		t.Error("Get unknown should throw error")
	}
}
