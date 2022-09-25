package LeetCode.num_0215;

/**
 * @author ellison
 * @date 2022-09-25 01:30
 */
public class MainClass {

    class Solution {

        public int findKthLargest(int[] nums, int k) {
            k = nums.length - k; // index of k
            return partition(nums, 0, nums.length - 1, k);
        }

        private int partition(int[] nums, int low, int high, int k) {
            int l = low, r = high;
            int pivot = nums[l];
            while (l < r) {
                while (l < r && nums[r] >= pivot) {
                    r--;
                }
                nums[l] = nums[r];
                while (l < r && nums[l] <= pivot) {
                    l++;
                }
                nums[r] = nums[l];
            }
            nums[l] = pivot;

            if (l == k) {
                return nums[l];
            } else if (l > k) {
                return partition(nums, low, l - 1, k);
            } else {
                return partition(nums, l + 1, high, k);
            }
        }

    }

}
