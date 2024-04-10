package models

import "time"

type Blog struct{
	Id int64
	Title string
	Content string
	CreatedAt time.Time
	UserId int64
}

func (e Blog) Save(){
	
}