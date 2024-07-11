package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Entry struct {
	Time   string
	IP     string
	Length string
	TTL    string
	ID     string
	SPT    string
	DPT    string
	Window string
}

var templateString = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Firewall Logs</title>
	</head>

	<body>
		<table>
			<tr>
				<th>Time</th>
				<th>IP</th>
				<th>Length</th>
				<th>TTL</th>
				<th>ID</th>
				<th>SPT</th>
				<th>DPT</th>
				<th>Window</th>
			</tr>	
			{{ range . }}
			<tr>
				<td>{{ .Time }}</td>	
				<td>{{ .IP }}</td>
				<td>{{ .Length }}</td>
				<td>{{ .TTL }}</td>
				<td>{{ .ID }}</td>
				<td>{{ .SPT }}</td>
				<td>{{ .DPT }}</td>
				<td>{{ .Window }}</td>
			</tr>
			{{ end }}
		</table>
	</body>

</html>
`

func main() {
	data := []Entry{
		Entry{Time: "Apr 11 23:03:12", IP: "192.168.1.1", Length: "40", TTL: "240", ID: "33841", SPT: "41818", DPT: "777", Window: "1024"},
		Entry{Time: "Apr 12 15:21:15", IP: "194.102.55.12", Length: "60", TTL: "230", ID: "332211", SPT: "32312", DPT: "111", Window: "1024"},
		Entry{Time: "May 14 08:23:01", IP: "72.102.33.109", Length: "80", TTL: "300", ID: "335562", SPT: "50000", DPT: "32313", Window: "2048"},
	}

	tmpl, err := template.New("table").Parse(templateString)
	if err != nil {
		fmt.Println("Error creating template:", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("Error executing template:", err.Error())
			return
		}
	})

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting server:", err.Error())
		return
	}
}
