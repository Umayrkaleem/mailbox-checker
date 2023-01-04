# mailbox-checker

mailbox-checker will check your mail. You may live in an apartment with slow elevators or 
just don't want to walk in bad weather. 

See [How it works](#how-it-works) for more info.

# Table of Contents
- [How it works](#how-it-works)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Contributing](#contributing)

# How it works

Board and sensor sits inside in the back of your mailbox pointed towards the opening door while connected to a network. 
When an item (mail) is placed in the mailbox the HC-SR04 sensor will not be able to read the distance from the 
opening door therefore read a much smaller distance. This change in distance is what we'll use to realize we have mail
in our mailbox.  

# Prerequisites
### Software
golang 1.19  
https://go.dev/doc/install

tinygo version 0.26.0 darwin/amd64 (using go version go1.19 and LLVM version 14.0.0)

### Hardware
Arduino Board
https://amzn.to/3VKozAS

HC-SR04 Sensor
https://amzn.to/3G7v62V

Bread board and Wiring
https://amzn.to/3Q7PKnW

# Installation
example
tinygo flash -port /dev/cu.usbmodem11301 -target=arduino api/main/board.go

# Contributing
Pull requests are def welcomed. Please open an issue first to discuss what you would like to change.
