# Application

tidy:
	go mod tidy

zip:
	go build main.go
	rm -f main.zip
	zip main.zip main
