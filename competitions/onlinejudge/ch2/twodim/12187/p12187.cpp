#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

// function to deep copy a 2D vector
vector<vector<int>> deep_copy(const vector<vector<int>> &src) {
    vector<vector<int>> dst(src.size(), vector<int>(src[0].size()));
    for (size_t i = 0; i < src.size(); i++) {
        for (size_t j = 0; j < src[0].size(); j++) {
            dst[i][j] = src[i][j];
        }
    }
    return dst;
}

// function to solve the problem
vector<vector<int>> solve(vector<vector<int>> land, int n, int r, int c, int k) {
    vector<vector<int>> current_land = land;
    vector<vector<int>> next_land = deep_copy(current_land);
    for (int _ = 0; _ < k; _++) {
        for (int i = 0; i < r; i++) {
            for (int j = 0; j < c; j++) {
                int current_brother = current_land[i][j];
                int hated_brother = (current_brother + 1) % n;
                pair<int, int> left = {i, j - 1};
                pair<int, int> right = {i, j + 1};
                pair<int, int> up = {i - 1, j};
                pair<int, int> down = {i + 1, j};
                if (left.second >= 0) {
                    int ir = left.first, ic = left.second;
                    if (current_land[ir][ic] == hated_brother) {
                        next_land[ir][ic] = current_brother;
                    }
                }
                if (right.second < c) {
                    int ir = right.first, ic = right.second;
                    if (current_land[ir][ic] == hated_brother) {
                        next_land[ir][ic] = current_brother;
                    }
                }
                if (up.first >= 0) {
                    int ir = up.first, ic = up.second;
                    if (current_land[ir][ic] == hated_brother) {
                        next_land[ir][ic] = current_brother;
                    }
                }
                if (down.first < r) {
                    int ir = down.first, ic = down.second;
                    if (current_land[ir][ic] == hated_brother) {
                        next_land[ir][ic] = current_brother;
                    }
                }
            }
        }
        current_land = deep_copy(next_land);
    }
    return current_land;
}


int main() {
    int n, r, c, k;
    while (true) {
        cin >> n >> r >> c >> k;
        if (n == 0) {
            break;
        }
        vector<vector<int>> land(r, vector<int>(c));
        for (int i = 0; i < r; i++) {
            for (int j = 0; j < c; j++) {
                cin >> land[i][j];
            }
        }
        vector<vector<int>> new_land = solve(land, n, r, c, k);
        for (int i = 0; i < r; i++) {
            for (int j = 0; j < c; j++) {
                cout << new_land[i][j] << (j == c - 1 ? '\n' : ' ');
            }
        }
    }
    return 0;
}
