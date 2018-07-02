FROM golang

COPY ./app /go/src/github.com/user/myProject/app
WORKDIR /go/src/github.com/user/myProject/app

RUN go get ./
RUN go build

# CMD go get github.com/pilu/fresh && fresh

CMD app;