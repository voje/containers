package tools

type DockerImg struct {
	Name string   `json:name`
	Tags []string `json:tags`
}
