# Simple https-redirect container
To use the simple https-redirect container, you can use `docker pull ensignprojects/https-redirect`.

To build this container use the following commands:
- to compile the golang source: `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirect .`
- to build the container: `docker build -t imagename . `
