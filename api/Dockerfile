FROM golang:1.9.2 as builder
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY main.go /usr/src/app
RUN go build -o lliwio_api.exe

FROM nginx:alpine
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/lliwio_api.exe .
COPY default.conf /etc/nginx/conf.d
EXPOSE 80 8080
CMD ["./lliwio_api.exe"]