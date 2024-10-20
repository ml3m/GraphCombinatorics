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

// Function to find the k-th permutation of {1, 2, ..., n}
void permWithRank(int k, int n) {
    // Decrement k to convert to 0-based index
    k -= 1;

    // List of available numbers from 1 to n
    vector<int> numbers;
    for (int i = 1; i <= n; ++i) {
        numbers.push_back(i);
    }

    // Vector to store the result permutation
    vector<int> result;

    // Construct the permutation
    for (int i = n; i >= 1; --i) {
        // Find the index of the next number to pick
        int factorialValue = factorial(i - 1);
        int index = k / factorialValue;
        
        // Select the number at the found index
        result.push_back(numbers[index]);

        // Remove the selected number from the list
        numbers.erase(numbers.begin() + index);

        // Update k for the next position
        k %= factorialValue;
    }

    // Print the result permutation
// Print the result permutation using a traditional for loop
    cout << "Permutation with rank " << (k + 1) << " is: ";
    for (size_t i = 0; i < result.size(); ++i) {
        cout << result[i] << " ";
    }
    cout << endl;
}

int main() {
    int k, n;
    cout << "Enter the value of n: ";
    cin >> n;
    cout << "Enter the rank k: ";
    cin >> k;

    // Find and print the permutation with rank k
    permWithRank(k, n);

    return 0;
}
