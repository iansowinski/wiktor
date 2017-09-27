#!/usr/bin/env python
# -*- coding: utf-8 -*-

from InstagramAPI import InstagramAPI
import socket
import getpass

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('localhost', 9999))
server.listen(1)

login = raw_input("Instagram login: ")
password = getpass.getpass("Instagram password: ")

print "Logged in!"

api = InstagramAPI.InstagramAPI(login, password)
api.login()

print "You are now free to start server.py by running `python3 server.py`"

while True:
    connection, client_address = server.accept()
    try:
        while True:
            filename = connection.recv(1024).decode('utf-8')
            photo = "images/" + filename + ".jpg"
            if len(filename) != 0:
                api.uploadPhoto(photo,caption="#PoliceBrutality",upload_id=None)
                print "Uploaded!"
    finally:
        connection.close()
