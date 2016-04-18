# SUMMARY
The first two group projects (Project 1 and Project 2) focus on building the Kademlia DHT -- a very popular DHT first published in 2002. Since this is a rather large task, it is split into two projects. In Project 1 you will start from a minimal base and, following the specification linked below, implement the routing underlay of Kademlia. Do not let the phrase "minimal base" fool you -- this project is a lot of work and you will need to familiarize yourself with a quite a bit of code, so be sure to start early!

Few distributed systems operate in homogeneous environments. Your project will be evaluated based on its ability to interoperate with other Kademlia implementations: if everyone implements the requirements correctly, we will be able to run a single DHT that includes clients from every group. We plan to provide access to a reference implementation later in the quarter for you to check your code against to make sure it works.

To ensure compatibility, you must strictly adhere to both the the Xlattice project's [kademlia spec](http://xlattice.sourceforge.net/components/protocol/kademlia/specs.html) and the README in the project tarball. In addition, project description below should help get you started. A link to the Kademlia paper is provided in the "Resources" section below; however, you do not need to implement all the features described there. Stick to the [XLattice spec](http://xlattice.sourceforge.net/components/protocol/kademlia/specs.html)

# IMPORTANT DATES
Out: April 15, 2016

In: April 25, 2016 (11:59PM CST)

# SKELETON
[eecs345-kademlia.tar.gz]()

# TURNING IN
Before submitting your code, please be sure to add your Net IDs to the "netIds" variable in src/kademlia/kademlia.go. In fact, Kademlia won't run if you don't.

To turn in your project, simply make a gzipped tarball of your code. You should only need to submit your 'src' directory. If you do not know how to make a tarball, there is a make_handin.sh script that will make one for you. Upload the .tgz/.tar.gz file to the submission server. Do not upload any other format (e.g., bzip, xz, zip, etc). Please try to avoid uploading any extraneous files.

Note: Some students were confused about the message on the submission server to only upload ".txt" files. This only applies to homework assignments, not projects. Upload the gzipped tarball without modifying the file extension.

# PROJECT DESCRIPTION
## Getting started
After downloading the skeleton code below, unpack it using your favorite method. From the command line, you can run the following command to unpack the contents.

```
tar vfx eecs345-kademlia.tar.gz
```

Go's build tools depend on the value of the GOPATH environment variable. $GOPATH should be the project root -- the absolute path of the directory containing the src (and bin and pkg) folders. After you cd into the project folder (kademlia), you can set GOPATH by running the following command

```
export GOPATH=`pwd`
```

Once you've set that, you should be able to build the skeleton and create an executable at bin/kademlia with:

```
go install kademlia
```

Running main as:

```
./bin/main localhost:7890 localhost:7890
```

will cause it to start up a server bound to localhost:7890 (the first argument) and then ping the second contact. All it does by default is perform an (incomplete) PING RPC and wait for more commands. You should eventually replace this code with a call to the DoPing function, once it is completed.

## Skeleton code
The skeleton code for Project 1 is split into two directories. The "kademlia" directory contains the interpreter front end. "libkademlia" is the actual "heart" of Kademlia, this is where you'll be writing most of your code. For now, you can ignore the directory named "sss", this will be for Project 3.

Inside the "libkademlia" directory, there are three files that you need to familiarize youself with: libkademila.go, id.go, and rpcs.go. You can ignore the other files for now -- proj1_test.go is for testing and extra credit (more on this below) and vanish.go will be used in Project 3. The core functions within Kademlia such as the k-buckets, update, and protocol functions should go in the libkademlia.go file. RPC responses which are invoked by remote machines are placed in the rpcs.go file.

With that said, you may find it helpful to create more files to help organize your code. This is perfectly acceptable, however, try to do so in a logical way that keeps your code readable (e.g., kbuckets.go for the KBuckets structure, or splitting large RPC commands into separate files).

## Requirements
This section outlines the requirements for the project. We suggest that you also attack these items in the following order.

* Implement the routing tables with K-Buckets.
K-buckets are sorted by time of most recent contact (most recent at the tail of the k-bucket). Whenever your Kademlia successfully communicates with any node in the DHT, it updates (calling an Update() function) the node in the corresponding k-bucket for the following conditions:

  * If the contact exists, move the contact to the end of the k-bucket.
  * If the contact does not exist and the k-bucket is not full: create a new contact for the node and place at the tail of the k-bucket.
  * If the contact does not exist and the k-bucket is full: ping the least recently contacted node (at the head of the k-bucket), if that contact fails to respond, drop it and append new contact to tail, otherwise ignore the new contact and update least recently seen contact.

Once you have finished your K-Buckets structure, you should complete the function listed below. It should look through your K-Bucket structure to find and return the Contact with the provided nodeID or an error if the Contact is not found.

```
FindContact(nodeId ID) (*Contact, error)
```

NOTE: You will be modifying your KBuckets from different goroutines. Your KBuckets design must be designed in a threadsafe way. Again, you MUST use channels for this. Submissions that use mutexes will be penalized. The reason your code needs to be thread safe is because the RPC functions (discussed below) are running in a different goroutine and will be asynchronously reading from/updating KBuckets. In both cases, you need to be controlling access to your kbuckets.

* Implement basic RPCs for Ping(), Store(), FindNode() and FindValue()

If you have not done so already, please read about [RPCs and GO here](http://jan.newmarch.name/go/rpc/chapter-rpc.html).

As a refresher, all RPCs in GO are an interface with three restrictions. The function must be public (begin with a capital letter), have exactly two arguments, the first is a pointer to value data to be received by the function from the client, and the second is a pointer to hold the answers to be returned to the client, have a return value of type error.

In order to use RPC methods, you have to register a struct and accompanying methods with the global RPC module, then launch a server to handle RPC requests. Luckily this is already handled for you in the skeleton code.

The RPCs that you will need in order to implement the Kademlia protocol are the four given above. These RPCs are invoked remotely from other Kademlia nodes (hence the reason they are named Remote procedure calls). Be sure to note that this function returns an error when something goes wrong, but the results are actually "returned" to the client by storing them in the second argument to the function. This second object is passed by reference and changes on the server will appear on the client once the call has finished.

In addition to the RPC handlers in rpcs.go, you will need client versions of these methods which invoke RPCs in remote nodes and implement the logic of the Kademlia protocol. These are located in libkademlia/libkademlia.go. A short example of invoking an RPC call is in kademlia/kademlia.go. The four functions you will need to implement in libkademlia/libkademlia.go are below:

```
DoPing(remoteHost net.IP, port uint16)
```

The Ping is the primitive communication for node organization and update. A node sends out a ping request to other nodes. When a node receives a PingMessage, it responds with a PongMessage. Ping and Pong structs are provided in the skeleton in rpcs.go.

```
DoStore(contact *Contact, Key ID, Value []byte)
```

Store() is the primitive for storing a key-value pair remotely in the DHT. The Store() primitive takes three arguments, a Contact who should store the value, an ID for the key, and a byte array for the value.

```
DoFindValue(contact *Contact, Key ID)
```

The FindValue RPC includes a B=160-bit key. If a corresponding value is present on the recipient, the associated data is returned. Otherwise the RPC is equivalent to a FIND_NODE and a set of k triples is returned.

```
DoFindNode(contact *Contact, searchKey ID)
```

The FIND_NODE RPC includes a 160-bit key. The recipient of the RPC returns up to k triples (IP address, port, nodeID) for the contacts that it knows to be closest to the key. The recipient must return k triples if at all possible. It may only return fewer than k if it is returning all of the contacts that it has knowledge of.

The issuer of the DoFindNode RPC should then add the returned contacts to its kbuckets.

There are three DoIterative* functions provided in the skeleton. These are for Project 2 and do not need to be completed by the Project 1 deadline.

# TESTING AND DEBUGGING
As mentioned above, the Kademlia instance can be started with the following command:

```
./bin/kademlia localhost:7890 localhost:7890
```

Once you have started a Kademlia instance, it will do the intitial ping and provide you with a prompt for testing your client. A switch statement in kademlia/kademlia.go already checks for these commands. You can use these commands to do simple testing of your Kademlia code.

After starting your first client, you can start a second client with the following commonad:

```
./bin/kademlia localhost:7891 localhost:7890
```

This will start a second client on port 7891 and will ping the other Kademlia instance running on port 7890.

## Command prompt
