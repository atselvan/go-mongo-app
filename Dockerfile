FROM golang:1.9

RUN mkdir /appl

COPY src /appl/src

ENV GOPATH=/appl \
    GOOS=linux \
    GOARCH=amd64 \
    PATH=$PATH:/appl/src/com/privatesquare/go/mongo-app

RUN cd /appl/src/com/privatesquare/go/mongo-app && go build

EXPOSE 3000

CMD ["mongo-app"]