package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	cli "github.com/urfave/cli/v2"
)

func Pull(c *cli.Context) {
	// Open file.
	file, err := os.Open(c.String("list"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file.
	// Generate a list of DockerImg objects.
	var dockerImages []DockerImg
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colonSplit := strings.Split(line, ":")
		newImg := DockerImg{
			Name: string(strings.Join(colonSplit[:len(colonSplit)-1], ":")),
			Tags: []string{colonSplit[len(colonSplit)-1]},
		}
		dockerImages = append(dockerImages, newImg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Read template.
	templ, err := template.ParseFiles("../../package/tools/templates/skopeo_template.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Write config to tmp file.
	var buff bytes.Buffer
	templ.Execute(&buff, dockerImages)

	tmpfile, err := ioutil.TempFile("/tmp", "skopeo_config")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// Some info output.
	log.Println(tmpfile.Name())
	log.Println(string(buff.Bytes()))

	// Run process call: skopeo
	cmd := exec.Command(
		"skopeo",
		"sync",
		"--src-tls-verify=false",
		fmt.Sprintf("--src-creds=%s", c.String("creds")),
		"--src=yaml",
		"--dest-tls-verify=false",
		"--dest=docker",
		tmpfile.Name(),
		GlobalString["local-addr"],
	)
	var out bytes.Buffer
	cmd.Stdout = &out

	log.Println(cmd.Path)
	log.Println(cmd.Args)

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", out.String())

}
