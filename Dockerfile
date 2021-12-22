FROM golang:1.14 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o hello-world3 cmd/main.go

FROM alpine:3.12
EXPOSE 8080
WORKDIR /app
COPY --from=build /build/hello-world3 .
CMD ["./hello-world3"]
