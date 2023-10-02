FROM golang:1.21.1
WORKDIR /go/hello-world-go
COPY . . 
RUN go mod init github.com/Guilherme-Joviniano/hello-world-go
RUN go build -o server .
CMD [ "./server" ] 
