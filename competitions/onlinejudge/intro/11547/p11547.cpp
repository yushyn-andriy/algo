#include <iostream>
#include <string>
#include <cmath>

using namespace std;

int main() {
	int n;
	cin>>n;
	int number = 0;	


	while(n--) {
		string tmp("");
		cin>>number;
		tmp = to_string(abs(((((number * 567)  / 9) + 7492) * 235) / 47 - 498));
		cout<<tmp[tmp.length() - 2]<<endl;
	}

	return 0;
}
