#!/bin/sh

function run_tests()
{
    progargs=(800 900 3000)

    echo ""
    echo ":::::::::::::::::::::::::::::"
    echo ":: c sequential-matrix..."
    time ./C/sequential/sequential_matrix "${progargs[@]}"
    echo ""
    echo ":: gc sequential-matrix..."
    time ./go/sequential/sequential_matrix "${progargs[@]}"
    echo ""
    echo ":: gccgo sequential-matrix..."
    time ./go/sequential/sequential_matrix-gccgo "${progargs[@]}"

    echo ""
    echo ":::::::::::::::::::::::::::::"
    echo ":: c sequential-vector..."
    time ./C/sequential/sequential_vector "${progargs[@]}"
    echo ""
    echo ":: gc sequential-vector..."
    time ./go/sequential/sequential_vector "${progargs[@]}"
    echo ""
    echo ":: gccgo sequential-vector..."
    time ./go/sequential/sequential_vector-gccgo "${progargs[@]}"
}

function make_c_targets()
{
    (cd C/sequential; make clean; make)
}

function make_go_targets()
{
    (cd go/sequential; make clean; make)
}

function run() {
    make_c_targets
    make_go_targets

    run_tests
}

run
