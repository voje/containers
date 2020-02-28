package tools

import (
	"log"

	cli "github.com/urfave/cli/v2"
)

func push(c *cli.Context) {
	log.Printf("TODO push all docker images to %s, using credentials: %s.\n", c.String("dst-addr"), c.String("cred"))
}
