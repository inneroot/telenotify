# Stage 1: Build the Go binary
FROM golang:1.21 AS build
WORKDIR /src
COPY . .
RUN go build -o ./bin/telenotify ./cmd/telenotify/main.go

# Stage 2: Create a minimal runtime image
FROM scratch
COPY --from=build /src/bin/telenotify /bin/telenotify

# Specify the command to run your application
CMD ["/bin/telenotify"]
