package org.data.problem;

import java.util.HashMap;

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/
public class TwoSum {

    // Complexity O^2
    public static int[] run(int[] nums, int target) {
        for (int idx=0; idx<nums.length-1; idx++) {
            int expected = target - nums[idx];

            for (int idx2 = idx+1; idx2<nums.length; idx2++) {
                if (nums[idx2] == expected) {
                    int[] output = new int[]{idx, idx2};
                    return output;
                }
            }
        }

        return null;
    }

    // Time complexity O(n)
    public static int[] runOptimized(int[] nums, int target) {
        HashMap<Integer, Integer> values = new HashMap<>();

        for (int idx=0; idx<nums.length; idx++) {
            if (values.containsKey(target - nums[idx])) {
                int[] output = new int[] {values.get(target - nums[idx]), idx};
                return output;
            }

            values.put(nums[idx], idx);
        }

        return null;
    }
}
