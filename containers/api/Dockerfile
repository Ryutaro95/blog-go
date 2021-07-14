FROM golang:1.16.3

WORKDIR /go/src/github.com/Ryutaro95/news-go/containers/api/
COPY ./ /go/src/github.com/Ryutaro95/news-go/containers/api/
COPY go.mod go.sum ./

RUN rm -f /etc/localtime && \
ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
go get -u github.com/cosmtrek/air && \
go get github.com/stretchr/testify

COPY . .

ENV GO111MODULE=on

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]