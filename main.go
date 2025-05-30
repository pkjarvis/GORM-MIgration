package main

import (
	"fmt"

	"golang.org/x/text/date"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  
}

type Movies struct {
	gorm.Model
	Id uint 
	Name string 
	Duration date


}


func main() {

	dsn := `host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	if err != nil {
		return
	}
	db.AutoMigrate(&User{})

	u := User{Name:"Gigi",Age:26}
	
	// resp := db.Create(&u);

	db.First(&u);
	u.Name="Jinzhu"


	// db.Save(&User{Id: 13, Name: "Jiangly", Age: 100})  -> Updation


	


	

	
	// fmt.Println(resp.Error, resp.RowsAffected)

	// fmt.Println(result.Error,result.RowsAffected)



}
