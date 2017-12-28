LDFLAGS = "-X main.version=0.0.1"

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	gometalinter \
		--skip vendor \
		./...

build:
	go build -ldflags $(LDFLAGS) -o build/fastly-exporter

clean:
	rm -fr build