FROM golang:1.18

WORKDIR /usr/src/stockticker

COPY . .

#Pull all dependencies
RUN go get -d -v ./
#Install the package
RUN go install -v ./

CMD ["StockTicker"]