#include <iostream>


using namespace std;


int main() {
    int t;
    int min = 100;
    int max = -1;
    int places = 0;

    cin>>t;

    for(int i = 0; i<t; i++) {
        cin>>places;
        if(places == 0) {
            cout<<0<<endl;
            continue;
        }

        for(auto j = 0; j<places; j++) {
            int x;
            cin>>x;
            if(x<min) min = x;
            if(x>max) max = x;  
        }
        cout<<(max - min) *2 <<endl;
        min = 100;
        max = -1;
    }


    return 0;
}
