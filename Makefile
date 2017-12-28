LDFLAGS = "-X main.version=0.0.3"

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	gometalinter \
		--skip vendor \
		./...

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o build/fastly-exporter

clean:
	rm -fr build