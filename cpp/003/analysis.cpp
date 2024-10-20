#include <iostream>
#include <algorithm>
#include <ctime>
#include <vector>


using namespace std;


void bubble_sort(vector<int> &a, int n) {
	for(int i = 0; i<n; i++) {
		for(int j = i+1; j<n; j++) {
			if(a[j]<=a[i]) {
				int tmp = a[j];
				a[j] = a[i];
				a[i] = tmp;
			}
		}
	}
}

int main() {
	int n;
	cin>>n;
	vector<int> arr(n, 0);
	for(int i = 0; i < n; i++) 
	{
		arr[i] = n - i;
		// cout<<arr[i]<<" ";
	}
	auto b = clock(); 
	sort(arr.begin(), arr.end());
	// bubble_sort(arr, n);
	cout<<endl;
	
	for(int i = 0; i < n; i++) 
	{
		// cout<<arr[i]<<" ";
	}
	cout<<endl<<clock() - b<<" ms"<<endl;
	return 0;
}
