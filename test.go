/**
 *通用实验
 */

package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"fmt"
)


func main() {
	//fmt.Print(sessionId())
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func init() {
	http.Handle("/view", appHandler(viewRecord))
	http.Handle("/view1", appHandler(viewRecord1))
	http.ListenAndServe(":9090",nil)
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError{
	if _,err := os.Open("hehe.txt"); err != nil {
		return &appError{err, "Record not found", 404}
	}
	return nil
}

func viewRecord1(w http.ResponseWriter, r *http.Request) *appError{
	if _,err := os.Open("hehe.txt"); err != nil {
		return &appError{err, "Can't display record", 500}
	}
	return nil
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
		fmt.Println()
	}
}

type appError struct {
	Error   error
	Message string
	Code    int
}
