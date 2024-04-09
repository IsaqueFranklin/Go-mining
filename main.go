package main

import (
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "sync"
  "time"

  "github.com/davecgh/go-spew/spew"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)

const difficulty = 1

type Block struct {
  Index int
  Timestamp string 
  BPM int 
  Hash string 
  PrevHash string
  Difficulty int
  Nonce string
}

var Blockchain []Block

type Message struct {
  BPM int
}

var mutex = &sync.Mutex{}

func run() error {
  mux := makeMuxRouter()
  httpAddr := os.Getenv("ADDR")
  log.Println("Listening on ", os.Getenv("ADDR"))
  s := &http.Server{
    Addr: ":" + httpAddr,
    Handler: mux,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  if err := s.ListenAndServe(); err != nil {
    return err
  }

  return nil
}

func makeMuxRouter() http.Handler {
  muxRouter := mux.NewRouter()
  muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
  muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
  return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
  bytes, err := json.MarshalIndent(Blockchain, "", "")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  io.WriteString(w, string(bytes))
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
  w.Header().Set("Content-Type", "application/json")
  response, err := json.MarshalIndent(payload, "", "")
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("HTTP 500: Internal Server Error"))
    return
  }
  w.WriteHeader(code)
  w.Write(response)
}
