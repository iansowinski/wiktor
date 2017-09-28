## Wymagania

Działa tylko na mac os sierra z zainstalowanym pythonem 2 i go. Wymaga [imagesnap](http://iharder.sourceforge.net/current/macosx/imagesnap/).

- [Sterowniki](https://github.com/adrianmihalko/ch340g-ch34g-ch34x-mac-os-x-driver)
- [Arduino IDE](https://www.arduino.cc/en/Guide/ArduinoNano)
- [PIR](https://forbot.pl/blog/kurs-arduino-ii-4-przerwania-kontaktron-czujnik-pir-id16792)
- [Instrukcja działania czujnika PIR](https://www.youtube.com/watch?v=63TR_3kn76U)

## Co jest czym

`main.go` - po odpaleniu / skompilowaniu czyta z serialportu i w momencie gdy dostanie stringa "BANG!" odpala zdjęcie z kamerki w komputerze (`snap()`) oraz wysyła keystroke do odpowieniej aplikacji (do rozbudowania w `sendCommand()`). W założeniu również wysyła następnie zdjęcie na facebooka, lub sygnał do `uploader.py` za pomocą socketu tcp/ip)

`uploader.py` - serwer który wrzuca zdjęcie gdy dostanie jego ścieżkę przez socket