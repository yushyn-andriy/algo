#include <bits/stdc++.h>
using namespace std;

int gcd(int a, int b) {
    if (a % b == 0) {
        return b;
    }
    return gcd(b, a % b);
}

double estimate(vector<int> numbers) {
    int no_common_divisors = 0;
    int total_combinations = 0;
    for (int a = 0; a < numbers.size(); a++) {
        for (int b = a + 1; b < numbers.size(); b++) {
            total_combinations++;
            if (gcd(numbers[a], numbers[b]) == 1) {
                no_common_divisors++;
            }
        }
    }

    if (no_common_divisors == 0) {
        return -1;
    }

    return sqrt((6.0 * total_combinations) / no_common_divisors);
}

int main() {
    int n;
    while (cin >> n) {
        if (n <= 0 || n >= 50) {
            break;
        }

        vector<int> numbers(n);
        for (int i = 0; i < n; i++) {
            cin >> numbers[i];
        }

        cin.ignore();

        double pi = estimate(numbers);
        if (pi == -1) {
            cout << "No estimate for this data set." << endl;
        } else {
            cout << fixed << setprecision(6) << pi << endl;
        }
    }

    return 0;
}
