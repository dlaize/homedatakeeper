FROM askmike/golang-raspbian

RUN apt-get update && apt-get install -y --no-install-recommends git && apt-get clean

ENV GOPATH /gopath
WORKDIR /gopath/src/github.com/dlaize/homedatakeeper

RUN mkdir -p /gopath/src/github.com/dlaize/homedatakeeper
ADD . /gopath/src/github.com/dlaize/homedatakeeper

RUN cd /gopath/src/github.com/dlaize/homedatakeeper
RUN go get github.com/gorilla/mux
RUN go get github.com/influxdata/influxdb/client/v2
RUN go build -o goapp

ENTRYPOINT ["./goapp"]

EXPOSE 8000