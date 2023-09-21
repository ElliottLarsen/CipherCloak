all: keymaker encrypt_client encrypt_server

keymaker: keymaker.go
	go build keymaker.go

encrypt_client: encrypt_client.go
	go build encrypt_client.go

encrypt_server: encrypt_server.go
	go build encrypt_server.go
