package main

import (
	"bufio"
	"fmt"
	"github.com/ahmdrz/goinsta"
	"github.com/tarm/serial"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Global counter for uploadId
var uploadsCounter int

func main() {

	uploadsCounter = 1
	whereIAm()

	// Serialport settings
	config := &serial.Config{Name: "/dev/cu.wchusbserial1410", Baud: 9600}
	openedPort, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(openedPort)

	// Instagram login
	passReader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := passReader.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := passReader.ReadString('\n')
	insta := goinsta.New(username, password)
	if err := insta.Login(); err != nil {
		panic(err)
	}
	// Main loop
	for true {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		} else if string(data) == "BANG!" {
			go snap(insta)
			go sendCommand()
		}
	}

	insta.Logout()
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

func snap(insta *goinsta.Instagram) {
	now := time.Now()
	cmd := "imagesnap"
	fileName := []string{now.Format("20060102150405"), ".jpg"}
	fileNameFormatted := strings.Join(fileName, "")
	args := []string{"-w", "1", fileNameFormatted}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	upload(insta, fileNameFormatted, int64(uploadsCounter))
	uploadsCounter++
}

func upload(insta *goinsta.Instagram, fileName string, uploadId int64) {
	insta.UploadPhoto(fileName, "test", uploadId, 87, 0)
}
