//
// Created by ellisonzhao on 2023/2/5.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:
  string removeKdigits(string num, int k) {
    vector<char> stk;
    for (auto &digit : num) {
      while (k && stk.size() > 0 && stk.back() > digit) {
        stk.pop_back();
        k--;
      }
      stk.push_back(digit);
    }

    while (k && stk.size() > 0) {
      stk.pop_back();
      k--;
    }

    string ans(stk.begin(), stk.end());
    ans.erase(0, ans.find_first_not_of("0"));

    return ans == "" ? "0" : ans;
  }
};