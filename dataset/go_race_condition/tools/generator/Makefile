.PHONY: generate-samples
generate-samples:
	go run main.go && shuf ../../sample.csv -o ../../sample_shuffle.csv && split -b 100M ../../sample_shuffle.csv ../../sample.csv. && rm ../../sample.csv && rm ../../sample_shuffle.csv
