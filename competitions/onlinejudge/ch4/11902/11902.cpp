#include <iostream>
#include <unordered_map>
#include <vector>

class Graph {
    private:
        std::unordered_map<int, std::unordered_map<int, bool>> _outgoing;
        std::unordered_map<int, std::unordered_map<int, bool>> _incoming;
        bool directed;
    public:
        Graph(bool directed = false) : directed(directed) {
            if (!directed) {
                _incoming = _outgoing;
            }
        }

        bool is_directed() {
            return _outgoing != _incoming;
        }

        std::vector<int> vertices() {
            std::vector<int> v;
            for (auto &i : _outgoing) {
                v.push_back(i.first);
            }
            return v;
        }

        std::vector<int> incident_vertices(int v, bool outgoing = true) {
            std::unordered_map<int, std::unordered_map<int, bool>> adj;
            if (outgoing) {
                adj = _outgoing;
            } else {
                adj = _incoming;
            }

            if (adj.find(v) == adj.end()) {
                return std::vector<int>();
            }

            std::vector<int> vertices;
            for (auto &i : adj[v]) {
                vertices.push_back(i.first);
            }
            return vertices;
        }

        void insert_vertex(int v) {
            if (_outgoing.find(v) == _outgoing.end()) {
                _outgoing[v] = std::unordered_map<int, bool>();
            }
            if (is_directed() && _incoming.find(v) == _incoming.end()) {
                _incoming[v] = std::unordered_map<int, bool>();
            }
        }

        void insert_edge(int u, int v, bool x = true) {
            if (is_directed()) {
                _outgoing[u][v] = x;
            } else {
                _outgoing[u][v] = x;
                _outgoing[v][u] = x;
            }
        }
};



void DFS(Graph& g, int root, std::unordered_map<int, bool>& discovered, int exc = -1, int y = -1) {
    std::vector<int> incident = g.incident_vertices(root);
    for (auto e : incident) {
        if (exc == e) {
            continue;
        }
        if (y == e) {
            discovered[e] = true;
            return;
        }
        if (discovered.find(e) == discovered.end()) {
            discovered[e] = true;
            DFS(g, e, discovered, exc);
        }
    }
}

bool is_dominator(Graph& g, int root, int x, int y) {
    std::unordered_map<int, bool> discovered;
    discovered[root] = true;
    if (root == x) {
        DFS(g, root, discovered, -1, y);
        if (discovered.find(y) != discovered.end()) {
            return true;
        }
    }
    if (x == y) {
        DFS(g, root, discovered, -1, y);
        if (discovered.find(y) != discovered.end()) {
            return true;
        } else {
            return false;
        }
    }
    DFS(g, root, discovered, -1, y);
    if (discovered.find(y) == discovered.end()) {
        return false;
    }
    discovered.clear();
    discovered[root] = true;
    DFS(g, root, discovered, x, y);
    if (discovered.find(y) == discovered.end()) {
        return true;
    }
    if (discovered.find(y) != discovered.end()) {
        return false;
    }
}


int main() {
    int n;
    std::cin >> n;
    for (int c = 0; c < n; c++) {
        Graph g(true);
        int nodes;
        std::cin >> nodes;
        for (int i = 0; i < nodes; i++) {
            std::vector<int> row(nodes);
            for (int j = 0; j < nodes; j++) {
                std::cin >> row[j];
            }
            for (int j = 0; j < nodes; j++) {
                int v = row[j];
                if (v == 1) {
                    g.insert_vertex(i);
                    g.insert_vertex(j);
                    g.insert_edge(i, j);
                }
            }
        }
        std::cout << "Case " << c+1 << ":" << std::endl;
        std::string s;
        for (int i = 0; i < nodes; i++) {
            int x = i;
            std::vector<char> answers;
            for (int j = 0; j < nodes; j++) {
                int y = j;
                answers.push_back(is_dominator(g, 0, x, y) ? 'Y' : 'N');
            }
            s = "|";
            for (auto ans : answers) {
                s += ans;
                s += "|";
            }
            std::cout << "+" << std::string(s.length() - 2, '-') << "+" << std::endl;
            std::cout << s << std::endl;
        }
        std::cout << "+" << std::string(s.length() - 2, '-') << "+" << std::endl;
    }
    return 0;
}
