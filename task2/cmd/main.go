package main

import (
	"context"
	"os"
	"os/signal"
	"time"
	"fmt"
	"task2/server"
	"task2/client"
)

// for convenience
func printResponseVersion(body []byte, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func printResponseDecode(decodedString string, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodedString)
}

// Жизнь боль без перегрузки
func printResponseHardOp(status bool, code int, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	if status {
		fmt.Printf("%t, %d\n", status, code)
		return
	}
	fmt.Printf("%t\n", status)
}

func main() {
	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		fmt.Println(err)
		return
	}

	client := client.NewClient("http://localhost:8080")

	body, err := client.GetVersion()
	printResponseVersion(body, err)

	decodedString, err := client.PostDecode("R09PRCBHT0JMSU4gSVMgQSBERUFEIEdPQkxJTiBTTyBMRVQnUyBNQUtFIFRIRVNFIEdPQkxJTlMgR09PRCEhISEhIQ==")
	printResponseDecode(decodedString, err)

	status, code, err := client.GetHardOp()
	printResponseHardOp(status, code, err)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown signal recieved")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
	}
	fmt.Println("Server shutdown successfully")
}