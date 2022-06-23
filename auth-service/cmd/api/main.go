package main

type Config struct {
}

func main() {
	app := Config{}
	// Start gRPC server
	// go app.gRPCListen()
	app.gRPCListen()
}
