#include <iostream>
#include <vector>
using namespace std;

// Function to calculate factorial
int factorial(int n) {
    int result = 1;
    for (int i = 2; i <= n; ++i) {
        result *= i;
    }
    return result;
}

// Function to compute the rank of a given permutation
int rankPermutation(int n, int* p) {
    int rank = 0;

    // Use a vector to mark the availability of numbers from 1 to n, initialized to false
    vector<bool> seen(n + 1, false);

    // Iterate through each element in the permutation
    for (int i = 0; i < n; ++i) {
        // Count how many numbers less than p[i] are still available
        int count = 0;
        for (int j = 1; j < p[i]; ++j) {
            if (!seen[j]) {
                count++;
            }
        }

        // Update the rank using the count of smaller numbers
        rank += count * factorial(n - i - 1);

        // Mark the current element as seen
        seen[p[i]] = true;
    }

    // Add 1 to convert to 1-based rank (since rank starts from 1)
    return rank + 1;
}

int main() {
    int n;
    cout << "Enter the value of n: ";
    cin >> n;

    // Create an array of size n
    int* p = new int[n];

    // Read values into the array
    cout << "Enter " << n << " integers (as a permutation of 1 to " << n << "): ";
    for (int i = 0; i < n; ++i) {
        cin >> p[i];
    }

    // Compute the rank of the permutation
    int rank = rankPermutation(n, p);
    cout << "The rank of the permutation is: " << rank << endl;

    // Free the allocated memory
    delete[] p;

    return 0;
}
