This benchmark was to test how much overhead channels(buffered or unbuffered) have by testing sending single bytes vs arrays of bytes across a chanel.

( .)(. ) Go>time ./ArrayByteStream ArrayByteStream.go 

real    0m0.008s
user    0m0.004s
sys 0m0.004s

( .)(. ) Go>time ./ChanByteStream ArrayByteStream.go 

real    0m0.043s
user    0m0.038s
sys 0m0.003s


//10
( .)(. ) Go>go build BufferedChanByteStream.go 
( .)(. ) Go>time ./BufferedChanByteStream ArrayByteStream.go 

real    0m0.029s
user    0m0.024s
sys 0m0.004s

//25
( .)(. ) Go>go build BufferedChanByteStream.go 
( .)(. ) Go>time ./BufferedChanByteStream ArrayByteStream.go 

real    0m0.027s
user    0m0.021s
sys 0m0.003s

//50
( .)(. ) Go>go build BufferedChanByteStream.go 
( .)(. ) Go>time ./BufferedChanByteStream ArrayByteStream.go 

real    0m0.025s
user    0m0.020s
sys 0m0.003s

//1024 (for good measure)
( .)(. ) Go>go build BufferedChanByteStream.go 
( .)(. ) Go>time ./BufferedChanByteStream ArrayByteStream.go 

real    0m0.024s
user    0m0.019s
sys 0m0.003s