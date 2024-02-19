FROM golang:1.21 AS build-env
WORKDIR /build 

COPY go.mod go.sum ./ 
RUN go mod download

COPY . ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/service-greeting

FROM alpine:3.19
WORKDIR /app

COPY --from=build-env /app/service-greeting .

EXPOSE 8080
CMD ["/app/service-greeting"]
