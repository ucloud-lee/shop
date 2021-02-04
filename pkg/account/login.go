package account

import (
	"database/sql"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	log "github.com/gogap/logrus"
)

const (
	table = "admin"
)

type admin struct {
	name   string
	passwd string
	email  string
	state  int
	login  int
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user admin
	r.ParseForm()
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("html/login/login.html"))
		t.Execute(w, nil)
	} else {
		username := r.Form["username"]
		password := r.Form["password"]

		db := Mysql()
		if len(username) == 0 || len(password[0]) == 0 {
			t := template.Must(template.ParseFiles("html/login/login.html"))
			t.Execute(w, nil)
			return
		}

		row := db.QueryRow("select name,passwd,login from admin where name=?", username[0])
		if err := row.Scan(&user.name, &user.passwd, &user.login); err != nil {
			t := template.Must(template.ParseFiles("html/login/login.html"))
			t.Execute(w, nil)
			log.Info("管理员：", user.name, " 密码：", user.passwd, " 状态: 登陆失败")
			return
		}
		if user.passwd != password[0] {
			t := template.Must(template.ParseFiles("html/login/login.html"))
			t.Execute(w, nil)
			log.Info("管理员：", user.name, " 密码：", user.passwd, " 状态: 登陆失败")
			return
		}
		login(db, user.login, user.name)
		log.Info("管理员： ", user.name, " 密码：", user.passwd, " 状态: 登陆成功")

		goAdmin(w)
		defer db.Close()
	}

}
func goAdmin(w http.ResponseWriter) {
	curlAddress("http://113.31.106.132:4000/docs/address.html", "html/login/test.html")
	t := template.Must(template.ParseFiles("html/login/test.html"))
	t.Execute(w, nil)
}
func curlAddress(url string, filePath string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("http.Get err=", err)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ioutil.ReadAll err=", err)
		return
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644) //打开文件
	if err != nil {
		log.Error(err)
		return
	}
	io.WriteString(f, string(bytes))
	f.Close()
}
func login(db *sql.DB, login int, name string) { //UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	stmt, err := db.Prepare("UPDATE admin SET login=? where  name=?;")
	if err != nil {
		log.Error(err)
		return
	}
	login += 1
	_, err = stmt.Exec(login, name)
	if err != nil {
		log.Error(err)
		return
	}
}

