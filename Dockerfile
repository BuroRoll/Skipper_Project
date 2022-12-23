FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./pkg ./pkg
COPY ./media ./media
COPY ./index.html ./index.html
#VOLUME /home/ubuntu/static_files ./media/user/profile_picture

RUN go build pkg/cmd/main.go

EXPOSE 8000

CMD [ "./main" ]
