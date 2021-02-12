knn:
	go build -o bin/knn cmd/knn/main.go;

knn-start:
	go run cmd/knn/main.go \
		--open=dataset/images/00_000200.png \
		--save=save/img.png \
		--label=dataset/soft-dataset/labels.csv \
		1>tmp/log.log;

kmeans:
	go build -o bin/kmeans cmd/kmeans/main.go;

kmeans-start:
	go run cmd/kmeans/main.go \
		--open=dataset/images/00_000200.png \
		--label=dataset/labels/00_000200.png \
		--save=save/img.png \
		--centroid=8 \
		--draw=true \
		1>tmp/log.log;

hmrf:
	go build -o bin/hmrf cmd/hmrf/main.go

hmrf-start:
	go run cmd/hmrf/main.go \
		--open=dataset/images/00_000200.png \
		1>tmp/log.log;

soft-dataset:
	go build -o bin/soft-dataset cmd/knn-selection/main.go;

soft-dataset-test:
	go run cmd/soft-dataset/main.go \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits-general/train_test.txt" \
		--save="dataset/soft-dataset/labels.csv" \
		1>tmp/log.log;

soft-dataset-start:
	go run cmd/soft-dataset/main.go \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits-general/train.txt" \
		--save="dataset/soft-dataset/labels.csv" \
		1>tmp/log.log;