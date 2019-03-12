package LeetCode.Test_01;

import java.io.IOException;
import java.util.*;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (char c : s.toCharArray()) {
            if (c == '(' || c == '[' || c == '{')
                stack.push(c);
            else {
                if (stack.isEmpty())
                    return false;
                else {
                    if (c == ')' && stack.peek() == '(')
                        stack.pop();
                    else if (c == ']' && stack.peek() == '[')
                        stack.pop();
                    else if (c == '}' && stack.peek() == '{')
                        stack.pop();
                    else
                        stack.push(c);
                }
            }

        }
        if (stack.empty())
            return true;
        else
            return false;
    }

    public int[][] reconstructQueue(int[][] people) {
//        Arrays.sort(people, (o1, o2) -> o1[0] != o2[0] ? o2[0] - o1[0] : o1[1] - o2[1]);
        List<int[]> res = new LinkedList<>();
        for (int[] cur : people) {
            res.add(cur[1], cur);
        }
        Arrays.sort(people, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[0] == o2[0]) {
                    return o1[1] - o2[1];
                } else {
                    return o2[0] - o1[0];
                }
            }
        });
        return res.toArray(new int[people.length][]);
    }

    public List<Integer> topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        List<List<Integer>> frequency = new ArrayList<>();
        Set<Map.Entry<Integer, Integer>> entries = map.entrySet();
        Iterator<Map.Entry<Integer, Integer>> it = entries.iterator();
        while (it.hasNext()) {
            Map.Entry<Integer, Integer> next = it.next();
            List<Integer> temp = new ArrayList<>();
            temp.add(next.getKey());
            temp.add(next.getValue());
            frequency.add(temp);
        }

        Collections.sort(frequency, (o1, o2) -> o2.get(1) - o1.get(1));

        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < k; ++i) {
            ans.add(frequency.get(i).get(0));
        }
        return ans;
    }

}

public class MainClass {

    public static void main(String[] args) throws IOException {
//        String s = "(])";
//        System.out.println(new Solution().isValid(s));

//        int[][] people = {{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}};
//        System.out.println(new Solution().reconstructQueue(people));
        Scanner in = new Scanner(System.in);
        String str = in.next();
    }
}
