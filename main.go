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

var uploadsCounter int

func main() {

	uploadsCounter = 1

	config := &serial.Config{Name: "/dev/cu.usbmodem1411", Baud: 9600}
	openedPort, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(openedPort)

	passReader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := passReader.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := passReader.ReadString('\n')
	insta := goinsta.New(username, password)
	if err := insta.Login(); err != nil {
		panic(err)
	}

	for true {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		} else if string(data) == "BANG!" {
			fileName := capturePhoto()
			time.Sleep(1000 * time.Millisecond)
			go upload(insta, fileName)
			time.Sleep(30000 * time.Millisecond)
		}
	}

	insta.Logout()
}

func capturePhoto() string {
	timeStamp := time.Now().Format("200601021504")
	script := `tell application "System Events"
  tell process "Capture One 10" 
  set frontmost to true
  end tell
  keystroke "k" using {command down}
  end tell`
	_, err := exec.Command("osascript", "-e", script).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fileName := []string{"foto/", "", timeStamp, ".JPG"}
	return strings.Join(fileName, "")
}

func upload(insta *goinsta.Instagram, fileName string) {
	insta.UploadPhoto(fileName, "test", int64(uploadsCounter), 87, 0)
	uploadsCounter++
}
