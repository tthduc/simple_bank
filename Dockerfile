# Build stage
FROM golang:1.19-alpine3.16 AS builder
# Define a current working in directory inside the image
WORKDIR /app
COPY . .
# The first dot means copy everything from the current folder, where we run the docker build
# The second dot is a corrrect working directory inside the image
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.yaml .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
# The dot is current working folder and this dot is represent working dir above
# if have the same tag, the old tag should be remov

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]