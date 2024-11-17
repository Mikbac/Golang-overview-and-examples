package main

type messageModel struct {
	Message1 string `json:"message1" binding:"required"`
	Message2 string `json:"message2" binding:"required"`
}
