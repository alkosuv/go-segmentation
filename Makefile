knn:
	go build -o bin/knn cmd/knn/main.go;

knn-start:
	make knn;
	bin/knn -open=save/test.png -save=save/img.png 1>tmp/log.log;

knn-selection:
	go build -o bin/knn-selection cmd/knn-selection/main.go;

knn-selection-start:
	make knn-selection;
	bin/knn-selection \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits_knn/train_test.txt" \
		--save="dataset/knn-dataset/labels.csv";