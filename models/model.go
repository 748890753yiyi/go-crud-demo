package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

type User struct {
	Id 			int `PK`
	Username 	string
	Pwd 		string
}

func getLink() beedb.Model {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.81.18:3306)/test_go?charset=utf8")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func SaveUser(user User) error {
	orm := getLink()
	fmt.Println(user)
	err := orm.Save(&user)
	return err
}

func ValidateUser(user User) error  {
	orm := getLink()
	var u User
	orm.Where("username=? and pwd=?",user.Username,user.Pwd).Find(&u)
	if u.Username == "" {
		return errors.New("yonghuming cuowu")
	}
	return nil
}