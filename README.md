# EECS345
Distributed Systems - Spring 2016 - Northwestern University
## Project 0: NU Chitter
### SUMMARY
This is the first project for the course, which we have called "NU Chitter". It should serve as both an introduction to Go as well as a self-evaluation of proficiency in networking and systems programming. This project must be completed individually. All submissions should be 100% your own work and will be graded on a Pass/Fail scale -- either the submitted code works or it does not.

NU Chitter is a very simple chat server with support for private messaging. By default, incoming messages should be broadcast to all connected clients. Incoming messages that contain a colon are considered "commands". There are a total of three commands "whoami:", "all:", and private messages. More details on these commands are provided below.

Note: Your code *must* be thread safe! You should not be sharing objects across goroutines. Code that is not thread safe will be very heavily penalized (50% off). Furthermore, you *must* use channels to communicate between goroutines! Submissions that use locking mechanisms will also be very heavily penalized.

### TURNING IN
Simply upload your code to our submission server. The site will ask you to log in with your Northwestern Net ID and password. If the page has trouble loading, please hit try again (either refresh or follow the link again).

For this project, you should be able to fit all your code within a single file -- the reference implementation is relatively short (less than 150 lines of code with comments). Projects 1-3 will require more than a single file. For those, we will provide instructions on how to pack them into a submittable tar.gz file. However, for this project, just submit your single source code file named "chitter.go".

### PROJECT DESCRIPTION
The specification of this project is relatively straight forward. Remember, it's basically a chatroom that has support for private messages (PMs). For this project, we only need to implement the NU Chitter server. You can use a program such as nc or telnet as a client (more on this in the testing section below).

* Your NU Chitter implementation should require a port number when starting the sever. NU Chitter will listen for connections on this port. Below is an example of starting NU Chitter in the command line.
```
go run chitter.go 12345

## Project 1: Kademlia DHT Part 1
## Project 2: Kademlia DHT Part 2
## Project 3: Vanish
