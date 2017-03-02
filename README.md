# hello-go

Hello go webserver

``docker run -d -p 8080:8080 --rm --name hello-go vensder/hello-go``

Open http://localhost:8080/Hi%20there! or any other random path in your web browser.

``docker logs hello-go``

``docker stop hello-go``
