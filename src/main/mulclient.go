package main

import (
	coap "github.com/lorrin/go-coap"
	"log"
	"strconv"
	"os"
	"sync"
	"time"
)

var(
	wg sync.WaitGroup
	start time.Time
)

// create 100 clients, and send a message to the server, get the response back
// calculate the cost time
func main(){
	
	for i:=0;i<100;i++{
		wg.Add(1)
		go createClientAndSend(i)	
	}
	wg.Wait()
	dur := time.Since(start)
	log.Println("------ 100 clients send a message cost time:",dur)
	log.Println("------ average client:",dur/100)
}

func createClientAndSend(i int){
	// 142.39.161.16 is my own remote server,
	// which is running as a coap server in port NO:5683 
	c, err := coap.Dial("udp","142.93.161.16:5683")
	if err!= nil{
		log.Println("err in Dial..",err)
	}
	
	req := coap.Message{
		Type:		coap.Confirmable,
		Code:		coap.GET,
		MessageID:	uint16(i),
		Payload:	[]byte(strconv.Itoa(i)),
	}
	
	path := "my/test"
	
	if len(os.Args) >1{
		path = os.Args[1]
	}
	
	req.SetPathString(path)
	
	start = time.Now()
	
	rv, err := c.Send(req)
	if err == nil{
			if rv != nil{
				if err!= nil{
					log.Println("err in send..",err)
				}
				payload := string(rv.Payload)
				log.Println("Got response message payload:",payload,i)
				wg.Done()
			}
			rv, err = c.Receive()
		}
}