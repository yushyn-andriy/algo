#include <iostream>
#include <vector>
using namespace std;

#define max(i,j) i<j?j:i

class Graph {
	public:
	vector<vector<int>> adj;

	void extend(int max) {
		while(adj.size() <= max) adj.push_back({});	
	};
	
	void addEdge(int i, int j) {
		extend(max(i,j));
		adj[i].push_back(j);
	};

	void print() {
		for(int i = 0; i<adj.size(); i++) {
			if(adj[i].size() > 0) {
				cout<<i<<" --> ";
				for(auto node:adj[i]) cout<<node<<", ";
				cout<<endl;
			}
		}
	};

};

int main() {
	Graph a;
	a.addEdge(1, 2);
	a.addEdge(2, 1);
	a.addEdge(2, 30);
	a.addEdge(20, 30);
	a.addEdge(10, 30);
	a.print();
	return 0;
}

