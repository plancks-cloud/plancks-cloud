FROM golang:1.11.5-alpine as builder
WORKDIR /github.com/plancks-cloud/plancks-cloud
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o plancks-cloud .

FROM scratch
WORKDIR /
COPY --from=builder /github.com/plancks-cloud/plancks-cloud/plancks-cloud .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/plancks-cloud"]