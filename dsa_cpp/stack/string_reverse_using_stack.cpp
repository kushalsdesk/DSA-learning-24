#include <iostream>
#include <stack>
using namespace std;

int main() {

  stack<char> st;
  string input = "Kunal";

  for (char c : input) {
    st.push(c);
  }
  int size_of_string = input.length();

  for (int i = 0; i < size_of_string; i++) {
    input[i] = st.top();
    st.pop();
  }
  cout << "Output :: " << input << endl;
  return 0;
}
