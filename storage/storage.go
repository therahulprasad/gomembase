package storage

import (
	"errors"
	"github.com/rahulpache/gomembase/commandInterpreter"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Value   string
	Expiry  time.Time
	Settime time.Time
}

var Datastore map[string]Node

func Cleanup() {
	now := time.Now()
	for key, val := range Datastore {
		if val.Expiry.Before(now) {
			delete(Datastore, key)
		}
	}
}

func Set(key, value string, options []commandInterpreter.CommandOption) (int, error) {
	expiryTime := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	for _, val := range options {
		if val.OptionType == "EX" {
			expiry, err := strconv.Atoi(val.OptionValue)
			if err == nil {
				expiryDuration := time.Duration(expiry) * time.Second
				expiryTime = now.Add(expiryDuration)
			} else {
				return -1, errors.New("-ERR value is not an integer or out of range\r\n")
			}
		}
		if val.OptionType == "PX" {
			expiry, err := strconv.Atoi(val.OptionValue)
			if err == nil {
				expiryDuration := time.Duration(expiry) * time.Millisecond
				expiryTime = now.Add(expiryDuration)
			}
		}
		if val.OptionType == "NX" {
			_, ok := Datastore[key]
			if ok == true {
				return -1, errors.New("$-1\r\n")
			}
		}
		if val.OptionType == "XX" {
			_, ok := Datastore[key]
			if ok == false {
				return -1, errors.New("$-1\r\n")
			}
		}
	}

	key = strings.Trim(key, " ")
	Datastore[key] = Node{Value: value, Expiry: expiryTime, Settime: now}

	return 0, nil
}

func Get(key string) (Node, error) {
	key = strings.Trim(key, string(byte(0)))
	val, ok := Datastore[key]
	if ok == true {
		// Check if it is not expired and act accordingly
		if !val.Expiry.IsZero() && val.Expiry.Before(time.Now()) {
			delete(Datastore, key)
		} else {
			return val, nil
		}
	}
	return Node{}, errors.New("Value not found")
}
