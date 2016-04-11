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

  ``` go
  go run chitter.go 12345
  ```

* By default, messages should be broadcasted to all clients. In other words, any messages that does not contain a colon should be broadcast. Furthermore, incoming messages that begin with "all:" should also be broadcast (more on commands below). For example, if a client sends any of the following messages (one per line), the body of the message should be forwarded to all clients. In the examples below, the body of each message is in bold text ("Hello, world!"). 

  ```
  Hello, world!
  all: Hello, world!
  all : Hello, world!
  ```

* When messages are forwarded to other clients, they should be prefixed with the ID of the user that sent the message. In our example above, the "Hello, world!" messages from a user with an ID of 42 would appear as follows:

 ```
 42: Hello, world!
 ```
 
  * Commands: Your NU Chitter program will need to support three types of "commands" that are parsed out by the server. This includes the "all:" command, the "whoami:" command, and sending PMs. If a user sends an invalid command, you can ignore the message.
  * all: - The "all:" command allows a user to explicitly broadcast a message. This also allows a user to use a colon ":" in their message. Remember that by default, if NU Chitter receives a message with one or more colons in it, everything before the first colon is interpretted as a command. Using the "all:" command allows users to have colons in their messages.
  * whoami: - Users should be able to get their own ID number by using the whoami: command. Below is an example of user 13 asking what their ID is with the NU Chitter server responding only to that user.
  
    ```
    whoami:
    chitter: 13
    ```
  
  Note that anything after a "whoami:" command can simply be ignored by the server.
  
  * Private messages - Users should also be able to send a message to a specific user. To do so, their message should be prefixed with the ID of the user they wish to message. Private messages to nonexistent users can just be ignored (i.e., the server does nothing). If user 1 wants to message user 2, their submitted message should be in the following format:

    ```
    2: Hey, number two!
    ```
  
  And should appear as the following on user 2's (and only user 2's) connection:
  
    ```
    1: Hey, number two!
    ```
  
## Project 1: Kademlia DHT Part 1
## Project 2: Kademlia DHT Part 2
## Project 3: Vanish
