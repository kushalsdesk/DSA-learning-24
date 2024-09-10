#include <iostream>
#include <queue>
using namespace std;

int main() {

  int arr[] = {1, 2, 3, 4, 5};
  int size = sizeof(arr) / sizeof(arr[0]);
  queue<int> q;

  for (int i = 0; i < size; i++) {
    q.push(arr[i]);
  }

  while (!q.empty()) {
    cout << q.front() << " ";
    q.pop();
  }
  return 0;
}
