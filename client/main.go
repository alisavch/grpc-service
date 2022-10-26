package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alisavch/grpc-service/pkg/api"
)

func convert(client api.HasherClient, request *api.InputNote) (string, error) {
	output, err := client.Convert(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("cannot convert the message")
	}
	return output.Message, nil
}

func newInputMessage(message string) *api.InputNote {
	return &api.InputNote{
		Message: message,
	}
}

func main() {
	var inputMessage string
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:8081", opt)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	fmt.Println("The client is connected to server")

	client := api.NewHasherClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for {
			fmt.Print("Enter the message: ")
			scanner.Scan()
			inputMessage = scanner.Text()
			if len(inputMessage) != 0 {
				fmt.Println(inputMessage)
				break
			} else {
				fmt.Println("Please enter the text")
				continue
			}
		}

		request := newInputMessage(inputMessage)

		response, err := convert(client, request)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Converted message: %v\n", response)
	}
}
