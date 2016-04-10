# EECS345
Distributed Systems - Spring 2016 - Northwestern University
## Project 0: NU Chitter
<h3>Summary</h3>
  <p>This is the first project for the course, which we have called "NU Chitter". It should serve as both an introduction to Go as well as a self-evaluation of proficiency in networking and systems programming. This project must be completed individually. All submissions should be 100% your own work and will be graded on a Pass/Fail scale -- either the submitted code works or it does not.</p>
  
    <p>NU Chitter is a very simple chat server with support for private messaging. By default, incoming messages should be broadcast to all connected clients. Incoming messages that contain a colon are considered "commands". There are a total of three commands "whoami:", "all:", and private messages. More details on these commands are provided below.</p>
  
    <p><i><b>Note:</b> Your code *must* be thread safe! You should not be sharing objects across goroutines. Code that is not thread safe will be very heavily penalized (50% off). Furthermore, you *must* use channels to communicate between goroutines! Submissions that use locking mechanisms will also be very heavily penalized.</i></p>
  



  
<h3>Turning in</h3>
    
  <p>Simply upload your code to our <a href="https://www.cs.northwestern.edu/~aqualab/assignments/DS/submission.htm">submission server</a>. The site will ask you to log in with your Northwestern Net ID and password. If the page has trouble loading, please hit try again (either refresh or follow the link again).</p>
  <p>For this project, you should be able to fit all your code within a single file -- the reference implementation is relatively short (less than 150 lines of code with comments). Projects 1-3 will require more than a single file. For those, we will provide instructions on how to pack them into a submittable tar.gz file. However, for this project, just submit your single source code file named "chitter.go".</p>  
  
  
<h3>Project description</h3>
  <p>The specification of this project is relatively straight forward. Remember, it's basically a chatroom that has support for private messages (PMs). For this project, we only need to implement the NU Chitter server. You can use a program such as nc or telnet as a client (more on this in the testing section below).</p>
  <ul>
    <li>Your NU Chitter implementation should require a port number when starting the sever. NU Chitter will listen for connections on this port. Below is an example of starting NU Chitter in the command line.<pre>go run chitter.go 12345</pre></li>
    <li>By default, messages should be broadcasted to all clients. In other words, any messages that does not contain a colon should be broadcast. Furthermore, incoming messages that begin with "all:" should also be broadcast (more on commands below). For example, if a client sends any of the following messages (one per line), the body of the message should be forwarded to all clients. In the examples below, the body of each message is in bold text ("Hello, world!").<pre><b>Hello, world!</b><br/>all: <b>Hello, world!</b><br/>all : <b>Hello, world!</b></pre></li>
    
    <li>When messages are forwarded to other clients, they should be prefixed with the ID of the user that sent the message. In our example above, the "Hello, world!" messages from a user with an ID of 42 would appear as follows:<pre>42: Hello, world!</pre></li>
    
    <li><b>Commands:</b> Your NU Chitter program will need to support three types of "commands" that are parsed out by the server. This includes the "all:" command, the "whoami:" command, and sending PMs. If a user sends an invalid command, you can ignore the message.</li>
      <ul>
    
    <li><b>all:</b> - The "all:" command allows a user to explicitly broadcast a message. This also allows a user to use a colon ":" in their message. Remember that by default, if NU Chitter receives a message with one or more colons in it, everything before the first colon is interpretted as a command. Using the "all:" command allows users to have colons in their messages.</li>
<li><b>whoami:</b> - Users should be able to get their own ID number by using the whoami: command. Below is an example of user 13 asking what their ID is with the NU Chitter server responding only to that user.<pre>whoami:<br/>chitter: 13</pre> Note that anything after a "whoami:" command can simply be ignored by the server.</li>
          
      
      <li><b>Private messages</b> - Users should also be able to send a message to a specific user. To do so, their message should be prefixed with the ID of the user they wish to message. Private messages to nonexistent users can just be ignored (i.e., the server does nothing). If user 1 wants to message user 2, their submitted message should be in the following format:<pre>2: Hey, number two!</pre>And should appear as the following on user 2's (and only user 2's) connection:<pre>1: Hey, number two!</pre></li></ul>
    
  </ul>

      
<h3>Debugging &amp; testing</h3>  
  <p>To test your own NU Chitter implementation, you will want to use either netcat (nc) or telnet. Both should be able to connect to your server with your hostname and port.</p>
  <p>Additionally, a reference solution will be running on aqualab.cs.northwestern.edu on port 10345. Note that this server is only reachable from Northwestern's network. If you are trying to access it from another network, you will need to connect to <a href="http://www.it.northwestern.edu/oncampus/vpn/">Northwestern's VPN service.</a> In order to help simplify testing, we have also implemented a "who" command on our test server. Similar to it's Unix counterpart, "who" returns a list of all the users connected to the server. You do not need to implement this function (though it should be trivial).</p>
  <p>Your program also <b>MUST</b> be thread safe. We will be stress testing it to make sure that there are not any incorrect behaviors. You can check for race conditions by using the following command when you test your server:<pre>go run -race chitter.go 12345</pre>
<h3>Suggested Plan of attack</h3>
  <p> Below is an outline of our suggest plan of attack. If you are already comfortable with Go, the following can still serve as a useful guide for the suggested order in which to implement each feature. </p>
  <ul>
    <li>Write a simple echo server in Go. Make sure the server is able to handle concurrent connections (hint: use Goroutines).</li>
    <li>Modify the server so that incoming messages are forwarded to all connected clients. The server will need to maintain a record of clients that are currently connected to the server. You will need to begin using Go's channels and the select statement to do so. </li>
    <li>Assign unique IDs to the connected clients. Our reference implementation assigns integers (0, 1, 2, etc.), but you can assign IDs however you prefer -- just do not assign the IDs "all" and "whoami", those are reserved for commands. Once users have IDs, prefix outgoing messages with the ID of the user that sent the message.</li>
    <li>Add support for the "whoami" and "all" commands. You can assume any line of text with a colon contains a command, and that all text before the colon is the command. If a user sends an invalid command, you can ignore the message. If there is no colon, it is a broadcast message. Go's <a href="https://godoc.org/strings">"strings"</a> package should make parsing messages from users straightforward.</li>
    <li>Add support for PMs. Clients should be able to send a PM by prefacing a message with the desired user's ID and a colon (e.g., 1: &lt;message&gt;). </li>
  </ul>
    
</div>
## Project 1: Kademlia DHT Part 1
## Project 2: Kademlia DHT Part 2
## Project 3: Vanish
