package org.data.algorithm;

// O(n^2)
public class BubbleSort {
    public static int[] of(final int[] numbers) {
        if (numbers.length <= 1)
            return numbers;

        int limit = numbers.length - 1;
        boolean sorted = false;

        while (!sorted) {
            for (int idx = 0; idx < limit; idx++) {
                if (numbers[idx] > numbers[idx + 1]) {
                    int temp = numbers[idx];
                    numbers[idx] = numbers[idx + 1];
                    numbers[idx + 1] = temp;
                }
            }

            limit--;

            if (limit == 0) {
                sorted = true;
            }
        }

        return numbers;
    }
}
