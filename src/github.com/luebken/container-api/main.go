package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fsouza/go-dockerclient"
)

// https://github.com/fsouza/go-dockerclient

func main() {
	// -- usage --
	if len(os.Args) != 2 {
		log.Println("Usage: container-api <imagename:tag>")
		os.Exit(1)
	}
	imageName := os.Args[1]

	// -- try to connect to docker api and find image --
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Panicf("Initialization error: %v\n", err)
	}
	images, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		log.Printf("Initialization error: %v\nIs docker running?\n", err)
		os.Exit(1)
	}
	found := false
	for i := 0; i < len(images) && !found; i++ {
		img := images[i]
		for j := 0; j < len(img.RepoTags) && !found; j++ {
			tag := img.RepoTags[j]
			if imageName == tag {
				found = true
			}
		}
	}
	if !found {
		log.Printf("Couldn't find image: %v\n", imageName)
		os.Exit(0)
	}
	img, err := client.InspectImage(imageName)
	if err != nil {
		log.Printf("Error while InspectImage error: %v\n", err)
		os.Exit(1)
	}

	// -- image analysis --
	fmt.Println(" -----------------------------------------------------------")
	fmt.Println("| Image: ", imageName)
	fmt.Println("|-----------------------------------------------------------")
	fmt.Println("| Author:  ", img.Author)
	fmt.Println("| Size:    ", img.Size/1000/1000, "MB")
	fmt.Println("| Created: ", img.Created.Format("2006-01-02 15:04"))
	fmt.Println("|-----------------------------------------------------------")
	fmt.Println("| Container API:")
	fmt.Println("|")
	var dat []map[string]interface{}
	fmt.Println("| * Expected Links:  ")
	expectedLinks := img.Config.Labels["api.expected-links"]
	if err := json.Unmarshal([]byte(expectedLinks), &dat); err != nil {
		fmt.Printf("|   'Couldn't parse: %v'\n", err)
	}
	for _, o := range dat {
		fmt.Println("|   - Name:         ", o["name"])
		fmt.Println("|   - Port:         ", o["port"])
		fmt.Println("|   - Description:  ", o["description"])
		fmt.Println("|   - Mandatory:    ", o["mandatory"])
	}
	fmt.Println("|")
	fmt.Println("| * Expected ENVs:  ")
	expectedEnvs := img.Config.Labels["api.expected-envs"]
	if err := json.Unmarshal([]byte(expectedEnvs), &dat); err != nil {
		fmt.Printf("|   'Couldn't parse: %v'\n", err)
	}
	for _, o := range dat {
		fmt.Println("|   - ENV:          ", o["key"])
		fmt.Println("|   - Description:  ", o["description"])
		fmt.Println("|   - Mandatory:    ", o["mandatory"])
	}
	fmt.Println("|")
	fmt.Println("| * Expected args:  ")
	expectedArgs := img.Config.Labels["api.expected-args"]
	if err := json.Unmarshal([]byte(expectedArgs), &dat); err != nil {
		fmt.Printf("|   'Couldn't parse: %v'\n", err)
	}
	for _, o := range dat {
		fmt.Println("|   - Arg:          ", o["arg"])
		fmt.Println("|   - Description:  ", o["description"])
	}
	fmt.Println("|")
	fmt.Println("| * Available ports:   ")
	for key, value := range img.Config.ExposedPorts {
		fmt.Println("|     - ", key, value)
	}
	fmt.Println("|")
	fmt.Println("| * Volumes:   ")
	for key, value := range img.Config.Volumes {
		fmt.Println("|     - ", key, value)
	}
	fmt.Println(" -----------------------------------------------------------")
	// TODO
	// add lifecyle hooks
	// add args
}
