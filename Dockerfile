FROM golang:1.22-alpine AS builder
WORKDIR /usr/local/src
# install utils
RUN go install github.com/jackc/tern/v2@latest
# for testing inside docker
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

COPY . .
# build
RUN go build -o ./bin/telenotify ./cmd/telenotify/main.go

# Stage 2: Create a minimal runtime image
FROM alpine AS runner
WORKDIR /app
COPY --from=builder /go/bin/tern ./
COPY --from=builder /go/bin/grpcurl ./
COPY ./migrations ./migrations
COPY --from=builder /usr/local/src/bin/telenotify ./

# Specify the command to run your application
CMD [ "sh", "-c", "/app/tern migrate --migrations ./migrations/ --conn-string $MIGRATION_CONN_STRING && exec /app/telenotify" ]
