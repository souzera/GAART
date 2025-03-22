FROM baseImage

WORKDIR /

COPY . .

RUN mkdir -p /data

RUN mkdir -p /cached

RUN go mod tidy

RUN go build -o /cmd/api/

RUN go build -o /cmd/cli/

EXPOSE 8000

CMD [ "main" ]

