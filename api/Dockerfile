FROM golang:1.9.2
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY main.go /usr/src/app
RUN go build -o lliwio_api.exe
EXPOSE 8080
CMD ["/usr/src/app/lliwio_api.exe"]