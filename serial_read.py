from time import sleep
import serial
ser = serial.Serial('/dev/cu.wchusbserial1410', 9600) # Establish the connection on a specific port
# counter = 32 # Below 32 everything in ASCII is gibberish
while True:
    # counter +=1
    print(str(ser.readline())) # Read the newest output from the Arduino
    sleep(.1) # Delay for one tenth of a second
    # if counter == 255:
    #     counter = 32

