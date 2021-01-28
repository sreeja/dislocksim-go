package main

import (
	// "github.com/sreeja/etcd-exp/rwlock"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	// "time"
	// "net"
	"errors"
	"net/http"
)

var replicas map[string]int
var whoami string

var cli *clientv3.Client
var session *concurrency.Session

func main() {
	replicas = map[string]int{"houston": 1, "paris": 2, "singapore": 3}
	whoami = os.Getenv("WHOAMI")

	//setting up log file
	filename := "/usr/data/log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	// create etcd client
	log.Println("CREATE CLIENT")
	cli, err = clientv3.New(clientv3.Config{Endpoints: []string{"etcd-" + strconv.Itoa(replicas[whoami]) + ":2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create etcd client session
	log.Println("CREATE SESSION")
	//send TTL updates to server each 1s. If failed to send (client is down or without communications), lock will be released
	session, err = concurrency.NewSession(cli, concurrency.WithTTL(1))
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	handleRequests()

}

func handleRequests() {
	http.HandleFunc("/do", do)
	log.Fatal(http.ListenAndServe(":6000", nil))
}

func do(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	// args := r.URL.Query()
	// app := args["app"][0]
	// op := args["op"][0]
	// timetosleep := getexectime(app, op)
	// locks := getlocks(app, op)
	// var reallocks map[Lock]*RWMutex
	// for l := range locks {
	// 	//acquire each lock
	// 	l1 := rwlock.NewRWMutex(session, l.Name)
	// 	reallocks = append(reallocks, l1)
	// 	reallocks
	// 	if l.Mode == "shared" {
	// 		rlerr := l1.RLock()
	// 		defer l1.RUnlock()
	// 		check(rlerr)
	// 	} else {
	// 		wlerr := l1.Lock()
	// 		defer l1.Unlock()
	// 		check(wlerr)
	// 	}
	// }
	// time.Sleep(time.Duration(timetosleep) * time.Millisecond)
	// for l := range reallocks {
	// 	//release each lock

	// }
}

func getexectime(appname, opname string) (int, error) {
	configfile := "./config/application/" + appname + ".json"
	f, err := ioutil.ReadFile(configfile)
	if err != nil {
		return -1, err
	}

	var exectimes []Exectime
	json.Unmarshal([]byte(string(f)), &exectimes)

	for e := range exectimes {
		if exectimes[e].Name == opname {
			return exectimes[e].Time, nil
		}
	}
	return -1, errors.New("Operation not found")
}

func getlocks(appname, opname, oplock, locktype string) ([]Lock, error) {
	oplockfile := "./config/locker/" + appname + "/oplock" + oplock + ".json"
	oplockf, err := ioutil.ReadFile(oplockfile)
	if err != nil {
		return nil, err
	}
	var oplocks []OpLock
	json.Unmarshal([]byte(string(oplockf)), &oplocks)

	locktypefile := "./config/locker/" + appname + "/locktype" + locktype + ".json"
	locktypef, err := ioutil.ReadFile(locktypefile)
	if err != nil {
		return nil, err
	}
	var locktypes []LockType
	json.Unmarshal([]byte(string(locktypef)), &locktypes)

	var locks []Lock
	for o := range oplocks {
		if oplocks[o].Op == opname {
			for l := range oplocks[o].Locks {
				for t := range locktypes {
					if locktypes[t].Name == oplocks[o].Locks[l].Name {
						lock := Lock{oplocks[o].Locks[l].Name, oplocks[o].Locks[l].Mode, locktypes[t]}
						locks = append(locks, lock)
					}
				}
			}
		}
	}
	return locks, nil
}
