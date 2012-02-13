#include $(GOROOT)/src/Make.inc

OBJS=\
 sequential_vector \
 sequential_vector_inter \
 sequential_vector_noacum \
 sequential_matrix \
 sequential_matrix_inter \
 sequential_matrix_noacum \
 sequential_vector-gccgo \
 sequential_vector_inter-gccgo \
 sequential_vector_noacum-gccgo \
 sequential_matrix-gccgo \
 sequential_matrix_inter-gccgo \
 sequential_matrix_noacum-gccgo \


LOGS=\
 gosequential1.log \
 gosequential2.log \
 gosequential3.log \
 gosequential1-gccgo.log \
 gosequential2-gccgo.log \
 gosequential3-gccgo.log \


GCCGOFLAGS=-O2

all: $(OBJS)

sequential_vector:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_vector_inter:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_vector_noacum:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_matrix:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_matrix_inter:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_matrix_noacum:
	go tool 6g -B $@.go
	go tool 6l -o $@ $@.6

sequential_vector-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

sequential_vector_inter-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

sequential_vector_noacum-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

sequential_matrix-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

sequential_matrix_inter-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

sequential_matrix_noacum-gccgo:
	gccgo $(GCCGOFLAGS) -o $@.o -c $(subst -gccgo,,$@).go
	gccgo $(GCCGOFLAGS) -o $@ $@.o

test1:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 800 900 3000' ; (time ../../bin/$(obj) 800 900 3000) >> gosequential1.log 2>&1; echo 'done.';)
test2:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 8000 90 8000' ; (time ../../bin/$(obj) 8000 90 8000) >> gosequential2.log 2>&1; echo 'done.';)
test3:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 8000 10 80000' ; (time ../../bin/$(obj) 8000 10 80000) >> gosequential3.log 2>&1; echo 'done.';)

tests:
	make test1 && make test2 && make test3

clean:
	rm -f *.6 *.o $(OBJS) $(LOGS)
