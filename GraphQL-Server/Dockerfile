FROM golang:1.18.1-alpine
RUN apk add git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

ENV DBHOST=host.docker.internal
ENV DBPORT=5432
ENV DBUSER=root
ENV DBPASS=4321
ENV DBNAME=postgres_gql

RUN go build -o ./build/go-gql.exe ./main.go

CMD [ "./build/go-gql.exe" ]