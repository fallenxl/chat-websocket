// client.go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(p))
			fmt.Print("> ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Println(err)
			return
		}
	}
}
