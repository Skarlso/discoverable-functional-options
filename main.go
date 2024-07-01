package main

import "options-pattern/pkg"

func main() {
	serverOpts := pkg.ServerOpts{}
	server := pkg.NewServer(
		serverOpts.WithAddress("address"),
		serverOpts.WithName("name"),
		serverOpts.WithPort(9998))

	server.Start()
}
