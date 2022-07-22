##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o /api .

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /api /api

EXPOSE 8080

#ENV TZ Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER nonroot:nonroot

ENTRYPOINT ["/api"]

