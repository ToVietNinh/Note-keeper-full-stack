FROM golang:alpine AS build
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /bin/yma-backend

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/config /config
COPY --from=build /bin/yma-backend /bin/yma-backend
EXPOSE 8000
ENTRYPOINT ["/bin/yma-backend"]
