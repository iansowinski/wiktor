import socket
import os
from datetime import datetime

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('localhost', 10001))
server.listen(1)

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('localhost', 9999))

while True:
    connection, client_address = server.accept()
    try:
        while True:
            data = connection.recv(1024).decode('utf-8')
            if data == "True":
                print(data)
                timestamp = datetime.now().strftime("%Y-%m-%d-%H:%M:%S")
                os.system("cd images && imagesnap -w 1 " + timestamp + ".jpg > /dev/null")
                client.sendall(timestamp.encode('utf-8'))
            elif len(data) != 0:
                print(data)
    finally:
        connection.close()

client.close()