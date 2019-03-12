package LeetCode.num_889;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

class Solution {
    public TreeNode constructFromPrePost(int[] preorder, int[] postorder) {
        Deque<TreeNode> dq = new ArrayDeque<>();
        dq.offer(new TreeNode(preorder[0]));
        for (int i = 1, j = 0; i < preorder.length; ++i) {
            TreeNode node = new TreeNode(preorder[i]);
            while (dq.getLast().val == postorder[j]) {
                dq.pollLast();
                if (dq.isEmpty()) {
                    dq.offer(node);
                }
                j++;
            }
            if (dq.getLast().left == null) dq.getLast().left = node;
            else dq.getLast().right = node;
            dq.offer(node);
        }
        return dq.getFirst();
    }
}

public class MainClass {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
            return new int[0];
        }

        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static String treeNodeToString(TreeNode root) {
        if (root == null) {
            return "[]";
        }

        String output = "";
        Queue<TreeNode> nodeQueue = new LinkedList<>();
        nodeQueue.add(root);
        while (!nodeQueue.isEmpty()) {
            TreeNode node = nodeQueue.remove();

            if (node == null) {
                output += "null, ";
                continue;
            }

            output += String.valueOf(node.val) + ", ";
            nodeQueue.add(node.left);
            nodeQueue.add(node.right);
        }
        return "[" + output.substring(0, output.length() - 2) + "]";
    }

    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int[] pre = stringToIntegerArray(line);
            line = in.readLine();
            int[] post = stringToIntegerArray(line);

            TreeNode ret = new Solution().constructFromPrePost(pre, post);

            String out = treeNodeToString(ret);

            System.out.print(out);
        }
    }
}
