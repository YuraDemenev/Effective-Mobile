FROM golang:1.22.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /serve ./cmd/main.go


FROM debian:bookworm

WORKDIR /

COPY --from=build-stage /serve /serve

EXPOSE 8080

ENTRYPOINT ["/serve"]

# pg_dump.exe -U postgres -d Effective_Mobile -f D:\effective_mobile.sql
