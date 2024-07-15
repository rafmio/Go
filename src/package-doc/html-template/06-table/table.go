package main

import (
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

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("table").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Styled Table</title>
    <style>
        body { font-family: 'lato', sans-serif; }
        .container { max-width: 1000px; margin-left: auto; margin-right: auto; padding-left: 10px; padding-right: 10px; }
        h2 { font-size: 26px; margin: 20px 0; text-align: center; }
        h2 small { font-size: 0.5em; }
        .responsive-table li { border-radius: 3px; padding: 25px 30px; display: flex; justify-content: space-between; margin-bottom: 25px; }
        .table-header { background-color: #95A5A6; font-size: 14px; text-transform: uppercase; letter-spacing: 0.03em; }
        .table-row { background-color: #ffffff; box-shadow: 0px 0px 9px 0px rgba(0,0,0,0.1); }
        .col-1 { flex-basis: 12.5%; }
        .col-2 { flex-basis: 12.5%; }
        .col-3 { flex-basis: 12.5%; }
        .col-4 { flex-basis: 12.5%; }
        .col-5 { flex-basis: 12.5%; }
        .col-6 { flex-basis: 12.5%; }
        .col-7 { flex-basis: 12.5%; }
        .col-8 { flex-basis: 12.5%; }
        @media all and (max-width: 767px) {
            .table-header { display: none; }
            .table-row { display: block; }
            .col { flex-basis: 100%; display: flex; padding: 10px 0; }
            .col:before { color: #6C7A89; padding-right: 10px; content: attr(data-label); flex-basis: 50%; text-align: right; }
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Network Data Table <small>Responsive Design</small></h2>
        <ul class="responsive-table">
            <li class="table-header">
                <div class="col col-1">Time</div>
                <div class="col col-2">IP</div>
                <div class="col col-3">Length</div>
                <div class="col col-4">TTL</div>
                <div class="col col-5">ID</div>
                <div class="col col-6">SPT</div>
                <div class="col col-7">DPT</div>
                <div class="col col-8">Window</div>
            </li>
            {{ range . }}
            <li class="table-row">
                <div class="col col-1" data-label="Time">{{.Time}}</div>
                <div class="col col-2" data-label="IP">{{.IP}}</div>
                <div class="col col-3" data-label="Length">{{.Length}}</div>
                <div class="col col-4" data-label="TTL">{{.TTL}}</div>
                <div class="col col-5" data-label="ID">{{.ID}}</div>
                <div class="col col-6" data-label="SPT">{{.SPT}}</div>
                <div class="col col-7" data-label="DPT">{{.DPT}}</div>
                <div class="col col-8" data-label="Window">{{.Window}}</div>

				            </li>
            {{ end }}
        </ul>
    </div>
</body>
</html>
		`))

	data := []Entry{
		Entry{Time: "Apr 11 23:03:12", IP: "192.168.1.1", Length: "40", TTL: "240", ID: "33841", SPT: "41818", DPT: "777", Window: "1024"},
		Entry{Time: "Apr 12 15:21:15", IP: "194.102.55.12", Length: "60", TTL: "230", ID: "332211", SPT: "32312", DPT: "111", Window: "1024"},
		Entry{Time: "May 14 08:23:01", IP: "72.102.33.109", Length: "80", TTL: "300", ID: "335562", SPT: "50000", DPT: "32313", Window: "2048"},
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
	// To test, open http://localhost:5000 in your browser.
}
