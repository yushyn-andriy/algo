#include <stdio.h>
#include <math.h>

int parent(int n) {
    return n / 2;
}

int solve(int N, int i, int j) {
    int total_participants = (int)pow(2.0, (double)N);
    int last_level_nodes_count = total_participants / 2;

    int curr_node1 = (i + 1) / 2 + (last_level_nodes_count - 1);
    int curr_node2 = (j + 1) / 2 + (last_level_nodes_count - 1);

    int round = 1;
    while (curr_node1 != curr_node2) {
        curr_node1 = parent(curr_node1);
        curr_node2 = parent(curr_node2);
        round++;
    }

    return round;
}

int main() {
    int N, i, j;
    while (scanf("%d %d %d", &N, &i, &j) != EOF) {
        printf("%d\n", solve(N, i, j));
    }
    return 0;
}
