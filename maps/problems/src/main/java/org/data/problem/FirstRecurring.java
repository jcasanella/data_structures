package org.data.problem;



import java.util.*;

// Find the first duplicated element
public class FirstRecurring {
    public static Integer run(final int[] numbers) {
        Set<Integer> values = new HashSet<>();
        for(int idx = 0; idx < numbers.length; idx++) {
            if (values.contains(numbers[idx])) {
                return numbers[idx];
            }

            values.add(numbers[idx]);
        }

        return null;
    }
}
