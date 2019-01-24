FROM alpine:3.8

RUN mkdir /app

WORKDIR /app/

COPY plancks-cloud .
CMD ["/app/plancks-cloud"]