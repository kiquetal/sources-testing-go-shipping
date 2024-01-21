FROM golang:1.20 AS deps

WORKDIR /hello-api
ADD *.mod *.sum ./
RUN go mod download

FROM deps as devs
ADD . .
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -X main.docker=true"  \
    -o api ./cmd/main.go

CMD ["/hello-api/api"]

FROM scratch as prod
WORKDIR /
EXPOSE 8080
COPY --from=devs /hello-api/api .
CMD ["/api"]
