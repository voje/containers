package tools

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func splitCredentials(s string) []string {
	slc := strings.Split(s, ":")
	if len(slc) != 2 {
		log.Fatal("Check your credentials - format should be 'user:password'")
	}
	return slc
}

func cmdRun(cmd *exec.Cmd) {
	var cmdOut, cmdErr bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	log.Println(cmd.Path)
	log.Println(cmd.Args)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		log.Fatal(cmdErr.String())
	}
	log.Printf(cmdOut.String())
}
