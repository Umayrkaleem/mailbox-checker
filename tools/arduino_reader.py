#!/usr/bin/env python3

import serial
import re
import sys
from time import sleep

USBPATH = '/dev/cu.usbserial-0001'

arduino = serial.Serial(USBPATH, 9600)
rawData = arduino.readline()
data = rawData.decode()

print("Raw Data = ", rawData, "\n", "data = ", data)
