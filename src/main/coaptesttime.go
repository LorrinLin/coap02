package main

import (
	coap "github.com/lorrin/go-coap"
	"strconv"
	"os"
	"log"
	"time"
)

var(
	i int
	t = 10
)

// This is a coap client used to send request messages to server and then receive response messages,
//what is more, it will calcuate the time duration it used
// PS: it must use server.go in coap-go-first as a server 
func main(){
	
	start := time.Now()
	c, err := coap.Dial("udp", "localhost:5683")
	if err!= nil{
		log.Println("err in dial..",err)
	}
	
	for i=0; i<t; i++{
	
		req := coap.Message{
			Type:		coap.Confirmable,
			Code:		coap.GET,
			MessageID:	uint16(i),
			Payload:	[]byte(strconv.Itoa(i)),
		}
	
		path := "my/test"
		if len(os.Args) >1 {
			path = os.Args[1]
		}
		req.SetPathString(path)
		rv, err := c.Send(req)
		if err == nil{
			if rv != nil{
				if err!= nil{
					log.Println("err in send..",err)
				}
				payload := string(rv.Payload)
				log.Println("Got response message payload:",payload,i)
			}
			rv, err = c.Receive()
		}
		//log.Println("------",i)
	}
	
	duration := time.Since(start)
	log.Println("duration time :",duration)
	log.Println("average duration time :",duration/10)
	log.Println("------done------")
}