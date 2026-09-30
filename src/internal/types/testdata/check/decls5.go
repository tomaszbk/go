package main

// declarations of main
const _, main /* ERROR "cannot declare main" */ , _ = 0, 1, 2
type main /* ERROR "cannot declare main" */ struct{}
var _, main /* ERROR "cannot declare main" */ int
