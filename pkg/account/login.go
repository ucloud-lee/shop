package account

import (
	"database/sql"
	"net/http"
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
	t := template.Must(template.ParseFiles("html/login/admin.html"))
	t.Execute(w, nil)
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
