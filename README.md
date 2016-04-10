# EECS345
Distributed Systems - Spring 2016 - Northwestern University
## Project 0: NU Chitter
### SUMMARY
This is the first project for the course, which we have called "NU Chitter". It should serve as both an introduction to Go as well as a self-evaluation of proficiency in networking and systems programming. This project must be completed individually. All submissions should be 100% your own work and will be graded on a Pass/Fail scale -- either the submitted code works or it does not.

NU Chitter is a very simple chat server with support for private messaging. By default, incoming messages should be broadcast to all connected clients. Incoming messages that contain a colon are considered "commands". There are a total of three commands "whoami:", "all:", and private messages. More details on these commands are provided below.

Note: Your code *must* be thread safe! You should not be sharing objects across goroutines. Code that is not thread safe will be very heavily penalized (50% off). Furthermore, you *must* use channels to communicate between goroutines! Submissions that use locking mechanisms will also be very heavily penalized.
## Project 1: Kademlia DHT Part 1
## Project 2: Kademlia DHT Part 2
## Project 3: Vanish
