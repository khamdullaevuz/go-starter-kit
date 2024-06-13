package models

type Error struct{
	HttpStatus int
	Data       interface{}
}