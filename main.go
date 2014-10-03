package main

import (
	"fmt"
	"github.com/therahulprasad/gomembase/commandInterpreter"
	"github.com/therahulprasad/gomembase/config"
	"github.com/therahulprasad/gomembase/storage"
	"net"
	"strconv"
	"time"
)

const (
	RECV_BUF_LEN = 1024
)

func main() {
	storage.Datastore = make(map[string]storage.Node)
	defer hello()

	// Run garbage cleaner as a separate routine
	go garbageCleaner()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error 1")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error 2")
			continue
		}
		fmt.Println("Connection received")

		// Handle every connection as separte routine
		go handleConnection(conn)
	}
}

func hello() {
	fmt.Println("World")
}

// Cleanup expired value every minute
func garbageCleaner() {
	for _ = range time.Tick(1 * time.Minute) {
		storage.Cleanup()
	}
}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, RECV_BUF_LEN)
		n, err := conn.Read(buf)
		if err != nil {
			println("Error reading:", err.Error())
			return
		}
		//println("received ", n, " bytes of data =", string(buf[:n-2]))

		// TODO FIX Command Interpreter
		command, err := commandInterpreter.ParseCommand(string(buf[:n-2]))
		if err != nil {
			reply(conn, []byte(err.Error()))
			continue
		}

		if command.CommandType == "SET" {
			// Send proper data to SET command
			_, err := storage.Set(command.CommandValue["key"], command.CommandValue["value"], command.Options)
			if err != nil {
				reply(conn, []byte(err.Error()))
			} else {
				reply(conn, []byte("+OK\r\n"))
			}
		} else if command.CommandType == "GET" {
			data, err := storage.Get(command.CommandValue["key"])
			if err != nil {
				//error_message := err.Error() + "\r\n"
				error_message := "$-1\r\n"
				reply(conn, []byte(error_message))
				continue
			}
			// TODO implement output formatter as separate package
			output_length := len(data.Value)
			output := "$" + strconv.Itoa(output_length) + "\r\n" + data.Value + "\r\n"
			reply(conn, []byte(output))
		}
	}
}

func reply(conn net.Conn, buf []byte) {
	//_,err := conn.Write(reverse(buf))
	_, err := conn.Write(buf)
	if err != nil {
		if config.Debug == true {
			println("Error send reply:", err.Error())
		}
	} else {
		if config.Debug == true {
			//println("Reply sent")
		}
	}
}

func reverse(buf []byte) (buff []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = buf[len(buf)-i]
	}
	return buf
}
