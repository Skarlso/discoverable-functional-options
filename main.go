package main

import "github.com/Skarlso/discoverable-functional-options/pkg"

func main() {
	serverOpts := pkg.ServerOpts{}
	server := pkg.NewServer(
		serverOpts.WithAddress("address"),
		serverOpts.WithName("name"),
		serverOpts.WithPort(9998))

	server.Start()
}
