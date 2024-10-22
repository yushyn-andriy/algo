#include <iostream>
#include <cstdio>

using namespace std;

int main() {

	int x;
	int n;
	cin>>x>>n;
	printf("%.8b\n", x);

	unsigned int one = 0x1;
	
	// set n th bit
	x = (x | (one<<n));
	printf("set %.8b\n", x);
	
	// clear n th bit
	x = (x & (~(0x1<<n)));
	printf("clear %.8b\n", x);

	if( ((x & (one << n))>>n) ) cout<<n<<"th bit is 1"<<endl;
	else
		cout<<n<<"th bit is 0"<<endl;
	return 0;
}
