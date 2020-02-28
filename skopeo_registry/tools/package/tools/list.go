package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cli "github.com/urfave/cli/v2"
)

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
