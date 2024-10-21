#include <iostream>
#include <stack>
#include <queue>
#include <string>

using namespace std;

class Compare{
	public:
	bool operator()(int a, int b) {
		return a<b;
	}
};

int main() {
	int arr[] = {4, 1, 3, 19, 5, 8};	
	// priority_queue<int, vector<int>, greater<int>> heap;
	priority_queue<int, vector<int>, Compare> heap;
	for(auto x: arr) {
		heap.push(x);
	}

	while(!heap.empty()) {
		cout<<heap.top()<<endl;
		heap.pop();
	}

	return 0;
}
