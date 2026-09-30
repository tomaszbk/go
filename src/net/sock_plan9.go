package net

func maxListenerBacklog() int {
	// /sys/include/ape/sys/socket.h:/SOMAXCONN
	return 5
}
