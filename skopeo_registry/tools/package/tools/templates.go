package tools

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
