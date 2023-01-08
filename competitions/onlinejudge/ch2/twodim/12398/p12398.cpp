#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

unordered_map<char, pair<int, int>> board_map = {
    {'a', {0, 0}},
    {'b', {0, 1}},
    {'c', {0, 2}},
    {'d', {1, 0}},
    {'e', {1, 1}},
    {'f', {1, 2}},
    {'g', {2, 0}},
    {'h', {2, 1}},
    {'i', {2, 2}},
};


vector<vector<int>> solve(string row) {
    vector<vector<int>> board(3, vector<int>(3));
    if (row.empty()) {
        return board;
    }
    for (char ch : row) {
        pair<int, int> curr = board_map[ch];
        int i = curr.first, j = curr.second;
        pair<int, int> u = {i - 1, j};
        pair<int, int> d = {i + 1, j};
        pair<int, int> lf = {i, j - 1};
        pair<int, int> r = {i, j + 1};
        vector<pair<int, int>> coords = {u, d, lf, r, curr};
        for (pair<int, int> coord : coords) {
            int y = coord.first, x = coord.second;
            if (y >= 0 && y < board.size() && x >= 0 && x < board[0].size()) {
                int val = board[y][x];
                val = (val + 1) % 10;
                board[y][x] = val;
            }
        }
    }
    return board;
}


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    int i = 1;
    string row;
    while (getline(cin, row)) {
        vector<vector<int>> board = solve(row);
        cout << "Case #" << i << ":\n";
        for (int i = 0; i<3; i++) {
            for (int j = 0; j < 3; j++) {
                cout << board[i][j] << (j == 3 - 1 ? '\n' : ' ');
            }
        }
        i++;
    }
    return 0;
}

