/**
 *	session相关
 * 1.面向接口编程，2.关于为什么要在main中import memory，可以参考main中import database/sql的实现包，是一样的逻辑。
 */
package main

import (
	"fmt"
	"hello-world/main/session"
	_ "hello-world/main/session/providers/memory" //只需要import就执行了memory中init函数
	"net/http"
	"time"
)

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		if value := sess.Get("username"); value != nil {
			fmt.Println(value)
			fmt.Fprint(w, value)
		}
	} else {
		sess.Set("username", r.Form["username"][0])
	}
	fmt.Println("end")
}

func count(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	w.Header().Set("Content-Type", "text/html")
	if value := sess.Get("countnum"); value != nil {
		fmt.Fprint(w, value)
	}
	fmt.Println("end")
}

func main() {
	fmt.Println("main")
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)
	http.ListenAndServe(":9090", nil)
}
