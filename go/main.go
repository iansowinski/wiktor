package main

import (
  "fmt"
  "time"
  "bufio"
  "os"
  "os/exec"
  "strings"
)


func main(){
  whereIAm()
  commandCatcher()
}

func commandCatcher() {
  reader := bufio.NewReader(os.Stdin)
  for true {
    fmt.Print("command: ")
    text, _ := reader.ReadString('\n')
    if text == "snap\n" {
      go snap()
    }
    if text == "exit\n" {
      break
    }
  }
}

func whereIAm(){
  out, err := exec.Command("pwd").Output()
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  fmt.Printf("your output will be stored in: %s\n", out)
}

func snap(){
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