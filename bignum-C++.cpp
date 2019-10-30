#include <iostream>
using namespace std;

int main ()
{
	long a;
	long sum = 0;
	long n = 1000000000;
	cout << "n = " << n << "\n";
	for( a = 0; a < n; a++ ){
		sum += a;
	}
	cout << "sum = " << sum << "\n";
	return 0;
}