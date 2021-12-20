FROM golang:latest


WORKDIR /app

COPY . .

RUN cd tests
RUN go get -u -t ./...
RUN go mod download
RUN go mod tidy

CMD go test -race -count=1 ./...