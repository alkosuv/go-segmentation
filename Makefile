knn:
	go build -o bin/knn cmd/knn/main.go;

knn-start:
	go run -o cmd/knn/main,go \
		--open=dataset/images/00_000200.png \
		--save=save/img.png \
		--label=dataset/knn-dataset/labels.csv \
		1>tmp/log.log;

knn-selection:
	go build -o bin/knn-selection cmd/knn-selection/main.go;

knn-selection-test:
	go run -o cmd/knn-selection/main,go \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits_knn/train_test.txt" \
		--save="dataset/knn-dataset/labels.csv";

knn-selection-start:
	go run -o cmd/knn-selection/main,go \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits_knn/train.txt" \
		--save="dataset/knn-dataset/labels.csv";

kmeans:
	go build -o bin/kmeans cmd/kmeans/main.go;

kmeans-start:
	go run -o cmd/kmeans/main,go \
		--open=dataset/images/00_000200.png \
		--save=save/img.png \
		--label=dataset/knn-dataset/labels.csv \
		1>tmp/log.log;

hmrf:
	go build -o bin/hmrf cmd/hmrf/main.go

hmrf-start:
	go run cmd/hmrf/main.go \
		--open=dataset/images/00_000200.png \
		1>tmp/log.log;