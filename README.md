# SRE Coding Challenge

## Temperature Conversion Check

Teperature conversion equations used in this application were found at <https://www.rapidtables.com/convert/temperature/index.html>

### Running with Docker

1) Install Docker (<https://www.docker.com/community-edition>)
2) Build the docker image by running the following command:

```docker
docker build -t sre_challenge .
```

3) After the image is built run the following to use the app:

```docker
docker run -it sre_challenge
```

### Running with Go

1) Install Go 1.10.x (<https://golang.org/doc/install>)
2) Run the following to run the application:

```bash
go run main.go
```

-OR-

Run via executable(Windows) or binary(Linux + Mac) by running the following:

```bash
go build
```

This will create an executable file in the current directory. To run the app run the following:

Linux/Mac

```bash
./SRECodingChallenge
```

Windows

```bash
SRECodingChallenge.exe
```