knn:
	go build -o bin/knn cmd/knn/main.go ;

knn-start:
	make knn;
	bin/knn -open=dataset/images/00_000200.png -save=save/img.png 1>log.log;

knn-selection:
	go build -o bin/knn-selection cmd/knn-selection/main.go;

knn-selection-start:
	make knn-selection;
	bin/knn-selection \
		--pathImages="dataset/images" \
		--pathLabels="dataset/labels" \
		--splits="dataset/splits_knn/train.txt" \
		--save="dataset/knn_dataset/labels.csv";