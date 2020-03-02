package tools

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

// It seems we need to include the user credentials in template if we're using src=yaml

// If we leave .Registry blank (""), we can use full iamge paths int .Images
// The end result is always .Registry/.Images[i]
const skopeo_template = `
"{{ .Registry }}":
    tls-verify: false
    credentials:
        username: {{ .User }}
        password: {{ .Password }}
    images:
        {{- range $x := .Images }}
        {{ $x.Name }}: {{ $x.Tags }}
        {{- end }}

`

// GenSkopeoTemplateTmpFile generates a skopeo template from data in SkopeoTemplate,
// saves the template to a temporary file and returns the filename.
func WriteSkopeoTemplate(d SkopeoTemplate, f *os.File) {

	// Read template from string.
	templ := template.Must(template.New("skopeo_template").Parse(skopeo_template))

	// Write config to tmp file.
	var buff bytes.Buffer
	templ.Execute(&buff, d)

	if _, err := f.Write(buff.Bytes()); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// Some info output.
	// log.Println(string(buff.Bytes()))
}
