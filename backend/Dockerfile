FROM golang:alpine
RUN mkdir /017
ADD . /017
WORKDIR /017
RUN go mod download
RUN go get -v github.com/cosmtrek/air
CMD [ "air"]