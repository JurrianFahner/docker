# Simple https-redirect container
To run the simple https-redirect container: `docker run -p 80:80 -it ensignprojects/https-redirect`.

To build this container use the following commands:
- compile the golang source (including libs): `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirect .`
- build the container: `docker build -t imagename . `

**Since version 2.0** the following environment variables are added to change the behavior of the redirect:

* `HOST` example: `-e HOST=google.com`
* `REDIRECTPATH` example: `-e REDIRECTPATH=true`, this configuration forwards also the path

Full example:

```docker
docker run -it -p 80:80 -e HOST=google.nl -e REDIRECTPATH=false --rm ensignprojects/https-redirect:v2.0
```