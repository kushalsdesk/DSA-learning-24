#include <iostream>
using namespace std;
int sum_of_natural_numbers(int end_bound) {
  int sum = 0;
  for (int i = 0; i <= end_bound; i++) {
    sum += i;
  }
  return sum;
}

int main() {
  int sum, input;
  cout << "Enter the end point: ";
  cin >> input;
  sum = sum_of_natural_numbers(input);
  cout << "Result: " << sum << endl;
  return 0;
}
