package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func readLine() string {
	fmt.Printf(">> ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return line
}

func sendMessage(conn net.Conn, message string) {
	message = strings.TrimSpace(message)
	conn.Write([]byte(message))
}

func sendMessages(messages chan string, address string) {
	/* establish a connection
	   if connection fails, panic
	   read message from messages
	   send
	*/
}

func sayClient(address string) {
	/*
		create messages channel
		call coroutine to read messages channel and send

		loop forever
		    get line from terminal
			put line in message channel
	*/
}

func sayMessages(messages chan string) {
	var message string
	for {
		message = <-messages
		say(message)
	}
}

func sayServer(address string) {
	messages := make(chan string)
	go sayMessages(messages)

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening at %s...\n", listener.LocalAddr().String())
	buffer := make([]byte, 1024)
	for {
		nBytesReceived, remoteAddr, err := listener.ReadFromUDP(buffer)
		if err == nil {
			message := string(buffer[0:nBytesReceived])
			fmt.Printf("%s -> \"%s\"\n", remoteAddr.String(), message)
			messages <- message
		}
	}
}

func say(message string) {
	cmd := exec.Command("say", message)
	cmd.Output()
	return
}

func parseArgs() (string, string, error) {
	usageMsg := fmt.Sprintf("Usage: %s <\"listen\"|\"send\"> <host> <port>\n", os.Args[0])
	err := errors.New(usageMsg)
	if len(os.Args) < 3 {
		return "", "", err
	}
	processType, host, port := os.Args[1], os.Args[2], os.Args[3]
	if processType != "send" && processType != "listen" {
		return "", "", err
	}
	address := net.JoinHostPort(host, port)
	return processType, address, nil
}

func main() {
	processType, addr, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if processType == "listen" {
		sayServer(addr)
	} else {
		sayClient(addr)
	}
}
