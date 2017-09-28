package main

import (
	"bufio"
  "log"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
  "github.com/tarm/serial"
)

func main() {
	whereIAm()
  config := &serial.Config{Name: "/dev/cu.wchusbserial1410", Baud: 9600}
  s, err := serial.OpenPort(config)
  if err != nil {
      log.Fatal(err)
  }
  r := bufio.NewReader(s)
  for true {
    data, _, err := r.ReadLine()
    if err != nil {
        log.Fatal(err)
    } else if string(data) == "BANG!" {
      go snap()
    }
  }
}


func whereIAm() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("your output will be stored in: %s\n", out)
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
