.PHONE: knn
knn:
	go build -o bin/knn cmd/knn/main.go ;
	bin/knn -open=dataset/images/00_000200.png -save=save/img.png 1>log.log;