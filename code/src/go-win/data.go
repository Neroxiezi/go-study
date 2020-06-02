package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	data_file_path     = "/cache/weibo"
	dsn                = "root:root@tcp(127.0.0.1:3306)/golang"
	weibo_tpl          = "/tpl/weibo.html"
	init_tpl_file_path = "/tpl/init.tpl"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getWeiBoData() (string, error) {
	str, err := os.Getwd()
	if err != nil {
		return "", err
	}
	f, err := os.Open(str + data_file_path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(fd), nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getDataforDatabase() (string, error) {
	initDB()
	sqlStr := "SELECT `id`,`name`,`pass` FROM weibo_member"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	mjson, _ := json.Marshal(tableData)
	mString := string(mjson)
	data_json, err := set_cache(mString)
	fmt.Println(data_json)
	return data_json, nil
}

func set_cache(mString string) (string, error) {
	str, err := os.Getwd()
	if err != nil {
		return "", err
	}
	f, err := os.Create(str + data_file_path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.Write([]byte(mString))
	if err != nil {
		return "", err
	}
	return mString, nil
}

type AccountInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func update_weibo_data() {
	str, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	status, err := PathExists(str + data_file_path)
	if err != nil {
		log.Fatal(err)
	}
	if status == true {
		data_string, err := getWeiBoData()
		if err != nil {
			log.Fatal(err)
		}
		set_html(data_string)
	} else {
		data_string, err := getDataforDatabase()
		if err != nil {
			log.Fatal(err)
		}
		set_html(data_string)
	}
}

func set_html(data_string string) {
	var account []AccountInfo
	err := json.Unmarshal([]byte(data_string), &account)
	if err != nil {
		fmt.Println(err)
	}
	str, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	userFile := str + weibo_tpl
	init_html, err := get_init_tpl()
	if err != nil {
		fmt.Println(err)
	}
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
	}
	init_html = init_html + "<table class='table'><thead><tr><th>编号</th><th>账号</th><th>密码</th></tr></thead><tbody>"
	for _, v := range account {
		init_html = init_html + "<tr><td>" + v.Id + "</td><td>" + v.Name + "</td><td>" + v.Pass + "</td></tr>"
	}
	init_html = init_html + "</tbody></table></body>\n</html>"
	fout.WriteString(init_html)
}

func get_init_tpl() (string, error) {
	str, err := os.Getwd()
	if err != nil {
		return "", err
	}
	f, err := os.Open(str + init_tpl_file_path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(fd), nil
}
