#
# Build stage
#

FROM golang:1.15-alpine as compiler

ENV JWT_SECRET_KEY "secret"
WORKDIR /go/src/github.com/mastanca/SALES_MARTIN_STANCANELLI
COPY . .
RUN apk add git bash sqlite build-base \
  && go build -o challenge cmd/web/*.go

#
# Run stage
#

FROM alpine

COPY --from=compiler /go/src /go/src
WORKDIR /go/src/github.com/mastanca/SALES_MARTIN_STANCANELLI
CMD ./challenge
EXPOSE 8080