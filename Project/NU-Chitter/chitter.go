// Author: Yuanhui Yang                  
// NetId: yyd198
// Email: yuanhui.yang@u.northwestern.edu

package main

import (
    "net"
    "bufio"
    "strconv"
    "fmt"
    "strings"
    "os"
)

type Client struct {
    id int
    connection net.Conn
    incoming chan string
    outgoing chan string
}

type ChatRoomInComing struct {
    clientId int
    clientOutComing string
}

type ChatRoom struct {
    clients []*Client
    joinConn chan net.Conn
    incoming chan ChatRoomInComing
}

type Message struct {
    senderId int
    receiverId int
    text string
}

func NewMessage(senderId_ int, receiverId_ int, text_ string) *Message {
    message := &Message {
        senderId: senderId_,
        receiverId: receiverId_,
        text: text_,
    }
    return message
}


func (client* Client) Read() {
    buf := bufio.NewReader(client.connection)
    for {
        line, err := buf.ReadString('\n')
        if err != nil {
            break
        } else {
            client.outgoing <- line
        }
    }
}

func (client* Client) Write() {
    for message := range client.incoming {
        client.connection.Write([]byte(message))
    }
}

func (client* Client) PrintId() {
    id_ := "chitter: " + strconv.Itoa(client.id) + "\n"
    client.incoming <- id_
}

func (client* Client) Listen() {
    go client.Read()
    go client.Write()
}

func NewClient(conn net.Conn, id_ int) *Client {
    client := &Client {
        id: id_,
        connection: conn,
        incoming: make(chan string),
        outgoing: make(chan string),
    }
    client.Listen()
    return client
}

func (chatRoom* ChatRoom) sendPersonalMessage(message *Message) {
    if message.receiverId <= len(chatRoom.clients) {
        if message.receiverId >= 1 {
            if message.senderId <= len(chatRoom.clients) {
                if message.senderId >= 1 {
                    text := strconv.Itoa(message.senderId) + ": " + message.text
                    chatRoom.clients[message.receiverId - 1].incoming <- text
                }
            }
        }
    }
}

func (chatRoom* ChatRoom) Broadcast(message *Message) {
    if (message.receiverId == -2) {
        for _, client := range chatRoom.clients {
            text := strconv.Itoa(message.senderId) + ": " + message.text
            client.incoming <- text
        }
    }
}

func (chatRoom* ChatRoom) Add(conn net.Conn) {
    client := NewClient(conn, len(chatRoom.clients) + 1)
    chatRoom.clients = append(chatRoom.clients, client)
    go func() {
        for {
         chatRoomInComing := ChatRoomInComing{
             clientId: client.id,
             clientOutComing: <-client.outgoing,
         }
            chatRoom.incoming <- chatRoomInComing
        }
    }()
}

func (chatRoom* ChatRoom) Listen() {
    go func() {
        for {
            select {
                case augmentedSay := <-chatRoom.incoming: {
                 senderId := augmentedSay.clientId
                 say := augmentedSay.clientOutComing
                 if len(say) >= 6 && string(say[0:6]) == "whoami" {
                     senderClient := chatRoom.clients[senderId - 1]
                     senderClient.PrintId()
                 } else if chatRoom.isPersonalMessage(say) {
                     message := NewMessage(senderId, -1, "")
                     chatRoom.handlePersonalMessage(say, message)
                     chatRoom.sendPersonalMessage(message)
                 } else {
                     message := NewMessage(senderId, -1, "")
                     chatRoom.handleBroadcastMessage(say, message)
                     chatRoom.Broadcast(message)                     
                 }
                }
                case conn := <-chatRoom.joinConn: {
                    fmt.Println("New Client ID:", len(chatRoom.clients) + 1)
                    chatRoom.Add(conn)
                }
            }
        }
    }()
}

func (chatRoom* ChatRoom) isPersonalMessage(say string) bool {
    result := false
    index := 0
    for index = 0; index < len(say); index++ {
        if string(say[index]) == ":" {
            break
        }
    }
    if index < len(say) {
        TrimSay := strings.Trim(say[0:index], " ")
        id, err := strconv.Atoi(TrimSay)
        if err == nil {
            if id >= 1 {
                if id <= len(chatRoom.clients) {
                    result = true
                }
            }
        }
    }
    return result
}

func (chatRoom* ChatRoom) handlePersonalMessage(say string, message *Message) {
    index := 0
    for index = 0; index < len(say); index++ {
        if string(say[index]) == ":" {
            break
        }
    }
    if index < len(say) {
        SayTrim := strings.Trim(string(say[0:index]), " ")
        receiverId, err := strconv.Atoi(SayTrim)
        if err == nil {
            message.receiverId = receiverId
            if index + 1 < len(say) {
                message.text = string(say[(index + 1) :])
                message.text = strings.Trim(message.text, " ")
            } else {
             message.text = ""
         }
        }
    }
}

func (chatRoom* ChatRoom) handleBroadcastMessage(say string, message *Message) {
    index := 0
    for index = 0; index < len(say); index++ {
        if string(say[index]) == ":" {
            break
        }
    }
    message.text = say
    if index < len(say) {
        sayTrim := strings.Trim(string(say[0:index]), " ")
        if sayTrim == "all" {
            if index + 1 < len(say) {
                message.text = say[(index + 1):]
            } else {
                message.text = ""
            }
        }
    }
    message.text = strings.Trim(message.text, " ")
    message.receiverId = -2
}

func NewChatRoom() *ChatRoom {
    chatRoom := &ChatRoom{
        clients: make([]*Client, 0),
        joinConn: make(chan net.Conn),
        incoming: make(chan ChatRoomInComing),
    }
    chatRoom.Listen()
    return chatRoom
}

func (chatRoom* ChatRoom) buildConn(listener net.Listener) {
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Listening and Unable to Accept")
            continue
        } else {
            chatRoom.joinConn <- conn
        }
    }
}

func main() {
    chatRoom := NewChatRoom()
    if len(os.Args) < 2 {
        os.Exit(1)
        return
    } else {
        listener, err := net.Listen("tcp", os.Args[len(os.Args) - 2] + ":" + os.Args[len(os.Args) - 1])
        if err != nil {
            fmt.Println("\nUnable to Listen\n")
        } else {
            fmt.Println("\nListening\n")
            chatRoom.buildConn(listener)
        }    	
    }
    fmt.Println("\nYuanhui Yang\n")
}
