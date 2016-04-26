test:
	go test ./...

install:
	go install ./...

install-deps-mac:
	brew install geos
	go get ./...

install-deps-linux:
	sudo apt-get install libgeos-dev
	go get ./...
