FROM golang:1.14

WORKDIR /go/src

ENV SRC_DIR=./gitbub.com/stetsd/monk-app\
    GOOS=linux \
    GOARCH=amd64

COPY . $SRC_DIR

RUN cd $SRC_DIR; make build; mv monkapp /usr/local/bin;

CMD ["monkapp", "start"]