FROM golang:1.18
RUN mkdir src/app
WORKDIR /src/app
COPY . .
RUN go build -o server .
CMD ["./server"]