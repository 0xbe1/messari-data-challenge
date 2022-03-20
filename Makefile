build:
	go build

test:
	go test

sample-run:
	go build && cat sample.txt | ./messari-txagg

run:
	go build && ./stdoutinator_amd64_darwin.bin | ./messari-txagg
