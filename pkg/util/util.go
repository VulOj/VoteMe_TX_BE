package util

import (
	"VoteMe_BE_TX/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func ReadSettingsFromFile(settingFilePath string) (config models.Config) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func CreateTableIfNotExist(db *gorm.DB, tableModels []interface{}) {
	for _, value := range tableModels {
		if !db.HasTable(value) {
			db.CreateTable(value)
			fmt.Println("Create table ", reflect.TypeOf(value), " successfully")
		}
	}
}

func CheckError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}
