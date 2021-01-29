package user

import (
	"html/template"
	"net/http"
	"time"

	log "github.com/gogap/logrus"

	"git.ucloudadmin.com/leesin/shop/pkg/account"
)

func Index(w http.ResponseWriter, r *http.Request) {
	go userview()
	t := template.Must(template.ParseFiles("html/index/index.html"))
	t.Execute(w, nil)
}

func userview() {
	db := account.Mysql()
	stmt, err := db.Prepare("INSERT INTO view(created_at, operator) VALUES(?, ?);")
	if err != nil {
		log.Error(err)
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(t, "enter index")
	if err != nil {
		log.Error(err)
		return
	}

	defer db.Close()
}
