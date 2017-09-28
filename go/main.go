package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	whereIAm()
	config := &serial.Config{Name: "/dev/cu.wchusbserial1410", Baud: 9600}
	openedPort, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(openedPort)
	for true {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		} else if string(data) == "BANG!" {
			go snap()
			go sendCommand()
		}
	}
}

func pwd() string {
	output, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return string(output)
}

func sendCommand() {
	_, err := exec.Command("osascript", "-e", "tell application \"Google Chrome\"\nactivate\ntell application \"System Events\"\nkeystroke \"r\" using {command down}\nend tell\nend tell").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func whereIAm() {
	fmt.Println("your output will be stored in:", pwd())
}

func snap() {
	now := time.Now()
	cmd := "imagesnap"
	fileName := []string{now.Format("20060102150405"), ".jpg"}
	fileNameFormatted := strings.Join(fileName, "")
	args := []string{"-w", "1", fileNameFormatted}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
