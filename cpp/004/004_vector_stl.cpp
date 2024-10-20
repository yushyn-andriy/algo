#include<iostream>
#include<array>
#include<vector>
#include<algorithm>
#include<string>
using namespace std;


int main() {
	vector<string> first;
	
	string names[] = {"A", "B", "Andriy", "Maria", "Svitlana"};
	vector<string> third(names, names+5);
	vector<string> fourth(third);
	for(auto name: fourth) cout<<name<<endl; 
	cout<<"Cap is "<<fourth.capacity()<<endl;	
	cout<<"Size is "<<fourth.size()<<endl;	
	return 0;
}

