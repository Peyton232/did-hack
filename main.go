package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// test cases
// 1. presdenetial election 2024
// 2. Houstan mayoral election 2023
// 3. University of North City pizza topping vote

// crud elections
// create election
// get result of election
// I am the authorirty for this elecetioon here is my did
// wallet connect - so you can say here is my wallet

// crud on voter - don't store - creted and done at vote time - connect wallet?
// create voter - when you say I want to vote on this - ok then give authorizale credential
// represented by dids and be verifiable

// vote on election
// I'd like to vote, then auth their verifiable credential given from a wallet
//i.e. I am ben linville and I can vote in this election, I'm a resident of X and
// verifiable credential when I try to vote

// crux - verify a vote that was cast was voted as it was cast
