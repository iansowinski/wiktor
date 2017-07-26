import socket
import os
from datetime import datetime

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('localhost', 10001))
server.listen(1)

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('localhost', 9999))

print("You are now free to start client.py by running `python3 client.py`")

while True:
    connection, client_address = server.accept()
    try:
        while True:
            data = connection.recv(1024).decode('utf-8')
            if data == "True":
                timestamp = datetime.now().strftime("%Y-%m-%d-%H-%M-%S")
                print("Snap!")
                os.system("cd images && imagesnap -w 1 " + timestamp + ".jpg > /dev/null")
                client.sendall(timestamp.encode('utf-8'))
            elif len(data) != 0:
                print("Wrong message!")
    finally:
        connection.close()

client.close()