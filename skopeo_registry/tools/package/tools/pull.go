package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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
		dockerImages = append(dockerImages, NewDockerImg(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Prepare template data.
	tmplData := SkopeoTemplate{
		User:     splitCredentials(c.String("creds"))[0],
		Password: splitCredentials(c.String("creds"))[1],
		Images:   dockerImages,
	}

	// Generate a temporary file with the skopeo template.
	tmpFile, err := ioutil.TempFile("/tmp", "skopeo_pull")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	WriteSkopeoTemplate(tmplData, tmpFile)

	// Some info output.
	log.Println(tmpFile.Name())

	// Run process call: skopeo
	cmd := exec.Command(
		"skopeo",
		"sync",
		"--src-tls-verify=false",
		fmt.Sprintf("--src-creds=%s", c.String("creds")),
		"--src=yaml",
		"--dest-tls-verify=false",
		"--dest=docker",
		tmpFile.Name(),
		TrimHttp(GlobalString["local-addr"]),
	)
	cmdRun(cmd)
}
