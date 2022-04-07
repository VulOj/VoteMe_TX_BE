package services

import (
	"VoteMe_BE_TX/models"
	"VoteMe_BE_TX/pkg/util"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

func dsn(settings models.DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", settings.Username, settings.Password, settings.Hostname, settings.Dbname)
}

//The following variables are defined for Database services

//The following variables are defined for email services
var db *gorm.DB
var RedisClient *redis.Client

func init() {
	databaseInit()
}

func databaseInit() {
	conf := util.ReadSettingsFromFile("Config.json")
	settings := conf.DbSettings
	connStr := dsn(settings)

	dbStr := strings.Replace(connStr, settings.Dbname, "", 1)
	msdb, e := sql.Open("mysql", dbStr)
	util.CheckError(e)
	msdb.Exec("create database if not exists " + settings.Dbname + " character set utf8")
	msdb.Close()

	var err1 error
	db, err1 = gorm.Open("mysql", connStr)
	//db.DB().SetMaxIdleConns(0)
	util.CheckError(err1)

	var temp []interface{}
	var candidates models.Candidates
	var tickets models.Ticket
	temp = append(temp, &candidates, &tickets)
	util.CreateTableIfNotExist(db, temp)

	var u models.Candidates
	db.Where("name=?", "张三").Find(&u)
	if (u == models.Candidates{}) {
		user := models.Candidates{
			Name:  "sam",
			Score: 0,
		}
		db.Create(user)
	}
	db.Where("name=?", "李四").Find(&u)
	if (u == models.Candidates{}) {
		user := models.Candidates{
			Name:  "mike",
			Score: 0,
		}
		db.Create(user)
	}
	db.Where("name=?", "王五").Find(&u)
	if (u == models.Candidates{}) {
		user := models.Candidates{
			Name:  "QQ",
			Score: 0,
		}
		db.Create(user)
	}
	var v models.Ticket
	db.Where("ticketstring=?", "init").Find(&v)
	if (v == models.Ticket{}) {
		ticket := models.Ticket{
			TicketString: "init",
			Time:         time.Now(),
			Limit:        10,
		}
		db.Create(ticket)
	}
}
