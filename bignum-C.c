#include <stdio.h>

int main ()
{
	long a;
	long sum = 0;
	long n = 1000000000;
	printf("n= %ld\n", n);
	for( a = 0; a < n; a++ ){
		sum += a;
	}
	printf("sum= %ld\n", sum);
	return 0;
}
