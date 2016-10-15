package main

import (
	// standard
	"flag"
	"log"
	"net/http"
	"os"

	// external
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// config options
	index_files StringArgs
	uploads_dir string
	address     string
	port        string
	addrport    string

	// channels
	requests  chan string
	responses chan *Response

	// loggers
	info *log.Logger
	elog *log.Logger
)

func init() {
	flag.Var(&index_files, "i", "path to yara rules")
	flag.StringVar(&uploads_dir, "uploads", "uploads", "path to uploads directory")
	flag.StringVar(&address, "address", "0.0.0.0", "address to bind to")
	flag.StringVar(&port, "port", "9999", "port to bind to")
	flag.Parse()

	// initialize logger
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	//build address string
	addrport = address + ":" + port
}

func main() {
	// create channels
	info.Println("Initializing channels")
	requests = make(chan string)
	responses = make(chan *Response)

	// create scanner
	info.Println("Initializing scanner")
	scanner, err := NewScanner(requests, responses)
	if err != nil {
		panic(err)
	}

	// load indexes
	for _, index := range index_files {
		info.Println("Loading index: " + index)
		err = scanner.LoadIndex(index)
		if err != nil {
			panic(err)
		}
	}

	// launch scanner
	go scanner.Run()

	//go scanner()

	// setup http server and begin serving traffic
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/scanner/v1/files/", ListHandler).Methods("GET")
	r.HandleFunc("/scanner/v1/files/{filename}", UploadHandler).Methods("PUT")
	r.HandleFunc("/scanner/v1/files/{filename}", DownloadHandler).Methods("GET")
	r.HandleFunc("/scanner/v1/files/{filename}", RemoveHandler).Methods("DELETE")
	r.HandleFunc("/scanner/v1/files/{filename}/scan/", ScanHandler).Methods("GET")
	http.Handle("/", r)
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	http.ListenAndServe(addrport, loggedRouter)
}
