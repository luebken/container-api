package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"log"
	"os"
)

// https://github.com/fsouza/go-dockerclient

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: container-api <imagename>")
		os.Exit(1)
	}
	imageName := os.Args[1]

	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Panicf("Initialization error: %v\n", err)
	}
	img, err := client.InspectImage(imageName)
	if err != nil {
		log.Printf("Initialization error: %v\nIs docker running?\n", err)
		os.Exit(1)
	}

	fmt.Println("Container API for: ", imageName)
	fmt.Println("---")
	fmt.Println("Author:  ", img.Author)
	fmt.Println("Created: ", img.Created)
	fmt.Println("Config: ", img.Config.Labels)
	fmt.Println("Ports: ", img.Config.ExposedPorts)
}
