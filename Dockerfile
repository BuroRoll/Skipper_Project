FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
#VOLUME /home/ubuntu/static_files ./media/user/profile_picture

RUN go build cmd/main.go

EXPOSE 8000

CMD [ "./main" ]
