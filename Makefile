build:
	go build

test:
	go test

sample-run:
	go build && cat sample.txt | ./messari-txagg
