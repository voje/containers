package tools

import (
	"fmt"
)

// Annoying urfave/cli "feature" where you can't user global flags inside commands...
// This is a workaround.
var GlobalString = make(map[string]string)

type DockerImg struct {
	Name string   `json:name`
	Tags []string `json:tags`
}

func (d *DockerImg) ToString() string {
	str := ""
	for _, tag := range d.Tags {
		str += fmt.Sprintf("%s:%s\n", d.Name, tag)
	}
	return str
}
