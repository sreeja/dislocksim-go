package main

import (
	"github.com/sreeja/etcd-exp/rwlock"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"

	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var replicas map[string]int
var whoami string

// var cli *clientv3.Client
var sessions map[string]*concurrency.Session

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

	sessions = map[string]*concurrency.Session{}

	placements := []string{"cent", "clust", "dist"}
	for place := range placements {
		// create etcd client
		etcd_client_name := "etcd" + strconv.Itoa(replicas[whoami]) + "-" + placements[place] + ":2379"
		log.Println("Creating etcd client", etcd_client_name, time.Now())
		cli, err := clientv3.New(clientv3.Config{Endpoints: []string{etcd_client_name}})
		if err != nil {
			log.Fatal(err)
		}
		defer cli.Close()

		// create etcd client session
		log.Println("Creating etcd session", time.Now())
		//send TTL updates to server each 1s. If failed to send (client is down or without communications), lock will be released
		session, err := concurrency.NewSession(cli, concurrency.WithTTL(1))
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		sessions[placements[place]] = session
	}

	handleRequests()

}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/do", do).Methods("GET")
	router.HandleFunc("/do", do).Methods("POST")
	router.HandleFunc("/do", do).Methods("PUT")
	router.HandleFunc("/do", do).Methods("DELETE")
	log.Println("Listening at port 6000", time.Now())
	log.Fatal(http.ListenAndServe(":6000", router))
}

func do(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	op := args["op"][0]
	paramstring := args["params"][0]
	// log.Println(paramstring)
	params := []string{}
	splitparams := strings.Split(paramstring, ",")
	for each := range splitparams {
		kv := strings.Split(splitparams[each], "-")
		params = append(params, kv[1])
	}
	log.Println(params)
	// for each in paramstring.split(","):
	//     kv = each.split("-")
	//     params[kv[0]] = kv[1]

	// print(op, params, flush=True)
	app := os.Getenv("APP")
	granularity := os.Getenv("GRANULARITY")
	oplock := os.Getenv("MODE")
	locktype := os.Getenv("PLACEMENT")
	err := execute(app, op, params, granularity, oplock, locktype)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error() + "\n"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successful request!\n"))
	}
}

func execute(app, op string, params []string, granularity, oplock, locktype string) error {
	start := time.Now()
	defer logexectime(app, op, start)

	timetosleep, err := getexectime(app, op)
	if err != nil {
		return err
	}

	locks, err := getlocks(app, op, params, granularity, oplock, locktype)
	if err != nil {
		return err
	}

	// reallocks := map[string]*rwlock.RWMutex{}
	for l := range locks {
		// var l1 *rwlock.RWMutex
		// if val, ok := reallocks[locks[l].Name]; ok {
		// 	l1 = val
		// } else {
		l1 := rwlock.NewRWMutex(sessions[locks[l].Type.Placement], locks[l].Name)
		// log.Println(l1.s.Client())
		// 	reallocks[locks[l].Name] = l1
		// }
		if locks[l].Mode == "shared" {
			log.Println("Asked read lock", locks[l].Name, time.Now())
			rlerr := l1.RLock()
			defer l1.RUnlock()
			defer logwithtime("Releasing read lock")
			if rlerr != nil {
				return rlerr
			}
			log.Println("Got read lock", time.Now())
		} else {
			log.Println("Asked lock", locks[l].Name, time.Now())
			wlerr := l1.Lock()
			defer l1.Unlock()
			defer logwithtime("Releasing lock")
			if wlerr != nil {
				return wlerr
			}
			log.Println("Got lock", time.Now())
		}
	}
	log.Println("Executing op", time.Now())
	time.Sleep(time.Duration(timetosleep) * time.Millisecond)
	log.Println("Finished execution", time.Now())
	return nil
}

func logwithtime(text string) {
	log.Println(text, time.Now())
}

func logexectime(app, op string, start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("started at %v, finished at %v", start, t)
	log.Printf("execution time for %v, %v :%v", app, op, elapsed)
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

func getlocks(appname, opname string, params []string, granularity, oplock, locktype string) ([]Lock, error) {
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
						lockName := oplocks[o].Locks[l].Name + strings.Join(params, "_")
						lock := Lock{lockName, oplocks[o].Locks[l].Mode, locktypes[t]}
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
