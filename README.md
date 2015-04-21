# _*chat350*_
A realtime chat application built using AngularJS and Go

## Technologies used
* AngularJS
* Go
* WebSockets API

## Overview
The server written in Golang listens and serves back the angular application. It
also listens for incoming websocket connections and stores them in a hash map. 
Then, the websocket server simply broadcasts incoming message data to all
connections in the map, it also logs them to the console for inspection. In this
manner, the websocket server is functioning as a chat server. When regular http
requests come in, the server will just serve back the files in the www/ directory.
If a request comes in asking for the `/ws` resource, the websocket handler is
triggered. 

The angular application generates a JSON message to be sent to the chat server
when the user submits a message. It also unpacks and displays the contents of
newly received messages from the chat server. Each JSON message contains the 
following fields:
* The contents of the chat message
* The type of message (ex. text)
* The user name of the author
* A timestamp for the message

## AngularJS
This framework was used in order to make use of real time two way data binding.
ngAnimate was also used for smooth transitions. I added my own angular directive
`ngEnter` in order to to capture key-presses of the enter key in order to allow
users to simply hit enter to submit their message, as opposed to having to always
click the send button. This directive could be included as an attribute to other
input boxes.

## Go
Go is a fun way to write server side code. It reduces the amount of boilerplate
that you have to write, but is also extensible and scalable. I chose to use Go
in order to get more experience using it to write applications. This application
requires the installation of the gorilla websockets library. Some elements from
the Std library were used such as net/http, log and flag.

## WebSockets API
The standard HTTP protocol provides a stateless half-duplex connection. The 
WebSockets API on the other hand creates a full duplex connection between the 
server and the client which can make it more responsive than traditional
HTTP in real time applications and generates less data traffic.

## Bootstrap
Twitter Bootstrap was used for CSS thememing and UI layout. The modal constructs
were used which requires jQuery as well. The Google fonts API was used for the 
title on the jumbotron. 

## Further Ideas
Here are a list of further ideas that could be added to the design project:
* Make the Golang server operate concurrently
* Proper user authentication scheme
* Support for posting images
* Admin account with ability to boot certain users
* Whisper mode to send private messages to a single recipient
* Display list of currently logged in users

## Final Comments
This project gave me the opportunity to use git and github for revision control.
This allowed me to significantly increase my skills at using git and some of its
basic features. I also learned some of the basics of how to use Markdown. The 
technologies I used were not covered much in class which required a large amount 
of extra reading of documentation, tutorials and instructional videos.

My primary goal with this project was to gain further skills that I could use
to complement my electrical engineering background. I plan to specialize in 
embedded systems development and the ability to use websockets and angularJS
will help me in developing powerful, user friendly interfaces for my applicaitons.
Golang as well is a very interesting way to write systems software that can
easily be scaled for distributed systems for maximum concurrency. The use
of interfaces to maximize abstraction and productivity is also very valuable.

## Sources
_calicoJake_ provided some useful starting material which I was inspired by so
I continued to develop upon his basic starting application.
https://www.youtube.com/watch?v=U80k7fTEqNw
