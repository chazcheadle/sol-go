# Stage 1
# Build stage to compile Go binary
FROM golang:1-alpine as build
ADD ./sol.go /go/src
RUN apk add build-base
ADD ./JPLEPH /JPLEPH
WORKDIR /go/src

RUN go mod init sol
RUN go mod tidy
RUN go build sol.go

# Stage 2
# Build stage to run minimal OS and binary
FROM alpine:latest
WORKDIR /app
# Copy executable to new image
COPY --from=build /go/src/sol /app/sol
COPY --from=build /JPLEPH /go/pkg/mod/github.com/pebbe/novas@v1.1.1/jpleph/JPLEPH
RUN du /go
ENV GOPATH=/go
RUN echo $GOPATH
EXPOSE 3000
ENTRYPOINT ["./sol"]
