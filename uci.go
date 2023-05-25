package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//In the provided code snippet, the uci function takes a function tell as a parameter.
//The type of the tell parameter is func(text ...string),
//which means it is a function that takes a variable number of string arguments.

var tell = mainTell //set default tell

func uci(frGUI chan string, myTell func(text ...string)) {

	tell = myTell

	tell("Hello from uci")

	//2 channels from the engine
	frEng, toEng := engine()

	quit := false
	cmd := ""
	bm := ""

	for quit == false {
		select {
		case cmd = <-frGUI:
		case bm = <-frEng:
			handleBm(bm)
			continue
		}
		switch cmd {
		case "uci":
			handleUci()
		case "debug on":
			handleDebugOn()
		case "isready":
			handleIsReady()
		case "setoption":
			handleSetoption()
		case "stop":
			handleStop(toEng)
		case "quit", "q":
			quit = true
			continue
		}
	}
}

func handleDebugOn() {
	tell("info string debug on")
}

func handleSetoption(option []string) {
	cmd := strings.Join(option, " ")
	tell("info string set option", cmd)
	tell("testing")
}

func handleUci() {
	tell("id name Ashish")
	tell("id author Ashish")
	tell("uciok")
}

func handleIsReady() {
	tell("readyok")
}

func handleBm(bm string) {
	tell(bm)
}

func handleStop(toEng chan string) {
	toEng <- "stop"
}

// go routine waits for commands and sends to uci : standard way to use anonymous func as go routines
func input() chan string {
	line := make(chan string)
	go func() {
		var reader *bufio.reader
		//A buffered reader that reads from the standard input (os.Stdin).
		reader = bufio.NewReader(os.Stdin)
		for {
			//To read a line of text from the standard input until a newline character ('\n') is encountered
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()
	return line
}

// mainTell is the tell function run when we not run test
func mainTell(text ...string) {

	//Empty string
	toGUI := ""

	//In each iteration, the value of t is concatenated to the toGUI string using the += operator.
	//This appends the current value of t to the end of toGUI.
	// range gives me index and value , _ means index is not interested in (blank identifier)
	for _, t := range text {
		toGUI += t
	}

	fmt.Println(toGUI)

}
