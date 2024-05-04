
all: clean l2scan-indexer trace-indexer

l2scan-indexer:
	go build ./

trace-indexer:
	go build -o bin/trace_indexer cmd/trace-indexer/*

clean:
	rm ./l2scan-indexer