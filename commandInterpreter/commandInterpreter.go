package commandInterpreter

import (
	"errors"
	"strings"
)

type CommandOption struct {
	OptionType, OptionValue string
}

type Command struct {
	CommandType  string
	Options      []CommandOption
	CommandValue map[string]string
}

func ParseCommand(string_buf string) (Command, error) {
	var commandValue = make(map[string]string)
	command_split := strings.Split(string_buf, " ")
	numCommandSplit := len(command_split)
	commandOptions := []CommandOption{CommandOption{}}

	if numCommandSplit > 0 {
		// Compare commands
		if strings.ToLower(command_split[0]) == "set" {
			// Check if key and value is supplied
			if numCommandSplit >= 3 {
				for i := 1; i < numCommandSplit; i++ {
					if numCommandSplit > i+1 {
						if strings.ToLower(command_split[i]) == "ex" {
							if len(strings.Trim(command_split[i+1], " ")) > 0 {
								commandOptions = []CommandOption{CommandOption{OptionType: "EX", OptionValue: strings.Trim(command_split[i+1], " ")}}
								i++
							}
						} else if strings.ToLower(command_split[i]) == "px" {
							if len(strings.Trim(command_split[i+1], " ")) > 0 {
								commandOptions = []CommandOption{CommandOption{OptionType: "PX", OptionValue: strings.Trim(command_split[i+1], " ")}}
								i++
							}
						} else if len(strings.Trim(command_split[i], " ")) > 0 {
							commandValue["key"] = strings.Trim(command_split[i], " ")
							value := strings.Trim(command_split[i+1], " ")
							if len(value) > 0 {
								var invertedComma byte = 0
								if value[0] == '"' {
									invertedComma = '"'
								} else if value[0] == '\'' {
									invertedComma = '\''
								}

								if invertedComma != 0 {
									broken := false
									value = ""
									first := true
									for j := i + 1; j < numCommandSplit; j++ {
										lenValue := len(command_split[j])
										if first == true {
											first = false
											value += command_split[j]
										} else {
											value += " " + command_split[j]
										}

										i++
										if lenValue >= 1 {
											if lenValue == 1 && command_split[j][lenValue-1] == invertedComma {
												broken = true
												break
											} else if command_split[j][lenValue-1] == invertedComma && command_split[j][lenValue-2] != '\\' {
												broken = true
												break
											}
										}
									}
									if broken == true {
										commandValue["value"] = strings.Trim(value, string(invertedComma))
									} else {
										return Command{}, errors.New("-ERR Protocol error: unbalanced quotes in request\r\n")
									}
								} else {
									commandValue["value"] = value
									i++
								}
							} else {
								return Command{}, errors.New("-ERR wrong number of arguments for 'set' command\r\n")
							}
						} else if command_split[i] == "" {

						} else {
							return Command{}, errors.New("-ERR wrong number of arguments for 'set' command\r\n")
						}
					} else if numCommandSplit > i {
						if strings.ToLower(command_split[i]) == "nx" {
							commandOptions = []CommandOption{CommandOption{OptionType: "NX"}}
						} else if strings.ToLower(command_split[i]) == "xx" {
							commandOptions = []CommandOption{CommandOption{OptionType: "XX"}}
						}
					}
				}
			} else {
				return Command{}, errors.New("-ERR wrong number of arguments for 'set' command\r\n")
			}

			return Command{CommandType: "SET", CommandValue: commandValue, Options: commandOptions}, nil
		} else if strings.ToLower(command_split[0]) == "get" {
			if len(command_split) >= 2 {
				if len(strings.Trim(command_split[1], " ")) > 0 {
					commandValue["key"] = strings.Trim(command_split[1], " ")
				} else {
					return Command{}, errors.New("-ERR wrong number of arguments for 'get' command\r\n")
				}
			} else {
				return Command{}, errors.New("-ERR wrong number of arguments for 'get' command\r\n")
			}
			return Command{CommandType: "GET", CommandValue: commandValue}, nil
		}
	}
	return Command{}, errors.New("")
}
