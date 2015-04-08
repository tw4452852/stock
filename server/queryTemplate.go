package server

import (
	"html/template"
)

var queryTem = template.Must(template.New("query").Parse(`
<html>
<body>
<p> key is {{.Key}} </p>
{{range .Pairs}} <p> {{.Name}} : {{.Value}} </p> {{end}}
</body>
</html>
`))
