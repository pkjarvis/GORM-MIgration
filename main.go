package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID     uint 
  Name   string
  Age    int
  Gender string
  Email  string
  
}

type Movies struct {
	gorm.Model
	Id uint 
	Name string 
	Runtime string 
	Language string
	Country string
	Genre string

}


type Bookings struct{
	gorm.Model
	Id uint
	BookingDate string
	UserId int64
	ShowId int64
	BookingStatus string
	TheatreId int64
	NoOfSeats int64
	TransactionStatus string
}

type Theatre struct{
	gorm.Model
	Id uint
	Name string
	CityId int64
	Capacity int64
	TotalScreens int64
	ContactId int64
}

type City struct{
	gorm.Model
	Id int64
	Name string
	Zipcode int64
}

type Show struct{
	gorm.Model
	StartTime string
	EndTime string
	TheatreId int64
	Date  string
	MovieId int64
	ScreenHallId int64
}

type ShowSeat struct{
	gorm.Model
	SeatId int64
	Status string
	Price int64 
	ShowId int64
	BookingId int64
}
type Transaction struct {
	gorm.Model
	Id int64
	TransactionAmount float64
	TransactionStatus string
	BookingId int64
	TransactionTime string
}

type Payments struct {
	gorm.Model
	Id int64
	BookingId int64 
	PaymentMethod string 
	PaymentStatus string
	TransactionTime  string
	Timestamp  string

}

type Review struct{
	gorm.Model
	Id int64
	UserId int64 
	Content string
	Stars int64 
	SeatNo int64
}


func main() {

	dsn := `host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	if err != nil {
		return
	}


	db.AutoMigrate(&User{})
	db.AutoMigrate(&Movies{})
	db.AutoMigrate(&Bookings{})
	db.AutoMigrate(&Theatre{})
	db.AutoMigrate(&City{})
	db.AutoMigrate(&Show{})
	db.AutoMigrate(&ShowSeat{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Payments{})
	db.AutoMigrate(&Review{})



	u := User{Name:"Gigi",Age:26}
	m := Movies{Name:"Ghazani",Runtime:"240",Language: "Hindi",Country:"India",Genre:"Adventure,Action"}
	b := Bookings{BookingDate: "21May2025",UserId:1,ShowId: 2,BookingStatus: "booked",TheatreId:3,NoOfSeats: 4,TransactionStatus: "Success"}
	t := Theatre{Name:"NehruTheatre",CityId: 1,Capacity: 2,TotalScreens: 3,ContactId: 4}
	c := City{Name:"New Delhi",Zipcode: 110044}
	s := Show{StartTime: "08:15AM",EndTime: "10:00AM",TheatreId: 2,Date:"22May2025",MovieId: 3,ScreenHallId: 4}
	ss:= ShowSeat{SeatId: 1,Status: "success",Price: 350,ShowId: 3,BookingId: 4}
	tt := Transaction{TransactionAmount: 350,TransactionStatus: "success",BookingId: 4,TransactionTime: "07:34AM"}
	p := Payments{BookingId: 1,PaymentMethod: "online",PaymentStatus: "success",TransactionTime: "success",Timestamp: "11:00AM 22May2025"}
	r := Review{UserId: 1,Content: "wwwwwwww",Stars: 4,SeatNo: 3}



	
	resp    := db.Create(&u);
	resp1   := db.Create(&m);
	resp2   := db.Create(&b);
	resp3   := db.Create(&t);
	resp4   := db.Create(&c);
	resp5   := db.Create(&s);
	resp6   := db.Create(&ss);
	resp7   := db.Create(&tt);
	resp8   := db.Create(&p);
	resp9   := db.Create(&r);


	db.First(&u);
	u.Name="Jinzhu"


	// db.Save(&User{Id: 13, Name: "Jiangly", Age: 100})  -> Updation


	


	

	
	fmt.Println(resp.Error, resp.RowsAffected)
	fmt.Println(resp1.Error,resp1.RowsAffected)
	fmt.Println(resp2.Error,resp2.RowsAffected)
	fmt.Println(resp3.Error,resp3.RowsAffected)
	fmt.Println(resp4.Error,resp4.RowsAffected)
	fmt.Println(resp5.Error,resp5.RowsAffected)
	fmt.Println(resp6.Error,resp6.RowsAffected)
	fmt.Println(resp7.Error,resp7.RowsAffected)
	fmt.Println(resp8.Error,resp8.RowsAffected)
	fmt.Println(resp9.Error,resp9.RowsAffected)

	// fmt.Println(result.Error,result.RowsAffected)



}
