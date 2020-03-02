package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cli "github.com/urfave/cli/v2"
)

// List makes a simple REST query to local (inside conatiner) docker registry and
// prints a lits of all images with tags.
func List(c *cli.Context) []DockerImg {
	resp, err := http.Get(fmt.Sprintf("%s/v2/_catalog", GlobalString["local-addr"]))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type RespJSON struct {
		Repositories []string `json:repositories`
	}
	var respJSON RespJSON
	err = json.Unmarshal(body, &respJSON)

	var imgList []DockerImg
	for _, r := range respJSON.Repositories {
		resp, err := http.Get(fmt.Sprintf("%s/v2/%s/tags/list", GlobalString["local-addr"], r))
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
