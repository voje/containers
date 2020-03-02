package main

import (
	"fmt"
	"log"
	"os"

	"github.com/voje/containers.git/skopeo_registry/tools/package/tools"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Tools",
		Usage: "Tools for skopeo and registry:v2",
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "local-addr",
			Usage:       "Local docker registry address. Only set this if you know what you are doing, leave default instead.",
			DefaultText: "http://127.0.0.1:5000",
		},
	}
	app.Before = func(c *cli.Context) error {
		for _, flag := range app.Flags {
			switch v := flag.(type) {
			default:
				continue
			case *cli.StringFlag:
				if v.Value == "" {
					tools.GlobalString[v.Name] = v.DefaultText
				} else {
					tools.GlobalString[v.Name] = v.Value
				}
			}
		}
		return nil
	}
	app.Commands = cli.Commands{
		{
			Name:  "pull",
			Usage: "Pull a list of docker images and save them to local registry.",
			Action: func(c *cli.Context) error {
				tools.Pull(c)
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "list",
					Usage:    "List of tagged docker images. Simple image-per-line format.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "creds",
					Usage:    "Source registry credentials in format: 'user:password'.",
					Required: true,
				},
			},
		},
		{
			Name:  "push",
			Usage: "Push all docker images from local registry to destination.",
			Action: func(c *cli.Context) error {
				tools.Push(c)
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "dst-addr",
					Usage: "Destination registry address; e.g. 'k-vm-repo-server.docker.iskratel.mak:4567/deploy/infrastructure. " +
						"Note that destination image address is created by combining the destination registry address and the image name.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "creds",
					Usage:    "Destination registry credentials in format: 'user:password'.",
					Required: true,
				},
			},
		},
		{
			Name:  "list",
			Usage: "List docker images in local registry.",
			Action: func(c *cli.Context) error {
				imgList := tools.List(c)
				fmt.Printf("%s images:\n-------\n", tools.GlobalString["local-addr"])
				for _, img := range imgList {
					log.Println(img.ToString())
				}
				return nil
			},
		},
	}

	log.Fatal(app.Run(os.Args))
}
