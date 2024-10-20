#include<iostream>
#include<array>
#include<vector>
#include<deque>
#include<algorithm>
#include<string>
using namespace std;

class Point {
	public:
	int x, y;
		Point() {
			x = 0;
			y = 0;
		};
};

int main() {
	deque<Point> d;
	d.push_back(Point());
	Point p;
	p.x = 100;
	d.push_back(p);
	cout<<d.front().x<<endl; 
	cout<<d.back().x<<endl; 
	cout<<"Size of deque is "<<d.size()<<endl;
	return 0;
}

