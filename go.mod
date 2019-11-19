module github.com/status-im/status-eth-node

go 1.13

replace github.com/ethereum/go-ethereum v1.9.5 => github.com/status-im/go-ethereum v1.9.5-status.5

require (
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.5
	github.com/mattn/go-pointer v0.0.0-20190911064623-a0a44394634f
	github.com/status-im/whisper v1.5.2
)
