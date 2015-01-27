#!/usr/bin/env python

import time

if __name__ == "__main__":
	N = 1000000000
	sum = 0
	i = 0
	start = time.time()
	while True:
		i += 1
		sum += i
		if i == N:
			break
	print("(Python) Run time for {0} loop : {1}ms".format(N, (time.time()-start)*1000))
