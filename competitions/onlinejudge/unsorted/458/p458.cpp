#include<bits/stdc++.h>
using namespace std;


int main()
{
    string s;
    while(cin>>s) {
        for (auto ch : s) 
        {
            cout<<(char)(ch - 7);
        }
        cout<<endl;
    }
}