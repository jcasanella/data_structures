package org.data.problem;

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
//
// Example1
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
//
// Example2
// Input: nums = [1]
// Output: 1
//
// Example3
// Input: nums = [5,4,-1,7,8]
// Output: 23
//
// Cost O (N)
public class MaximumSubArray {
    public static int doOperac(int[] nums) {
        int maxSoFar = Integer.MIN_VALUE;
        int maxEndingHere = 0;

        for (int idx = 0; idx < nums.length; idx++) {
            maxEndingHere += nums[idx];
            if (maxSoFar < maxEndingHere) {
                maxSoFar = maxEndingHere;
            }

            if (maxEndingHere < 0) {
                maxEndingHere = 0;
            }
        }

        return maxSoFar;
    }
}
