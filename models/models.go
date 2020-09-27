package models

//User schema of teh user table
type User struct{
	ID int64 `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int64 `json:"age"`
}