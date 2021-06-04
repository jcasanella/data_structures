package org.data.algorithm;

public class BinarySearch {
    public static int of(final int[] numbers, int target) {
        if (numbers.length == 1) {
            if (numbers[0] == target) {
                return 0;
            } else {
                return -1;
            }
        }

        int low = 0;
        int high = numbers.length - 1;
        while (low <= high) {
            int m = (low+high)/2;
            if(numbers[m] <  target) {
                low = m + 1;
            } else if (numbers[m] > target) {
                high = m - 1;
            } else {
                return m;
            }
        }

        return -1;
    }
}
