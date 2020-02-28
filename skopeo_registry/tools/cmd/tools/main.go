package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	cli "github.com/urfave/cli/v2"
)

type DockerImg struct {
	Name	string	`json:name`
	Tags	[]string	`json:tags`
}

func main() {
	app := &cli.App{
		Name:  "Tools",
		Usage: "Tools for skopeo and registry:v2",
		Commands: []*cli.Command{
			{
				Name:  "pull",
				Usage: "Pull a list of docker images and save them to local registry.",
				Action: func(c *cli.Context) error {
					pull(c)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "list",
						Usage:    "List of tagged docker images. Simple image-per-line format.",
						Required: true,
					},
				},
			},
			{
				Name:  "push",
				Usage: "Push all docker images from local registry to destination.",
				Action: func(c *cli.Context) error {
					push(c)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "dst-addr",
						Usage:    "Destination registry address. (Don't forget the port, e.g. :4567).",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "cred",
						Usage:    "Destination registry credentials in format: 'user:password'.",
						Required: true,
					},
				},
			},
			{
				Name:  "list",
				Usage: "List docker images in local registry.",
				Action: func(c *cli.Context) error {
					imgList := list(c)
					fmt.Printf("%s images:\n-------\n", c.String("local-addr"))
					for _, img := range imgList {
						for _, tag := range img.Tags {
							fmt.Printf("%s:%s\n", img.Name, tag)
						}
					}
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "local-addr",
				Usage:       "Local docker registry address. Only set this if you know what you are doing, leave default instead.",
				DefaultText: "http://127.0.0.1:8080",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func pull(c *cli.Context) {
	log.Printf("TODO pull docker images specified in %s.\n", c.String("list"))
}

func push(c *cli.Context) {
	log.Printf("TODO push all docker images to %s, using credentials: %s.\n", c.String("dst-addr"), c.String("cred"))
}

func list(c *cli.Context) []DockerImg {
	log.Println("TODO list local images")
	resp, err := http.Get(fmt.Sprintf("%s/v2/_catalog", c.String("local-addr")))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type RespJson struct {
		Repositories []string `json:repositories`
	}
	var respJson RespJson
	err = json.Unmarshal(body, &respJson)

	var imgList []DockerImg
	for _, r := range respJson.Repositories {
		resp, err := http.Get(fmt.Sprintf("%s/v2/%s/tags/list", c.String("local-addr"), r))
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var dockerImg DockerImg
		err = json.Unmarshal(body, &dockerImg)
		if err != nil {
			panic(err)
		}
		imgList = append(imgList, dockerImg)
	}
	return imgList
}
