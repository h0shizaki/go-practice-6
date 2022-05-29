FROM golang:1.18.1-alpine
RUN apk add git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o ./build/go-gql.exe ./main.go

CMD [ "./build/go-gql.exe" ]