#include <iostream>
#include <vector>
#include <sstream>


using namespace std;


struct Coord {
    int r, c;
};

int dr[] = {1, 1, 1, -1, -1, -1, 0, 0};
int dc[] = {-1, 0, 1, -1, 0, 1, -1, 1};



int floodfill(vector<vector<char>> &grid, int r, int c, char c1, char c2) {
    int ans = 0;
    int rows = grid.size();
    int columns = grid[0].size();

    if (r < 0 || r > rows - 1 || c < 0 || c > columns - 1) return 0;
    if (grid[r][c] != c1) return 0;

    ans = 1;

    grid[r][c] = c2;

    for(int d = 0; d<8; d++) {
        ans += floodfill(grid, dr[d] + r, dc[d] + c, c1, c2);
    }

    return ans;
}


int main() {
    int n;
    string s;
    stringstream ss;
    Coord c;

    cin>>n;
    getline(cin, s);
    getline(cin, s);

    for(int i = 0; i < n; i++) {
        vector<vector<char>> grid;
        vector<Coord> coords;
        while (getline(cin, s)) {
            if (s == "") break;

            if (s[0] == 'L' || s[0] == 'W') {
                grid.push_back({});
                for (int j = 0; j<s.length(); j++) {
                    grid[grid.size() - 1].push_back(s[j]);
                }
            } else {
                ss << s;
                ss >> c.r >> c.c;
                ss.str("");
                ss.clear();

                c.r = c.r - 1;
                c.c  = c.c - 1;
                coords.push_back(c);
            }
        }
        
        for(auto c: coords) {
            vector<vector<char>> new_grid = grid;
            cout<<floodfill(new_grid, c.r, c.c, 'W', '.')<<endl;
        }
        
        if(i == n -1) continue;
        cout<<endl;
    } 

    return 0;
}
