import socket
import time

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('localhost', 10001))

while True:
    message = str(input('>>> '))
    client.sendall(message.encode('utf-8'))
    time.sleep(2)

client.close()