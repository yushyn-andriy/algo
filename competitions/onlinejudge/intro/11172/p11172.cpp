#include <iostream>

using namespace std;


string get_operator(int x, int y) {
    switch (x>=y)
    {
    case true:
        if(x>y) return ">";
        return "=";
    default:
        return "<";
    }
}


int main() {

    int t;
    int x, y;
    cin>>t;
    
    for(int i = 0; i<t; i++) {
        cin>>x>>y;
        cout<<get_operator(x, y)<<endl;
    }

    return 0;
}
