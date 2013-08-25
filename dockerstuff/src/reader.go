package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"bufio"
	"bytes"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	fmt.Println("Listening for tcp traffic on 0.0.0.0:6666")
	if err != nil {
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go JuliaFunc(conn)
	}
}

func JuliaFunc(conn net.Conn) {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	f, err := ioutil.TempFile("/tmp", "onlinerepl")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		b, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
		got := bytes.TrimSpace(b)
		if string(got) == "exit()" {
			fmt.Println("Got some exit()")
			break
		}
		fmt.Println("got some bytes", got)
		fmt.Println("Writing to a file")
		f.Write(got)
	}

	fmt.Println("running the file on julia")
	cmd := exec.Command("julia", f.Name())
	out, err := cmd.Output()
	fmt.Println("Got this from julia:", out)
	if err != nil {
		fmt.Println("Got an error but I'm sure it's fine:", err.Error())
	}
	//send reply
	conn.Write(out)
}
