# Stock Ticker #
Toy project utilizing Golang, Docker, and Kubernetes

### Building Golang Project ###
From the base directory:
```
$ go build StockTicker
```

### Building Docker Image ###
Note: This image is publicly available as 'alphakappa/stockticker:latest'

From the base directory:
```
$ docker build ../StockTicker -t stockticker:latest
```

### Deploying With Kubernetes ###
From the base directory:
```
$ kubectl apply -f Kubernetes
```