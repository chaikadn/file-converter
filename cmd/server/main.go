package main

import "github.com/chaikadn/file-converter/internal/server"

func main() {
	s := server.NewServer("localhost:8080")

	if err := s.Run(); err != nil {
		panic(err)
	}
}
