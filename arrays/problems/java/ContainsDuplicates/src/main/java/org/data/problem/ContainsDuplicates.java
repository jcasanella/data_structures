package org.data.problem;

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
//
// Example 1:
// Input: nums = [1,2,3,1]
// Output: true

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false
//
// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

import java.util.HashSet;
import java.util.Set;

// Cost O(n)
public class ContainsDuplicates {
    public static boolean doOper(int[] nums) {
        Set<Integer> values = new HashSet<>();

        for (int idx=0; idx<nums.length; idx++) {
            if (values.contains(nums[idx]))
                return true;

            values.add(nums[idx]);
        }

        return false;
    }
}
