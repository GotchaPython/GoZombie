package main

import (
    "bufio"
    "github.com/zhouhui8915/go-socket.io-client"
    "log"
    "os"
    "strings"
    "io/ioutil"
    "net/http"
)

func main() {

    basicAuth()
    opts := &socketio_client.Options{
        Transport: "websocket",
        Query:     make(map[string]string),
    }
    opts.Query["user"] = "user"
    opts.Query["pwd"] = "pass"
    uri := "http://138.197.30.113:80/socket.io/"

    client, err := socketio_client.NewClient(uri, opts)
    if err != nil {
        log.Printf("NewClient error:%v\n", err)
        return
    }

    client.On("error", func() {
        log.Printf("on error\n")
    })
    client.On("connection", func() {
        log.Printf("on connect\n")

    })
    client.On("chat message", func(msg string) {
        log.Printf("on message:%v\n", msg)
    })
    client.On("disconnection", func() {
        log.Printf("on disconnect\n")
    })

    reader := bufio.NewReader(os.Stdin)
    for {
        data, _, _ := reader.ReadLine()
        command := string(data)
        if strings.Contains(command,"report"){
        	client.Emit("chat message", "Reporting in!")
        	
        	
        }
        client.Emit("chat message", command)
        log.Printf("send message:%v\n", command)
    }
}

func basicAuth() string {
    var username string = "trj2mch"
    var passwd string = "Security55%25%25"
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://google.com", nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    return s
}