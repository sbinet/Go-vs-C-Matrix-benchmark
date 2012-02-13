# Makefile

CC=gcc
CFLAGS=-g -Wall -O2

OBJS=sequential_vector sequential_vector_inter sequential_vector_noacum sequential_matrix sequential_matrix_inter sequential_matrix_noacum
LOGS=sequential1.log sequential2.log sequential3.log

all: $(OBJS)

debug:
	make "CFLAGS=-g -Wall -DDEBUG"

sequential_vector.o: sequential_vector.c
sequential_vector_inter.o: sequential_vector_inter.c
sequential_vector_noacum.o: sequential_vector_noacum.c
sequential_matrix.o: sequential_matrix.c
sequential_matrix_inter.o: sequential_matrix_inter.c
sequential_matrix_noacum.o: sequential_matrix_noacum.c

sequential_vector: sequential_vector.o
	$(CC) $(CFLAGS) -o $@ $@.o

sequential_vector_inter: sequential_vector_inter.o
	$(CC) $(CFLAGS) -o $@ $@.o

sequential_vector_noacum: sequential_vector_noacum.o
	$(CC) $(CFLAGS) -o $@ $@.o

sequential_matrix: sequential_matrix.o
	$(CC) $(CFLAGS) -o $@ $@.o

sequential_matrix_inter: sequential_matrix_inter.o
	$(CC) $(CFLAGS) -o $@ $@.o

sequential_matrix_noacum: sequential_matrix_noacum.o
	$(CC) $(CFLAGS) -o $@ $@.o

test1:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 800 900 3000' ; (time ../bin/$(obj) 800 900 3000) >> sequential1.log 2>&1; echo 'done.';)
test2:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 8000 90 8000' ; (time ../bin/$(obj) 8000 90 8000) >> sequential2.log 2>&1; echo 'done.';)
test3:
	$(foreach obj, $(OBJS), echo 'Starting test $(obj) 8000 10 80000' ; (time ../bin/$(obj) 8000 10 80000) >> sequential3.log 2>&1; echo 'done.';)

tests:
	make test1 && make test2 && make test3

clean:
	rm -f *.o $(OBJS) $(LOGS)
