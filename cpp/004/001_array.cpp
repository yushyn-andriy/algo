#include<iostream>
using namespace std;

// "int arr[]" is actually the same as int *arr
// det er det samme som int *arr  
void update_arr(int arr[], int i, int val) {
	arr[i] = val;
}

void print(int arr[], int n) {
	cout<<"Size of array is "<<sizeof(arr)<<endl;
	for(int i = 0; i<n; i++) cout<<arr[i]<<" ";
}


int main() {
	int arr[5] = {1, 2, 3, 4, 5};
	int n = sizeof(arr)/sizeof(int);
	print(arr, 5);
	update_arr(arr, 0, 10);
	cout<<endl;
	print(arr, 5);
	cout<<"Size of array is "<<sizeof(arr)<<endl;
	cout<<n<<endl;
	return 0;
}
