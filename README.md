# Go
git@github.com:rafmio/Go.git

Закрепить ввод с клавиатуры в переменную
Runes, strings, bytes - слабо представляю (https://go.dev/blog/strings)
Закрепить методы
Не до конца понимаю работу горутин/блокировки
Не до конца понимаю работу каналов/блокировки
Повторить каналы
Повторить горутины
Блокировка, условия блокировки, deadlock

time.Now() // func
time.Time // data type
time.Timer // data type
time.Ticker // data type
time.NewTimer
time.NewTicker
time.Stop(), C (field of struct)
time.Reset()

io.`Read`er (type Reader interface)
io.Reader
io.Writer
io.PipeReader
io.PipeWriter

empty interface{}
empty struct {}

panic, recover, defer

  understanding closures
  pointers and methods
  assertions
  static and dynamic types (interface)
  empty interfaces


  type assertion


  What is Err() method?

Learn later:

os.Stdin
os.Stat()
bufio.NewReader(os.Stdin)
reader.ReadString('\n')
strings.TrimSpace


Collection's memory model: arrays, slices, maps, memory, allocations in memory, new variables, copying

Закрепить применение модулей

What is Text() method?

Package "unicode/utf8"
error realization (type error interface{}, Err(), Error() method, etc)
bufio Scanner
bufio NewScanner
io/ioutil
path/filepath
filepath.Join()
ioutil.ReadDir()
ioutil Name()
ioutil IsDir



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

encoding/json
embedded structs
pointers to structs
fields of the struct as a pointers


----------------------------------------------------
type Product struct {
  name, category string
  price float64
}

Defining struct (literal):
kayak := &Product{ "Boat", "Watersports", 150.00 }
(Kayak is a pointer to var kayak of type *Product)

Defining a slice of pointers to struct:
products := []*Product {
  { "Kayak", "Watersports", 275.00 },
  { "Lifejacket", "Watersports", 48.95 },
  { "Soccer Ball", "Soccer", 19.50 },
}
----------------------------------------------------
INTERFACES

Для работы интерфейса должны быть (без применения указателей):
1.Объявлен тип (type VariableType string/struct/[]float64)
2.Объявлена метод с соответствующим типом и параметрами и возвращ.значениями:
    func(v VariableType) methodName()
3.Объявлен интерфейс с включением методов:
    type InterfaceName interface {
      methodName()
      methodName2(float64) []float64
    }
4.Объявлена переменная типа, который поддерживает интерфейс
  newVar := VariableType("text if underlying type is string") или
  newVar2 := VariableType { field1 : value1, field2 : "value2", field3 : 11.33 }

5.Объявить переменную с типом=названием интерфейс и присвоить ей значение
  ранее объявленной переменной с типом:
  var varForInterface InterfaceName = newVar

6.Через точку обращаться к последней переменной за методами:
  varForInterface.methodName()
  varForInterface.methodName2(3.14)

РАЗОБРАТСЬЯ С УКАЗАТЕЛЯМИ В ИНТЕРФЕЙСАХ
???Все вышеуказанные действия копируют данные в методы. Для того, чтобы передавать
по ссылке нужно
  a) varForInterface присвоить адрес newVar:
    var varForInterface InterfaceName = &newVar
  b) в получателях методов расставить * перед типом:
    func(v *VariableType) methodName


--------------------------------------------------------

Regular expressions - basics

//----------------------------------------
Рандомные инты (случайные значения):
1.Подключить  "math/rand" и "time"
2.rand.Seed(time.Now().UnixNano()) - "посеять зерно, от чего будет отталкиваться"
3.rand.Int()
4.инт в промежутке : rand.Intn()

rand.Seed()
