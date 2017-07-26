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

api = InstagramAPI.InstagramAPI(login, password)
api.login()

while True:
    connection, client_address = server.accept()
    try:
        while True:
            data = connection.recv(1024).decode('utf-8')
            photo = "images/" + data + ".jpg"
            if len(data) != 0:
                print(data)
                api.uploadPhoto(photo,caption="test",upload_id=None)
    finally:
        connection.close()
