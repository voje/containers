package tools

import (
	"io/ioutil"
	"log"
	"os/exec"

	cli "github.com/urfave/cli/v2"
)

func Push(c *cli.Context) {
	// List all images in local repo.
	dockerImages := List(c)

	// Prepare template data.
	tmplData := SkopeoTemplate{
		Registry: TrimHttp(GlobalString["local-addr"]),
		User:     "",
		Password: "",
		Images:   dockerImages,
	}

	// Generate a temporary file with the skopeo template.
	tmpFile, err := ioutil.TempFile("/tmp", "skopeo_push")
	if err != nil {
		log.Fatal(err)
	}
	// TODO defer os.Remove(tmpFile.Name())
	WriteSkopeoTemplate(tmplData, tmpFile)

	// Some info output.
	log.Println(tmpFile.Name())

	// Run the process.
	cmd := exec.Command(
		"skopeo",
		"sync",
		"--src=yaml",
		"--dest=docker",
		"--dest-tls-verify=false",
		"--dest-creds="+c.String("creds"),
		tmpFile.Name(),
		c.String("dst-addr"),
	)

	cmdRun(cmd)

}
