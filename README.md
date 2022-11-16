# Go
git@github.com:rafmio/Go.git

Learn later:
os.Stdin
os.Stat()
bufio.NewReader(os.Stdin)
reader.ReadString('\n')
strings.TrimSpace
rand.Seed()
time.Now() // func
time.Time // data type

What is Text() method?
What is Err() method?

Collection's memory model: arrays, slices, maps, memory, allocations in memory, new variables, copying

Закрепить ввод с клавиатуры в переменную
Закрепить применение модулей

type assertion
Runes
Package "unicode/utf8"
error realization (type error interface{}, Err(), Error() method, etc)
bufio Scanner
bufio NewScanner
io.Reader (type Reader interface)
io/ioutil
path/filepath
filepath.Join()
ioutil.ReadDir()
ioutil Name()
ioutil IsDir


panic, recover, defer

net/http
http.Get (go doc http Get)
http.Response (go doc http Response)
http.ResponseWriter
*http.Request
http.Write
http.HandleFunc
http.ListenAndServe

sync Package
(https://zetcode.com/golang/goroutine/)
sync.WaitGroup
sync.Add()
sync.Done()
sync.Wait()

text/template
template.New()
template.Parse()
template.Execute()

html/template
{{.}} {{if}} {{end}} {{range}}

http(html?).FormValue()

image
image/color
image/gif
