#include <iostream>
#include <cstdio>

using namespace std;

int main() {

	int x;
	cin>>x;
	if(x & 0x1 == 1) cout<<"Is odd"<<endl;
	else cout<<"Is even"<<endl;
	return 0;
}
