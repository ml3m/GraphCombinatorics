def is_permutation(n, p):
    # Create a boolean list to mark seen numbers
    seen = [False] * n
    # Iterate through the array p
    for val in p:
        # Check if the value is out of range or already seen
        if val < 1 or val > n or seen[val - 1]:
            print("not a permutation")
            return
        # Mark the number as seen
        seen[val - 1] = True

    # If all numbers are marked as seen, it's a permutation
    print("is a permutation")


# Main program to read input and check for permutation
if __name__ == "__main__":
    # Read n (the number of elements)
    n = int(input("Enter the number of elements (n): "))
    # Read the array p
    p = list(map(int, input(f"Enter {n} integers separated by spaces: ").split()))
    # Call the is_permutation function
    is_permutation(n, p)
