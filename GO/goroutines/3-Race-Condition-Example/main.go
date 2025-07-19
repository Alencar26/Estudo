package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var (
	numVisitas uint64 = 0
	m                 = sync.Mutex{}
)

func main() {

	http.HandleFunc("/", RaceCondition)                //com race condition
	http.HandleFunc("/mutex", RaceConditionMutex)      //sem race condition
	http.HandleFunc("/atomic", RaceConditionAtomicSum) //sem race condition

	http.ListenAndServe(":3000", nil)
}

func RaceCondition(w http.ResponseWriter, r *http.Request) {
	numVisitas++
	w.Write([]byte(fmt.Sprintf("Você é o Visitante número %d.", numVisitas)))
}

// RESOLVENDO RACE CONDITION COM MUTEX
func RaceConditionMutex(w http.ResponseWriter, r *http.Request) {

	m.Lock()
	numVisitas++
	m.Unlock()

	w.Write([]byte(fmt.Sprintf("Você é o Visitante número %d.", numVisitas)))
}

// RESOLVENDO RACE CONDITION COM SOMA ATÔMICA
func RaceConditionAtomicSum(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&numVisitas, 1)
	w.Write([]byte(fmt.Sprintf("Você é o Visitante número %d.", numVisitas)))
}
