#include <iostream>

int factorial (int n){
    /*computes factorial of n*/
    if (n==1) {
        return n;
    }
    return n * factorial(n-1);
}

int main() {

    std:: cout << factorial(8);

}
