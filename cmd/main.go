package main

import "github.com/studybox/auth"

func main() {
	conf := auth.Config{
		HttpPort: 8090,
	}
	service := auth.NewService(conf)
}
