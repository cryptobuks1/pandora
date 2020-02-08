# Pandora
Blockchain-based decentralized platform for issuing, verifying and recovering digital data.


### Usage

#### Node
To run nodes use following commands from the root pandora dir.
```
# start node-1
go run ./cmd/node \
  --log_level=TRACE \
  --listen=/ip4/127.0.0.1/tcp/9999 \
  --private=08011240308bba6aef0385a3c28da15975d8288802964d79b0ad15ac3f53a74ea30ea40e1fdb4e4401e2903b23ad259257785051f7bea9c75d431ffb085381ea1c7aee7f \
  --http_port=8887

# start node-2
go run ./cmd/node \
  --log_level=TRACE \
  --listen=/ip4/127.0.0.1/tcp/9998 \
  --boot=/ip4/127.0.0.1/tcp/9999/p2p/12D3KooWBximBRbRqQnFDJCeG6Mz9LWotj6dZXCoLVy53UsKPhFc \
  --http_port=8886
```

#### Tests
To run tests use following command from the root pandora dir.
```
go test -count=1 -race -cover ./...
```
