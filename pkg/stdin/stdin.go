package stdin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ConsulEvents []ConsulEvent

type ConsulEvent struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Payload []byte `json:"Payload"`
	LTime   int    `json:"LTime"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	line, err := ParseEvents()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(line)
}

func ParseEvents() (string, error) {
	//標準入力をバッファリングする
	reader := bufio.NewReader(os.Stdin)
	// 先頭の1byteを覗き見る
	b, _ := reader.Peek(1)
	if string(b) == "[" {
		var evs ConsulEvents
		dec := json.NewDecoder(reader)
		if err := dec.Decode(&evs); err != nil {
			return "", err
		}
		ev := &evs[len(evs)-1]
		return string(ev.Payload), nil
	} else {
		//JSONでなければ1行読み取る
		line, err := reader.ReadString('\n')
		return line, err
	}
}
