FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/api-bin ./cmd


FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/api-bin /app/api-bin
COPY --from=0 /app/configs/app.env /app/configs/app.env
EXPOSE 8080
CMD [ "/app//api-bin" ]
