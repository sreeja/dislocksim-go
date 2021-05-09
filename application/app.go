package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var replicas map[string]int
var whoami string

var lockmanagers map[string]net.Conn

func main() {
	time.Sleep(time.Duration(5000) * time.Millisecond)
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

	lockmanagers = map[string]net.Conn{}

	placements := []string{"houston", "paris", "singapore"}
	for _, place := range placements {
		// create lockmanager connection
		log.Println("Creating lock manager session at replica", place, strconv.Itoa(replicas[place]), time.Now())
		conn, err := net.Dial("tcp", "lm"+place+":8000") //+strconv.Itoa(replicas[place]+1))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		lockmanagers[place] = conn
	}

	// log.Println("reversed")
	// for rep := range placements {
	// 	revrep := len(placements) - 1 - rep
	// 	log.Println(placements[revrep])
	// }

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
	start := time.Now()
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
	// log.Println(params)
	// for each in paramstring.split(","):
	//     kv = each.split("-")
	//     params[kv[0]] = kv[1]

	// print(op, params, flush=True)
	app := os.Getenv("APP")
	granularity := os.Getenv("GRANULARITY")
	oplock := os.Getenv("MODE")
	locktype := os.Getenv("PLACEMENT")
	processtime := time.Since(start)
	err := execute(app, op, params, granularity, oplock, locktype)
	if err != nil {
		elapsed := time.Since(start)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error() + elapsed.String() + "processing time" + processtime.String() + "\n"))
	} else {
		elapsed := time.Since(start)
		w.WriteHeader(http.StatusOK)
		// w.Write([]byte("Successful request! Took " + elapsed.String() + "\n"))
		w.Write([]byte(elapsed.String() + "processing time" + processtime.String() + "\n"))
	}
}

func execute(app, op string, params []string, granularity, oplock, locktype string) error {
	start := time.Now()
	defer logexectime(app, op, start)

	timetosleep := 0 //5000

	locks, err := getlocks(app, op, params, granularity, oplock, locktype)
	if err != nil {
		return err
	}

	locklist := map[string]string{}
	for r, _ := range replicas {
		locklist[r] = ""
	}

	for l := range locks {
		if locks[l].Mode == "shared" {
			locklist[locks[l].Type.Placement] = locklist[locks[l].Type.Placement] + ";S:" + locks[l].Name
		} else {
			locklist[locks[l].Type.Placement] = locklist[locks[l].Type.Placement] + ";X:" + locks[l].Name
		}
	}

	placements := []string{"houston", "paris", "singapore"}
	for rep := range placements {
		if len(locklist[placements[rep]]) > 1 {
			log.Println("Asked locklist", placements[rep], locklist[placements[rep]], time.Now())
			// send to server
			fmt.Fprintf(lockmanagers[placements[rep]], "acquire"+locklist[placements[rep]]+"\n")
			// wait for reply
			_, err := bufio.NewReader(lockmanagers[placements[rep]]).ReadString('\n')
			if err != nil {
				log.Println("Lock acquisition failed", placements[rep], time.Now())
			}
			log.Println("Got locklist", placements[rep], time.Now())
		}
	}

	log.Println("Executing op", time.Now())
	time.Sleep(time.Duration(timetosleep) * time.Millisecond)
	log.Println("Finished execution", time.Now())

	for rep := range placements {
		revrep := len(placements) - 1 - rep
		if len(locklist[placements[revrep]]) > 1 {
			log.Println("Releasing locklist", placements[revrep], locklist[placements[revrep]], time.Now())
			// send to server
			fmt.Fprintf(lockmanagers[placements[revrep]], "release"+locklist[placements[revrep]]+"\n")
			// wait for reply
			_, err := bufio.NewReader(lockmanagers[placements[revrep]]).ReadString('\n')
			if err != nil {
				log.Println("Lock release failed", placements[revrep], time.Now())
			}
			log.Println("Released locklist", placements[revrep], time.Now())
		}
	}

	return nil
}

func logwithtime(text string) {
	log.Println(text, time.Now())
}

func logexectime(app, op string, start time.Time) {
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("started at %v, \n finished at %v", start, t)
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
