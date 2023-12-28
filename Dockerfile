# Builder image
FROM golang:1.21-alpine as build

RUN mkdir /app 
COPY . /app
WORKDIR /app
RUN go build -ldflags="-s -w" /app/cmd/shawty

# Server image
FROM gcr.io/distroless/static-debian12

COPY --from=build /app/.env /app/.env
COPY --from=build /app/shawty /app/shawty
WORKDIR /app
# docker-compose expects port 8080
# to change the server port on the HOST modify docker-compose.yml
EXPOSE 8080
CMD [ "/app/shawty" ]