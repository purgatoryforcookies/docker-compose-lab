FROM golang:1.15-buster
WORKDIR /src
COPY . /src
RUN go mod download

RUN CGO_ENABLED=0 go build -o /server1

CMD ["/server1"]