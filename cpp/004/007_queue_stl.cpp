#include <iostream>
#include <stack>
#include <string>
#include <queue>
using namespace std;


int main() {
	queue<string> books;

	books.push("C++");
	books.push("Iskra");
	while(!books.empty()) {
		cout<<books.front()<<endl;
		books.pop();
	}

	return 0;
}
