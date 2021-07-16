all:
	go build -buildmode=c-archive -o libgotest.a wrapper.go
	go build -buildmode=c-shared -o libgotest.so wrapper.go
	go build -buildmode=c-shared -o libgotest1.so wrapper1.go
	gcc -o a.out wrapper.c libgotest.a
	gcc -o a.out wrapper.c libgotest.so
	LD_LIBRARY_PATH=. gcc -o a.out wrapper.c libgotest.so