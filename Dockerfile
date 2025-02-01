FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ethiocal .


FROM alpine:3.7
WORKDIR /opt
COPY --from=builder /app/ethiocal /opt/ethiocal/
EXPOSE 8080
CMD ["./opt/ethiocal"]
