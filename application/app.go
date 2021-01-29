package main

import (
	"github.com/sreeja/etcd-exp/rwlock"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	// "net"
	"errors"
	"net/http"
	"sort"
)

var replicas map[string]int
var whoami string

var cli *clientv3.Client
var session *concurrency.Session

func main() {
	replicas = map[string]int{"houston": 0, "paris": 1, "singapore": 2}
	whoami = os.Getenv("WHOAMI")
	if _, ok := replicas[whoami]; !ok {
		log.Fatal("Replica not listed")
	}

	//setting up log file
	filename := "../data/log.txt"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	// create etcd client
	log.Println("Creating etcd client at", whoami, time.Now())
	cli, err = clientv3.New(clientv3.Config{Endpoints: []string{"etcd" + strconv.Itoa(replicas[whoami]) + ":2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create etcd client session
	log.Println("Creating etcd session", time.Now())
	//send TTL updates to server each 1s. If failed to send (client is down or without communications), lock will be released
	session, err = concurrency.NewSession(cli, concurrency.WithTTL(1))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("created session")
	defer session.Close()

	handleRequests()

}

func handleRequests() {
	http.HandleFunc("/do", do)
	log.Println("Listening at port 6000", time.Now())
	log.Fatal(http.ListenAndServe(":6000", nil))
}

func do(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	args := r.URL.Query()
	app := args["app"][0]
	op := args["op"][0]
	granularity := os.Getenv("GRANULARITY")
	oplock := os.Getenv("MODE")
	locktype := os.Getenv("PLACEMENT")
	err := execute(app, op, granularity, oplock, locktype)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func execute(app, op, granularity, oplock, locktype string) error {
	timetosleep, err := getexectime(app, op)
	if err != nil {
		return err
	}

	locks, err := getlocks(app, op, granularity, oplock, locktype)
	if err != nil {
		return err
	}

	reallocks := map[string]*rwlock.RWMutex{}
	for l := range locks {
		var l1 *rwlock.RWMutex
		if val, ok := reallocks[locks[l].Name]; ok {
			l1 = val
		} else {
			l1 = rwlock.NewRWMutex(session, locks[l].Name)
			reallocks[locks[l].Name] = l1
		}
		if locks[l].Mode == "shared" {
			rlerr := l1.RLock()
			defer l1.RUnlock()
			if rlerr != nil {
				return rlerr
			}
		} else {
			wlerr := l1.Lock()
			defer l1.Unlock()
			if wlerr != nil {
				return wlerr
			}
		}
	}
	time.Sleep(time.Duration(timetosleep) * time.Millisecond)
	return nil
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

func getlockconfigs(appname, opname, granularity, oplock, locktype string) ([]OpLock, []LockType, error) {
	oplockfile := "./config/locker/" + appname + "/granular" + granularity + "/oplock" + oplock + ".json"
	oplockf, err := ioutil.ReadFile(oplockfile)
	if err != nil {
		return nil, nil, err
	}
	var oplocks []OpLock
	json.Unmarshal([]byte(string(oplockf)), &oplocks)

	locktypefile := "./config/locker/" + appname + "/granular" + granularity + "/locktype" + locktype + ".json"
	locktypef, err := ioutil.ReadFile(locktypefile)
	if err != nil {
		return nil, nil, err
	}
	var locktypes []LockType
	json.Unmarshal([]byte(string(locktypef)), &locktypes)

	return oplocks, locktypes, nil
}

func getlocks(appname, opname, granularity, oplock, locktype string) ([]Lock, error) {
	oplocks, locktypes, err := getlockconfigs(appname, opname, granularity, oplock, locktype)
	if err != nil {
		return nil, err
	}

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
	sort.Slice(locks, func(i, j int) bool {
		return locks[i].Name < locks[j].Name
	})
	return locks, nil
}
