#include <iostream>
using namespace std;

bool checker(int n) {
  int counter = 0;
  bool is_prime;
  for (int i = 1; i <= n; i++) {
    if (n % i == 0) {
      counter++;
    }
  }
  if (counter > 2) {
    is_prime = false;
  } else if (counter <= 2) {
    is_prime = true;
  }
  return is_prime;
}

int main() {
  int input;
  cout << "Enter Number for Check : ";
  cin >> input;
  if (checker(input)) {
    cout << "The Number is Prime" << endl;
  } else {
    cout << "The Number is Non Prime" << endl;
  }
  return 0;
}
