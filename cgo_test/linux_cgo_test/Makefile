all:
	gofmt -e -s -w .
	go vet .
	gcc -g -fPIC -c -o lib/libtest.o lib/libtest.c
	gcc -g -fPIC -shared -o libtest.so lib/libtest.o
	ar rcs libtest.a lib/libtest.o
	LD_LIBRARY_PATH=. go run main.go
	LD_LIBRARY_PATH=. go run main2.go
