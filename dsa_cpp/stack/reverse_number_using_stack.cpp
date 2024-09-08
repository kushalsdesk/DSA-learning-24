#include <cmath>
#include <iostream>
#include <stack>
using namespace std;

int main() {
  int output = 0;
  stack<int> st;
  int input = 1234;
  while (input != 0) {
    st.push(input % 10);
    input = input / 10;
  }
  int size = st.size();
  for (int i = 0; i < size; i++) {
    output = output + (st.top() * pow(10, i));
    st.pop();
  }

  cout << "Output :: " << output << endl;
  return 0;
}
