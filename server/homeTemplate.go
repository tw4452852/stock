package server

import (
	"html/template"
)

var homeTem = template.Must(template.New("home").Parse(`
<html>
<body>

<form method="GET" action="/query">
sh:<input type="radio" checked="checked" name="kind" value="sh" />
sz:<input type="radio" name="kind" value="sz" />
<input type="text" name="id" />
<input type="submit" />
</form>

</body>
</html>
`))
