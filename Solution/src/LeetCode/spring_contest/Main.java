package LeetCode.spring_contest;

import java.util.HashSet;
import java.util.Set;

/**
 * @author charlsonz
 * @date 2019-04-14 15:16
 */
public class Main {

    public int shortestPath(String source, String target) {
        char[] src = source.toCharArray();
        char[] tar = target.toCharArray();
        Set<Character> set = new HashSet<>();
        for (char c : src)
            set.add(c);
        int res = 0;
        int i = 0, j = 0;
        while (j < tar.length) {
            if (!set.contains(tar[j])) {
                return -1;
            }

            if (src[i] == tar[j]) {
                ++j;
            }

            if (++i == src.length) {
                ++res;
                i = 0;
            }
        }
        return res;
    }

    public static void main(String[] args) {
        String source = "abc";
        String target = "abdbc";
        System.out.println(new Main().shortestPath(source, target));
    }
}
