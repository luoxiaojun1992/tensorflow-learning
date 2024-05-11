.PHONY: generate-samples
generate-samples:
	go run main.go && split -b 100M ../../sample.csv ../../sample.csv. && rm ../../sample.csv
