package AcWing.nums_65;

class Solution {

    /*
    * 假设目标值在闭区间[l, r]中，
    * 每次将区间长度缩小一半，
    * 当l = r时，我们就找到了目标值。
    * */
    public int binarySearch(int[] nums, int k) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int mid = l + (r - l) / 2; // 取中间位置，防止越界
            if (nums[mid] >= k) // check(mid)条件，根据实际选择
                r = mid; // 区间[l, r]的更新操作是r = mid; l = mid + 1;
            else
                l = mid + 1;
        }

        return l;
    }

    public int inversePairs(int[] nums) {
        if (nums.length == 0 || nums == null)
            return 0;
        int[] copy = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            copy[i] = nums[i];
        }
        int ans = merge(nums, copy, 0, nums.length - 1);
        return ans;
    }

    private int merge(int[] nums, int[] copy, int start, int end) {
        if (start == end) {
            copy[start] = nums[start];
            return 0;
        }

        int len = (end - start) / 2;
        int left = merge(copy, nums, start, start + len);
        int right = merge(copy, nums, start + len + 1, end);

        int i = start + len, j = end, indexCopy = end, count = 0;

        while (i >= start && j >= start + len + 1) {
            if (nums[i] > nums[j]) {
                copy[indexCopy--] = nums[i--];
                count += j - (start + len);
            } else {
                copy[indexCopy--] = nums[j--];
            }
        }

        while (i >= start)
            copy[indexCopy--] = nums[i--];
        while (j >= start + len + 1)
            copy[indexCopy--] = nums[j--];

        return left + right + count;
    }
}

public class MainClass {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 3, 4};
        System.out.println(new Solution().binarySearch(nums, 3));
    }
}
