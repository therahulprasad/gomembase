package main
import "fmt"
import "net"

const (
	RECV_BUF_LEN = 1024
)

func main() {
	fmt.Println("Hello world")
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
		if conn == nil {
			fmt.Println("Connection received : conn is nil")
		} else {
			fmt.Println("Connection received")
		}
		for {
			buf := make([]byte, RECV_BUF_LEN)
			n, err := conn.Read(buf)
			if err != nil {
				println("Error reading:", err.Error())
				return
			}
			println("received ", n, " bytes of data =", string(buf))
		 
			//send reply
			_, err = conn.Write(buf)
			if err != nil {
				println("Error send reply:", err.Error())
			}else {
				println("Reply sent")
			}
		}	
	}
}

