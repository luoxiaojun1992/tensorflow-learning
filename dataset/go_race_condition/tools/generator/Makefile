.PHONY: generate-samples
generate-samples:
	go run main.go && shuf ../../sample.csv -o ../../sample_shuffle.csv && echo "code,label" >> ../../sample_shuffle_with_label.csv && cat ../../sample_shuffle.csv >> ../../sample_shuffle_with_label.csv && split -b 100M ../../sample_shuffle_with_label.csv ../../sample.csv. && rm ../../sample.csv && rm ../../sample_shuffle.csv && rm ../../sample_shuffle_with_label.csv
