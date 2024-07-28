#include <iostream>


using namespace std;


int min_sonar_count(int n, int m) {
    return (n / 3) * (m / 3);
}


int main() {

    int t;
    int n, m;
    cin>>t;
    
    for(int i = 0; i<t; i++) {
        cin>>n>>m;
        cout<<min_sonar_count(n, m)<<endl;
    }

    return 0;
}
