#include <iostream>
#include <algorithm>

using namespace std;

void prevPermutation(int n, int* p) {
    // Step 1: Find the first element from the end that is greater than its next element
    int i = n - 2;
    while (i >= 0 && p[i] <= p[i + 1]) {
        i--;
    }

    // If no such element exists, the permutation is the first one, so reverse it
    if (i < 0) {
        reverse(p, p + n);
    } else {
        // Step 2: Find the rightmost element that is smaller than p[i]
        int j = n - 1;
        while (p[j] >= p[i]) {
            j--;
        }

        // Step 3: Swap p[i] and p[j]
        swap(p[i], p[j]);

        // Step 4: Reverse the sequence from p[i + 1] to the end to get the previous permutation
        reverse(p + i + 1, p + n);
    }

    // Display the result
    cout << "Previous permutation is: ";
    for (int k = 0; k < n; ++k) {
        cout << p[k] << " ";
    }
    cout << endl;
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

    // Find and display the next permutation
    prevPermutation(n, p);

    // Free the allocated memory
    delete[] p;

    return 0;
}
