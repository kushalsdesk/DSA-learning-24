#include <iostream>
using namespace std;
#define row 3
#define col 3
int main() {

  int matrix_1[row][col] = {{1, 21, 3}, {42, 6, 6}, {7, 8, 19}};
  int matrix_2[row][col] = {{9, 6, 7}, {2, 5, 4}, {13, 20, 1}};
  int sum[row][col];
  int i, j;
  for (i = 0; i < row; i++) {
    for (j = 0; j < col; j++) {
      sum[i][j] = matrix_1[i][j] + matrix_2[i][j];
    }
  }
  cout << "The Result :" << "\n";
  for (i = 0; i < row; i++) {
    for (j = 0; j < col; j++) {
      cout << sum[i][j] << " ";
    }
    cout << "\n";
  }
  return 0;
}
