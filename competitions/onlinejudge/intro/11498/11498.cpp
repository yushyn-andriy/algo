#include <iostream>


using namespace std;


int main() {
    int queries;
    int ox, oy;
    int x, y;

    while(cin>>queries) {
        if(queries == 0) break;

        cin>>ox>>oy;
        while(queries != 0) {
            cin>>x>>y;

            if(x < ox && y > oy) cout<<"NO"<<endl;
            if(x > ox && y > oy) cout<<"NE"<<endl;
            if(x > ox && y < oy) cout<<"SE"<<endl;
            if(x < ox && y < oy) cout<<"SO"<<endl;
            if(x == ox || y == oy) cout<<"divisa"<<endl;

            queries--;
        }
    }



    return 0;
}
