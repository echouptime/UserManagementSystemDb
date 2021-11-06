package controllers

import (
	"UserManagementSystem/models"
	"UserManagementSystem/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const (
	InsertDataCmd = `insert into user_info (name,department,addr,sex,salary,phone) values (?,?,?,?,?,?)`
	SelectDataCmd = `select id,name,department,addr,sex,salary,phone from user_info`
	QueryDataCmd  = `select * from user_info where id = ?`
	DeleteDataCmd = `delete from user_info where id = ?`
	UpdateDataCmd = `update user_info set name=?,department=?,addr=?,sex=?,salary=?,phone=? where id=?`
)

var (
	SexMap = map[int]string{
		0: "男",
		1: "女",
	}
	user models.User
)

func BaseInformation(w http.ResponseWriter, r *http.Request) {
	users := make([]models.User, 0, 10)
	rows, err := models.InitDB().Query(SelectDataCmd)
	if err == nil {
		for rows.Next() {
			var user models.User
			//将数据库中数据扫描到变量中
			err = rows.Scan(&user.Id, &user.Name, &user.Department, &user.Addr, &user.Sex, &user.Salary, &user.Phone)
			if err == nil {
				//将扫描到的数据添加到切片中
				users = append(users, user)
			} else {
				fmt.Println(err)
			}
		}
	}

	//模板函数
	funcs := template.FuncMap{
		"sex": func(sex int) string {
			return SexMap[sex]
		},
	}

	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("templates/index.html"))
	//传结构体
	if err := tpl.ExecuteTemplate(w, "index.html", struct {
		Users []models.User
	}{users}); err != nil {
		fmt.Println(err)
	}

}

//创建用户
func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("templates/create.html"))
		tpl.ExecuteTemplate(w, "create.html", nil)
	} else if r.Method == "POST" {
		//将用户信息添加到user切片中
		sal, err := strconv.Atoi(strings.TrimSpace(r.FormValue("salary")))
		if err != nil {
			utils.Errors["salary"] = "禁止输入非Int类型"
		}
		sex, _ := strconv.Atoi(strings.TrimSpace(r.FormValue("sex")))

		user = models.User{
			Name:       strings.TrimSpace(r.FormValue("name")),
			Department: strings.TrimSpace(r.FormValue("department")),
			Addr:       strings.TrimSpace(r.FormValue("addr")),
			Sex:        sex,
			Salary:     sal,
			Phone:      strings.TrimSpace(r.FormValue("phone")),
		}

		//插入校验数据
		utils.DataCheck("name", user.Name, 30)
		utils.DataCheck("department", user.Department, 50)
		utils.DataCheck("addr", user.Addr, 50)
		utils.DataCheck("phone", user.Phone, 11)

		salLength := len(strings.TrimSpace(r.FormValue("salary")))
		if salLength == 0 {
			utils.Errors["salary"] = "salary不能为空"
		} else if salLength > 11 {
			utils.Errors["salary"] = "salary字段不能超出11位"
		}

		//判断Errors长度为0则代表无错误,插入数据
		if len(utils.Errors) == 0 {
			if _, err := models.InitDB().Exec(InsertDataCmd, user.Name, user.Department, user.Addr, user.Sex, user.Salary, user.Phone); err == nil {
				fmt.Println("[+]数据插入成功")
			}
		} else {
			tpl := template.Must(template.ParseFiles("templates/create.html"))
			tpl.ExecuteTemplate(w, "create.html", struct {
				User   models.User
				Errors map[string]string
			}{user, utils.Errors})
		}
		//走完添加流程后,清空Map的错误信息
		utils.Errors = map[string]string{}

	}
	http.Redirect(w, r, "/", 302)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		queryId := r.FormValue("id")
		rows := models.InitDB().QueryRow(QueryDataCmd, queryId)
		err := rows.Scan(&user.Id, &user.Name, &user.Department, &user.Addr, &user.Sex, &user.Salary, &user.Phone)
		if err != nil {
			fmt.Println(err)
		}
		tpl := template.Must(template.ParseFiles("templates/update.html"))
		tpl.ExecuteTemplate(w, "update.html", struct {
			User models.User
		}{user})

	} else {
		//解析更新数据
		id := r.FormValue("id")
		name := r.FormValue("name")
		department := r.FormValue("department")
		addr := r.FormValue("addr")
		sex := r.FormValue("sex")
		salary := r.FormValue("salary")
		phone := r.FormValue("phone")

		//更新数据
		if result, err := models.InitDB().Exec(UpdateDataCmd, name, department, addr, sex, salary, phone, id); err == nil {
			fmt.Println("[+]数据更新成功")
			fmt.Println(result.RowsAffected())
		} else {
			fmt.Println(result.RowsAffected())
		}

	}
	http.Redirect(w, r, "/", 302)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//获取到用户ID,直接删除
	deleteId := r.FormValue("id")
	result, err := models.InitDB().Exec(DeleteDataCmd, deleteId)
	if err == nil {
		fmt.Println(result.RowsAffected())
	} else {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", 302)
}
