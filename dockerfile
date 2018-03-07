FROM golang:1.8
RUN mkdir -p /go/src/github.com/contacts_api_go/
WORKDIR /go/src/github.com/contacts_api_go/
COPY . /go/src/github.com/contacts_api_go/
RUN go get -u github.com/Masterminds/glide
RUN glide install
EXPOSE 5000
RUN go get github.com/pilu/fresh

CMD fresh -c contacts_runner.conf
