/*void isPermutation(int n, int* p) which takes as input arguments a
positive integer n and a (pointer to) an array of n integers, and prints the
message
• "is a permutation" if hp[0], . . . , p[n − 1]i is a permutation of the
numbers 1, . . . , n.
• "not a permutation" otherwise
*/
#include <iostream>
#include <vector>
using namespace std;

void isPermutation(int n, int* p) {
    // Create a vector to track the presence of numbers from 1 to n
    vector<bool> seen(n, false);
    
    // Iterate through the array
    for (int i = 0; i < n; ++i) {
        // If the number is out of range or already seen, it's not a permutation
        if (p[i] < 1 || p[i] > n || seen[p[i] - 1]) {
            cout << "not a permutation" << endl;
            return;
        }
        // Mark the number as seen
        seen[p[i] - 1] = true;
    }
    
    // If all numbers from 1 to n are seen exactly once, it's a permutation
    cout << "is a permutation" << endl;
}

int main() {
    int n;
    cout << "Enter the value of n: ";
    cin >> n;

    // Create an array of size n
    int* p = new int[n];

    // Read values into the array
    cout << "Enter " << n << " integers: ";
    for (int i = 0; i < n; ++i) {
        cin >> p[i];
    }

    // Call the function to check if it's a permutation
    isPermutation(n, p);

    // Free the allocated memory
    delete[] p;

    return 0;
}
