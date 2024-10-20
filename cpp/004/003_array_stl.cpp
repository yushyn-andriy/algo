#include<iostream>
#include<array>
#include<algorithm>

using namespace std;

// her er stor forksjel, hvis vi bruker & da kan vi 
// forandre verdier av array
// void updateArr(array<int,6> arr, int i, int val) {
void updateArr(array<int,6> &arr, int i, int val) {
	arr[i] = val;
}

void print(array<int, 6> arr) {
	for(int i = 0; i<arr.size(); i++) cout<<arr[i]<<" ";
	cout<<endl;
}

int main() {
	array<int, 6> arr = {1, 2, 3, 4, 5, 6};
	print(arr);
	updateArr(arr, 0, 100);
	print(arr);
	sort(arr.begin(), arr.end());
	print(arr);
	arr.fill(-100);
	print(arr);
	for(auto x: arr) {
		cout<<"auto: "<<x<<" "<<endl;
	}
	cout<<endl;

	return 0;
}

