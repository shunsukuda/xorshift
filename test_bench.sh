#!/bin/bash
RESULT=bench_result.txt
rm $RESULT
echo "\$ go test -bench Seed" >> $RESULT
go test -bench Seed >> $RESULT
echo -e >> $RESULT
echo "\$ go test -bench Uint64" >> $RESULT
go test -bench Uint64 >> $RESULT
echo -e >> $RESULT
echo "\$ go test -bench Int63" >> $RESULT
go test -bench Int63 >> $RESULT
echo -e >> $RESULT
echo "\$ go test -bench Int64" >> $RESULT
go test -bench Int64 >> $RESULT
echo -e >> $RESULT
echo "\$ go test -bench Float64" >> $RESULT
go test -bench Float64 >> $RESULT
echo -e >> $RESULT
echo "\$ go test -bench Jump" >> $RESULT
go test -bench Jump >> $RESULT