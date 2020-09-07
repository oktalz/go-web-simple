#builder
FROM golang:1.15.1-alpine AS builder

COPY / /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web_server 

#image
FROM scratch

COPY --from=builder /src/web_server .
CMD [ "/web_server" ]
