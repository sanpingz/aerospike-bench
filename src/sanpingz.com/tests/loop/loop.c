#include<stdio.h>
#include<sys/time.h>

//struct timeval{
//	long int tv_sec;	// sec
//	long int tv_usec;	// usec
//}

int main(){
	long N = 1000000000;
	long sum = 0;
	long i = 0;
	struct timeval start, end;
	gettimeofday(&start, NULL);
	for (;;) {
		i++;
		sum += i;
		if (i == N) {
			break;
		}
	}
	gettimeofday(&end, NULL);
	printf("(C) Run time for %ld loop: %fms\n", N, 1000*(end.tv_sec-start.tv_sec)+(end.tv_usec-start.tv_usec)/1000.0);
	return 0;
}
