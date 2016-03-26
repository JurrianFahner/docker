# Simple https-redirect container
To run the simple https-redirect container: `docker run -p 80:80 -it ensignprojects/https-redirect`.

To build this container use the following commands:
- compile the golang source (including libs): `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirect .`
- build the container: `docker build -t imagename . `
