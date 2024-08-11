#include <iostream>
using namespace std;

int main() {
  int matrix[3][3] = {{1, 2, 3}, {1, 2, 3}, {1, 2, 3}};
  int x;
  cout << "Press 1 for Primary Diagonal, Press 2 For Secondary Diagonal:";
  cin >> x;
  if (x == 1) {
    for (int i = 0; i < 3; i++) {
      for (int j = 0; j < 3; j++) {

        if (i == j) {

          cout << matrix[i][j] << " ";

        } else {
          cout << " " << " ";
        }
      }
      cout << "\n";
    }
  } else if (x == 2) {
    for (int i = 2; i >= 0; i--) {
      for (int j = 0; j < 3; j++) {
        if (i == j) {
          cout << matrix[i][j] << " ";
        } else {
          cout << " " << " ";
        }
      }
      cout << "\n";
    }

  } else {
    cout << "Invalid Input >< ";
  }
  return 0;
}
