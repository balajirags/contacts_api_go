FROM golang:latest as builder
RUN mkdir -p $GOPATH/src/github.com/contacts_api_go/
COPY . $GOPATH/src/github.com/contacts_api_go/
WORKDIR $GOPATH/src/github.com/contacts_api_go/
RUN go get -u github.com/Masterminds/glide
RUN glide install
RUN mkdir -p $GOPATH/src/github.com/contacts_api_go/out/migrations
COPY ./migrations $GOPATH/src/github.com/contacts_api_go/out/migrations
COPY ./application.yml $GOPATH/src/github.com/contacts_api_go/out/
RUN	go build -o $GOPATH/src/github.com/contacts_api_go/out/contacts

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 5000

CMD ["dockerize", "-wait", "tcp://db:5432","-timeout", "120s", "/go/src/github.com/contacts_api_go/out/contacts", "start"]
