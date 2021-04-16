package main

import (
	"bufio"
	"fmt"
	"github.com/sreeja/etcd-exp/rwlock"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var whereami string

var session *concurrency.Session

func main() {
	replicas := map[string]int{"houston": 0, "paris": 1, "singapore": 2}
	whereami = os.Getenv("WHEREAMI")
	if _, ok := replicas[whereami]; !ok {
		log.Fatal("Replica not listed")
	}

	//setting up log file
	filename := "../data/locklog.txt"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	// create etcd client
	etcd_client_name := "etcd-" + whereami + ":2379"
	log.Println("Creating etcd client", etcd_client_name, time.Now())
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{etcd_client_name}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create etcd client session
	log.Println("Creating etcd session", time.Now())
	//send TTL updates to server each 5s. If failed to send (client is down or without communications), lock will be released
	session, err = concurrency.NewSession(cli, concurrency.WithTTL(5))
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Printf("Starting lock manager at %v", whereami)

	// listen on port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer ln.Close()

	for {
		// accept connection
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error connecting:", err.Error())
		}
		log.Println("Client " + conn.RemoteAddr().String() + " connected")

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) error {
	// run loop forever (or until ctrl-c)
	locks := map[string]*rwlock.RWMutex{}
	log.Println("Waiting for msg")
	for {
		// get message, output
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Client left", time.Now())
		}
		log.Println("Some message received")
		text := "Message Received:" + string(message)
		log.Println(text)

		parts := strings.Split(message, ";")
		if parts[0] == "acquire" {
			for i, l := range parts {
				if i > 0 { //to ignore acquire release decision as it has been taken earlier
					//this is a lock in the format "mode(X/S):name"
					components := strings.Split(l, ":")
					l1 := rwlock.NewRWMutex(session, components[1])
					locks[components[1]] = l1
					if components[0] == "X" {
						log.Println("Asked lock", components[1], time.Now())
						wlerr := l1.Lock()
						if wlerr != nil {
							log.Fatal("Error while acquiring lock", components[1])
						}
						log.Println("Got lock", time.Now())
					} else {
						log.Println("Asked read lock", components[1], time.Now())
						rlerr := l1.RLock()
						if rlerr != nil {
							log.Fatal("Error while acquiring read lock", components[1])
						}
						log.Println("Got read lock", time.Now())
					}
				}
			}
		} else { //release
			for i, l := range parts {
				if i > 0 { //to ignore acquire release decision as it has been taken earlier
					//this is a lock in the format "mode(X/S):name"
					components := strings.Split(l, ":")
					l1 := locks[components[1]]
					if components[0] == "X" {
						log.Println("Releasing lock", components[1], time.Now())
						wlerr := l1.Unlock()
						if wlerr != nil {
							log.Fatal(wlerr, time.Now())
						}
						log.Println("Released lock", time.Now())
					} else {
						log.Println("Releasing read lock", components[1], time.Now())
						rlerr := l1.RUnlock()
						if rlerr != nil {
							log.Fatal(rlerr, time.Now())
						}
						log.Println("Released read lock", time.Now())
					}
				}
			}
		}
		//send message to client
		fmt.Fprintf(conn, "done\n")
	}
}

func logexectime(app, op string, start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("started at %v", start)
	log.Printf("finished at %v", t)
	log.Printf("execution time for %v, %v :%v", app, op, elapsed)
}
