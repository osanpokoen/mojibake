FROM golang:1.16-alpine as build

ENV GO111MODULE off
RUN apk --no-cache add git && go get github.com/kakitamama/mojibake
ENV GO111MODULE on

WORKDIR /go/src/github.com/kakitamama/mojibake
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/mojibake

FROM scratch as exec
COPY --from=build /go/src/github.com/kakitamama/mojibake/bin/mojibake /bin/

ENTRYPOINT ["/bin/mojibake"]