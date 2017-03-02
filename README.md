# Small Golang webserver in Docker

### How to run:

``docker run -d -p 8080:8080 --rm --name hello-go vensder/hello-go``

Open [http://localhost:8080/Hi%20there!](http://localhost:8080/Hi%20there!) or any other random path after slash in your web browser.

View logs:

``docker logs hello-go``

Stop container:

``docker stop hello-go``

### How to build your own image if you don't have Go (but have a docker):

``git clone https://github.com/vensder/hello-go.git``

``cd hello-go``

``docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8-alpine go build -v webserver.go``

Make binary file smaller:

``strip webserver``

`` mkdir -p bin/ && mv webserver bin/``

Build docker image:

``docker build -t hello-go .``

So you got smallest docker image with built-in web-server, it has size less than 8 MB!
You can use it even in embedded systems.

