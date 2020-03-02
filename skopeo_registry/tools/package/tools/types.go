package tools

import (
	"fmt"
	"strings"
)

// Annoying urfave/cli "feature" where you can't user global flags inside commands...
// This is a workaround.
var GlobalString = make(map[string]string)

type DockerImg struct {
	Name string   `json:name`
	Tags []string `json:tags`
}

func NewDockerImg(ref string) (di DockerImg) {
	ref = TrimHttp(ref)
	cs := strings.Split(ref, ":")
	di.Name = strings.Join(cs[:len(cs)-1], ":")
	di.Tags = []string{cs[len(cs)-1]}
	return
}

func (d *DockerImg) ToString() string {
	str := ""
	for _, tag := range d.Tags {
		str += fmt.Sprintf("%s:%s\n", d.Name, tag)
	}
	return str
}

func TrimHttp(url string) string {
	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "http://", "", -1)
	return url
}

type SkopeoTemplate struct {
	Registry string
	User     string
	Password string
	Images   []DockerImg
}
