#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
# Use text editor to edit the script and type in valid Instagram username/password

from InstagramAPI import InstagramAPI
import socket
import os
import os.path
from datetime import datetime
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
            data = connection.recv(1024).decode('utf-8')
            photo = "images/" + data + ".jpg"
            if len(data) != 0:
                api.uploadPhoto(photo,caption="test",upload_id=None)
                print "Uploaded!"
    finally:
        connection.close()
