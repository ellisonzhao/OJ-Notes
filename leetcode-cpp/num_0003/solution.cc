//
// Created by ellisonzhao on 2022/9/24.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:

  int lengthOfLongestSubstring(string s) {
    vector<int> dict(256, -1);
    int maxLen = 0, start = -1;
    for (int i = 0; i != s.length(); i++) {
      // 窗口内出现重复字符
      if (dict[s[i]] > start)
        start = dict[s[i]];
      // 遍历过程中更新每个字符最近出现的位置
      dict[s[i]] = i;
      // 更新长度
      maxLen = max(maxLen, i - start);
    }
    return maxLen;
  }
};

//  int lengthOfLongestSubstring(string s) {
//    unordered_map<char, int> window;
//    int left = 0, right = 0;
//    int res = 0, len = 0;
//    while (right < s.length()) {
//      char rc = s[right];
//      // 当且仅当窗口内出现重复字符时，需要更新 left
//      if (window.count(rc) > 0 && window[rc] >= left) {
//        left = window[rc] + 1;
//        // 更新窗口大小
//        len = right - left;
//      }
//      // 更新该重复字符最新出现的位置
//      window[rc] = right;
//      ++len;
//      ++right;
//
//      res = max(res, len);
//    }
//    return res;
//  }