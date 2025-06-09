FROM golang:1.24.2-alpine3.21 As build
WORKDIR /go/src
COPY . /go/src
RUN cd /go/src && go build -o main
 
FROM alpine:3.21
WORKDIR /app
COPY --from=build /go/src/main /app/
ENTRYPOINT ./main