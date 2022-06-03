package MySqlDB

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cust struct {
	CuId     int
	CuName   string
	CuFamily string
}

type MyConfig struct {
	ServerPort string      `json:"server_port"`
	MYSql      MySQLConfig `json:"mysql"`
}
type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func TestDB() (stat bool, comment string) {

	// Load config from CONFIG.json file
	file, _ := ioutil.ReadFile("CONFIG.json")
	data := MyConfig{}
	_ = json.Unmarshal([]byte(file), &data)

	var PNum string = ":" + data.MYSql.Port

	dns := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		data.MYSql.User,
		data.MYSql.Password,
		data.MYSql.Protocol,
		data.MYSql.Host,
		data.MYSql.Port,
		data.MYSql.DBName,
	)

	var CN_Stat bool = true

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	db.Begin()

	if err != nil {
		CN_Stat = false
	}
	return CN_Stat, PNum
}

func InsertDB() (stat bool, num string) {

	// Load config from CONFIG.json file
	file, _ := ioutil.ReadFile("CONFIG.json")
	data := MyConfig{}
	_ = json.Unmarshal([]byte(file), &data)

	dns := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		data.MYSql.User,
		data.MYSql.Password,
		data.MYSql.Protocol,
		data.MYSql.Host,
		data.MYSql.Port,
		data.MYSql.DBName,
	)

	var Stat bool = true

	db, _ := gorm.Open(mysql.Open(dns), &gorm.Config{})

	cust := Cust{CuId: 1, CuName: "vahid", CuFamily: "mashmooli"}

	result := db.Create(&cust) // pass pointer of data to Create

	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	num = fmt.Sprint(result.RowsAffected) // returns inserted records count

	if result.Error != nil {
		Stat = false
	}
	return Stat, num
}
