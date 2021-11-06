package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DBUser         = "xx"
	DBPassword     = "xxx"
	DBHost         = "xxx"
	DBPort         = 3306
	Database       = "usermanagesystem"
	Drive          = "mysql"
	CreateTableCmd = `create table if not exists user_info(
			id bigint primary key auto_increment,
			name varchar(30) not null default '' comment '姓名',
			department varchar(50) not null default '' comment '平台归属',
			addr varchar(50) not null default '' comment '用户住址',
			sex tinyint not null default 0 comment '用户性别',
			salary decimal(10.2) not null default 0.00 comment '用户薪资',
			phone char(11) not null default '' comment '用户联系方式'
		) engine=innodb default charset=utf8mb4;`
)

type User struct {
	Id         int
	Name       string
	Department string
	Addr       string
	Sex        int
	Phone      string
	Salary     int
}

func InitDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true", DBUser, DBPassword, DBHost, DBPort, Database)

	//创建数据库 create database usermanagesystem charset=utf8mb4;
	db, err := sql.Open(Drive, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if _, err = db.Exec(CreateTableCmd); err != nil {
		log.Fatal(err)
	}

	return db

}
