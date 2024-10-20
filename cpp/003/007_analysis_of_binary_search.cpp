#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int _binary_search(vector<int> &v, int start, int end, int value) {
	if(end<start) return -1;
	int mid = (start + end) / 2;
	cout<<start<<" "<<mid<<" "<<end<<endl;
	if(v[mid] == value) return mid;
	if(value > v[mid]) return _binary_search(v, mid+1, end, value);
	if(value < v[mid]) return _binary_search(v, start, mid-1, value);
	return -1;
};


int binary_search(vector<int> &v, int n, int value) {
	return _binary_search(v, 0, n, value);
};

int main() {
	vector<int> a = {1, 2, 3, 7, 9, 120, 130};
	cout<<binary_search(a, 6, 1)<<endl;
	return 0;
}

