package main

import (
  "github.com/ahmdrz/goinsta"
  "reflect"
  "fmt"
)

func main() {
  insta := goinsta.New("iansowinski@gmail.com", "multiflash56")
  if err := insta.Login(); err != nil {
    panic(err)
  }
  fmt.Println(reflect.TypeOf(insta).String())
  upload(insta)
  defer insta.Logout()
}
func upload(insta *goinsta.Instagram, uploadId int64) {
  insta.UploadPhoto("20170928181912.jpg", "test", int64(3), 87, 0)
}