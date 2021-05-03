package org.data.problem;

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]

*/
public class MoveZeroes {
    // O(n^2)
    public static int[] doFunc(int[] nums) {
        for (int idx=0; idx < nums.length; idx++) {
            if (nums[idx] == 0) {
                int idx2 = idx + 1;
                while (idx2 < nums.length) {
                    if (nums[idx2] != 0) {
                        int aux = nums[idx2];
                        nums[idx2] = nums[idx];
                        nums[idx] = aux;
                        break;
                    } else {
                        idx2++;
                    }
                }
            }
        }

        return nums;
    }

    // O(n)
    public static int[] doFuncOptimized(int[] nums) {
        int[] aux = new int[nums.length];

        int nonZero=0;
        for (int idx=0; idx<nums.length; idx++) {
            if (nums[idx] != 0) {
                aux[nonZero++] = nums[idx];
            }
        }

        int zeros = nums.length - nonZero;
        for (int idx2 = nonZero; idx2<zeros; idx2++) {
            aux[idx2] = 0;
        }

        for (int idx=0; idx<nums.length; idx++) {
            nums[idx] = aux[idx];
        }

        return nums;
    }
}