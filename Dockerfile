FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o nameDemo .

FROM scratch

COPY --from=builder /app/nameDemo /

ENTRYPOINT ["/nameDemo"]
