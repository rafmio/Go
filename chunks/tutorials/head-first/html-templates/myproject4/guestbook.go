package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Структура для отображения view.html
type Guestbook struct {
	SignaturesCount int
	Signatures      []string
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file) // создается сканер для содержимого файла
	for scanner.Scan() {							// <- для каждой строки файла
		lines = append(lines, scanner.Text()) // <- ее текст присоединяется к сегменту
	}
	check(scanner.Err())
	return lines
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// viewHandler читает записи гостевой книги и выводит их со счетчиком записей
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html") // создает шаблон на осн.view.html
	check(err)
	guestbook := Guestbook{
		SignaturesCount: len(signatures),
		Signatures:      signatures,
	}
	err = html.Execute(writer, guestbook) // данные структуры вставляются в шаблон,
	check(err)														// а результат вписывается в ResponseWriter
}

// newHandler отображает форму для ввода записи
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") // загружает форму HTML из шаблона
	check(err)
	err = html.Execute(writer, nil) // Записывает шаблон в ResponseWriter
	check(err)											// (нет данных для вставки)
}

// createHandler получает запрос POST с добавляемой записью
// и присоединяет ее к файлу signatures
func createHandler(writer http.ResponseWriter, request *http.Request) {
    signature := request.FormValue("signature") // получает значение поля формы "signature"
		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
		check(err)
		_, err = fmt.Fprintln(file, signature) // добавляет содержимое поля формы в файл
		check(err)
		err = file.Close()
		check(err)
		// перенаправление браузера на основную страницу гостевой книги:
		http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	// запросы на просмотр списка записей будут обрабатываться функцией viewHandler:
	http.HandleFunc("/guestbook", viewHandler)
	// запросы на отправку формы будут обрабатываться createHandler
	http.HandleFunc("/guestbook/new", newHandler)
	// запросы на отправку формы будут обрабатываться функцией createHandler
  http.HandleFunc("/guestbook/create", createHandler)
	// в бесконечном цикле запросы HTTP передаются соответствующим функциям
	// для обработки
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}


// type ResponseWriter interface {
//    Header() Header
//    Write([]byte)(int, error)
//    WriteHeader(statusCode int)
// }
// ResponseWriter используется HTTP-обработчиком для создания HTTP-ответа
//
// type Header map[string][]string
// -----------------------------------------------------------------------
// func HandleFunc(pattern string, handler func(ResponseWriter, *Requset))
// Функция HanldeFunc регистрирует обработчик функции для заданного
// шаблона DefaultServeMux

// func ListenAndServe(addr string, hanlder Handler) error
// Функция ListenAndServe прослушивает сетевой адрес TCP 'addr' а затем
// вызывает Serve с обработчиком для обработки запроса на взодящие соединения

// type Request struct {
//   Method string
//   URL *url.URL
//   Proto string
//   ProtoMajor int
//   ProtoMinor int
//   Header Header
//   Body io.ReadCloser
//   GetBody func() (io.ReadCloser, error)
//   ContentLength int64
//   TransferEncoding []string
//   Close bool
//   Host string
//   Form url.Values
//   PostForm url.Values
//   MultipartForm *multipart.Form
//   Trailer Header
//   RemoteAddr string
//   RequestURI string
//   TLS *tls.ConnectionState
//   Cancel <-chan struct{}
//   Response *Response
// }
// Request представляет собой HTTP-запрос, полученный сервером
// или подлежащий отправке клиентом

// func Open(name string)(*File, error)
// Opens named file for reading. If successful - the associated file descriptor
// has mode O_RDONLY


// type Scanner struct {
//     r            io.Reader // The reader provided by the client.
//     split        SplitFunc // The function to split the tokens.
//     maxTokenSize int       // Maximum size of a token; modified by tests.
//     token        []byte    // Last token returned by split.
//     buf          []byte    // Buffer used as argument to split.
//     start        int       // First non-processed byte in buf.
//     end          int       // End of data in buf.
//     err          error     // Sticky error.
//     empties      int       // Count of successive empty tokens.
//     scanCalled   bool      // Scan has been called; buffer is in use.
//     done         bool      // Scan has finished.
// }
//
// func NewScanner(r io.Reader) *Scanner
// NewScanner returns a new Scanner to read from r

// func (s *Scanner) Scan() bool
// Scan advances the Scanner to the next token, which will then be available
// through the Bytes or Text method.

// func (s *Scanner) Text() string
// Text() returns the most recent token generated by a coll go Scan as a newly
// allocated string holding its bytes

// html.ParseFiles
// func (t *Template) ParseFiles(filenames ...string)(*Template, error)

// html.Execute
// func (*Template) Execute(wr io.Writer, data any) error
// Execute applies parsed template to the specified data object, writing the
// output to wr

// func (r *Request) FormValue(key string) string
// FormValue returns the first value for the named component of the query.
// POST and PUT body parameters take precedence over URL query string values.
