all: keymaker

keymaker: keymaker.go
	go build keymaker.go
