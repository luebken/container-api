package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

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
	fmt.Println("| * Required Links:  ")
	fmt.Println("|   - Name:         TODO")
	fmt.Println("|")
	fmt.Println("| * Required ENVs:  ")
	labels := img.Config.Labels

	//envs as a map
	envMap := make(map[string]string)
	for _, envLine := range img.Config.Env {
		s := strings.Split(envLine, "=")
		envMap[s[0]] = s[1]
	}

	//parsing labels for "api.ENV" declarations
	envsInLabels := []string{}
	for label := range labels {
		r1, _ := regexp.Compile("(api.ENV).(.*)")
		envLabel := r1.FindStringSubmatch(label)
		if len(envLabel) > 2 {
			r2, _ := regexp.Compile("^[^.]*$") //no dots => ENV key
			values := r2.FindStringSubmatch(envLabel[2])
			if len(values) > 0 {
				envsInLabels = append(envsInLabels, values[0])
			}
		}
	}

	// printing env and documentation
	for _, env := range envsInLabels {
		fmt.Printf("|   - %v\n", env)
		fmt.Printf("|     > default value : %v\n", envMap[env])
		for label := range labels {
			r1, _ := regexp.Compile("(api.ENV." + env + ").(.*)")
			envLabel := r1.FindStringSubmatch(label)
			if len(envLabel) > 2 {
				fmt.Printf("|     > %-14v: %v\n", envLabel[2], labels[envLabel[0]])
			}
		}
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
