prefix?=/usr/local/bin
dockerimage?=jmervine/gitfish

gitfish:
	go build

clean:
	rm gitfish

install:
	cp gitfish $(prefix)/gitfish

test:
	go test -v

docker:
	docker build -t $(dockerimage) .

