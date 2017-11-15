# Backend for art instalation by Wiktor Wolski

Files:

- **arduino_controller.c** - arduin
- **main.go** - main program, compile it and run on mac attached to arduino running *arduino_controller.c*, Canon EOS 600D and Capture One Pro 10.

1. ```go build mian.go```
2. ```mkdir foto```
3. Open Capture One Pro 10
4. Connect your camera
5. Set target folder in Capture One Pro to ```<this repo's directory>/foto```
6. Set image size to some JPG
7. Connect arduino with PIR sensor
8. ```cat arduino_controller.c | pbcopy```
9. Open arduino studio and paste clipboard content to new sketch
10. Check if you have correct value in ```pirPin``` (it should equal pin to whitch you connected PIR sensor)
11. Click ```upload``` in arduino studio
12. ```./main```
