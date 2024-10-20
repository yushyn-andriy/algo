#include<iostream>
#include<algorithm>
#include<vector>


using namespace std;

int main() {
	int n;
	cin>>n;
	vector<int> arr(3, 0);
	int N = 1;
	while(n--) {
		for(int i = 0; i<3; i++) cin>>arr[i];
		sort(arr.begin(), arr.end());
		cout<<"Case "<<N<<": "<<arr[1]<<endl;
		N++;
	}

	return 0;
}

