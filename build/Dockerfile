WORKDIR /app


COPY . .
RUN go mod download

RUN go build -o /api-bin ./cmd

EXPOSE 8080

CMD [ "/api-bin" ]
