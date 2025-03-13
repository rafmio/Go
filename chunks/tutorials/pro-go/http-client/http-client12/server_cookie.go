package main

import (
  "net/http"
  "strconv"
  "fmt"
)

func init() {
  http.HandleFunc("/cookie", func(w http.ResponseWriter, r *http.Request) {
    counterVal := 1
    counterCookie, err := r.Cookie("counter")
    if (err == nil) {
      counterVal, _ = strconv.Atoi(counterCookie.Value)
      counterVal++
    }
    http.SetCookie(w, &http.Cookie{
      Name: "counter", Value: strconv.Itoa(counterVal),
    })
    if (len(r.Cookies()) > 0) {
      for _, c := range r.Cookies() {
        fmt.Fprintf(w, "Cookie Name: %v, Value: %v\n",
        c.Name, c.Value)
      }
    } else {
      fmt.Fprintln(w, "Request contains no cookies")
    }
  })
}

// func (r *Request) Cookie(name string) (*Cookie, error)
// Cookie returns the named cookie provided in the request or ErrNoCookie if not
// found. If multiple cookies match the given name, only one cookie will be returned.


// type Cookie struct {
// 	Name  string
// 	Value string
//
// 	Path       string    // optional
// 	Domain     string    // optional
// 	Expires    time.Time // optional
// 	RawExpires string    // for reading cookies only
//
// 	// MaxAge=0 means no 'Max-Age' attribute specified.
// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// 	// MaxAge>0 means Max-Age attribute present and given in seconds
// 	MaxAge   int
// 	Secure   bool
// 	HttpOnly bool
// 	SameSite SameSite
// 	Raw      string
// 	Unparsed []string // Raw text of unparsed attribute-value pairs
// }
