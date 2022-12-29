/*
*     Author: Yushyn Andriy
*	  Problem: UVA 434 - Matty's Blocks
*/
#include <bits/stdc++.h>

using namespace std;


class Pattern {
    public:
    Pattern(vector<vector<char>> &v) {
        p = v;
        n = v.size();
    }
    bool IsEqual(vector<vector<char>> &p1, vector<vector<char>> &p2) {
        for(int i = 0; i<p1.size(); ++i) {
            for(int j = 0; j<p2.size(); ++j) {
                if (p1[i][j] != p2[i][j]) return false;
            }
        }
        return true;
    }
    private:
        vector<vector<char>> p;
        int n;
};

void printVector(auto p) {
        for(int i = 0; i<p.size(); ++i) {
            for(int j = 0; j<p[i].size(); ++j) {
                cout<<" "<<p[i][j];
            }
            cout<<endl;
        }
}

int main(void) {
    int n;
    string line;
    while(getline(cin, line)) {
       n = atoi(line.c_str());
       vector<vector<char>> p1(n), p2(n);
       const char delim = ' ';
       string s;
       for(int i = 0; i<n; ++i) {
           bool flag = false;
           while(getline(cin, s, delim)) {
                for(int j = 0; j<s.length(); ++j) {
                    if(flag) p1[i].push_back(s[j]);
                    else
                        p2[i].push_back(s[j]);
                }
                flag = true;
           }
       }

        printVector(p1);

    }
    return 0;
}
