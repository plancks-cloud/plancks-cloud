
# Build a Go app build container
FROM golang:1.13.4-alpine3.10 as prebuilder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git


# Download dependencies
FROM prebuilder as depbuilder
WORKDIR /src
COPY go.mod .
RUN go mod download
RUN go mod vendor


# Download build the container
FROM depbuilder as builder
WORKDIR /src
COPY . .
RUN rm /src/go.sum
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Build final and run the application
FROM scratch as final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /app
COPY --chown=nobody:nobody data /.local
COPY --chown=nobody:nobody data /config
USER nobody:nobody
ENTRYPOINT ["/app"]