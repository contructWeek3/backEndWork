FROM golang:1.19-alpine3.16

##buat folder APP
RUN mkdir /alta-commerce

##set direktori utama
WORKDIR /alta-commerce

##copy seluruh file
ADD . .

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]