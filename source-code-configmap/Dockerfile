FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/simple-go-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM scratch
COPY --from=build-env /tmp/simple-go-app/gitops-promotion-source-code /app/gitops-promotion-source-code

# COPY settings.ini /config/settings.ini


EXPOSE 8080
CMD ["/app/gitops-promotion-source-code"]
