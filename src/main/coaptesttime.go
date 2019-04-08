package main

import (
	coap "github.com/lorrin/go-coap"
	"strconv"
	"os"
	"log"
	"time"
	"sync"
)

var(
	i int
	t = 1000
	wg sync.WaitGroup
)

// This is a coap client used to send 100 request messages to server and then receive response messages,
// what is more, it will calcuate the time duration it used
// PS: it must use server.go in coap-go-first as a server 
func main(){
	
//	start := time.Now()
	// 142.93.161.16 is my own remote server,
	// which is running as a coap server in port NO:5683 
	c, err := coap.Dial("udp", "142.93.161.16:5683")
	if err!= nil{
		log.Println("err in dial..",err)
	}
	start := time.Now()
	for i=0; i<t; i++{
		wg.Add(1)
		go sendRequestMesage(c, i)
	}
	wg.Wait()
	duration := time.Since(start)
	log.Println("----- 1 client send 1000 message cost time :",duration)
	log.Println("----- average duration time :",duration/1000)
	
}

func sendRequestMesage(c *coap.Conn,i int){

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
				wg.Done()
			}
			rv, err = c.Receive()
		}
		//log.Println("------",i)
		
}