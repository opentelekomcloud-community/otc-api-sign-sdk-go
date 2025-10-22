build:
	go build -o target/startECS startECS.go

zip:
	zip -j deploy.zip target/startECS

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip
