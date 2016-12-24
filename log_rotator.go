package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
	"net/http"
)


func main() {

	http.HandleFunc("/log", logData)
	fmt.Println("serving on http://localhost:8000/hello")
	http.ListenAndServe("localhost:8000", nil)

	// used for testing
	var quitServer = make(chan bool)
	var quitIndexer = make(chan bool)
	var quitPersist = make(chan bool)
	var quitEvent = make(chan bool)
	go TailLogs("[XRAY-SERVER] ","/Users/chenk/.xray/logs/xray_server.log",quitServer)
	go TailLogs("[XRAY-INDEXER] ","/Users/chenk/.xray/logs/xray_indexer.log",quitIndexer)
	go TailLogs("[XRAY-PERSIST] ","/Users/chenk/.xray/logs/xray_persist.log",quitPersist)
	go TailLogs("[XRAY-EVENT] ","/Users/chenk/.xray/logs/xray_event.log",quitEvent)

	time.Sleep(500 * time.Minute)
}
func TailLogs(logName,logPath string,quit chan bool) {
	t, _ := tail.TailFile(logPath, tail.Config{
		Follow: true,
		ReOpen: true})
	for line := range t.Lines {
		fmt.Println(logName + line.Text)
	}
	select {
	case  <-quit:
		return
	default:
		// do nothing
	}
}

func logData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Listen to log Rotator")
}

