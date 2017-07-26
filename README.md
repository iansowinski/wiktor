## Wymagania

Działa tylko na macu z zainstalowanym pythonem 2 i pythonem 3. wymaga [imagesnap](http://iharder.sourceforge.net/current/macosx/imagesnap/).

pliki odpalamy w tej kolejności:

1. `python uploader.py` - działa z pythonem 2 - należy zalogować się do instagrama i poczekać na komunikat.
2. `python3 server.py` - działa z pythonem 3 - należy poczekać na komunikat o gotowości.
3. `python3 client.py` - działa z pythonem 3 - należy poczekać na komunikat o gotowości przed wpisaniem wiadomości.

Po wpisaniu przez użytkownika wiadomość o treści "True" w konsoli działającego `client.py`, `server.py` odpala kamerkę i robi zdjęcie, przekazując nazwę zdjęcia do serwera `uploader.py`, który wstawia je na instagrama.

## TODO:

- arduino + czujniki
- robienie zdjęć canonem z linii polecań