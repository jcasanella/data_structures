package org.data.problem;

import org.junit.jupiter.api.Test;
import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class MergedSortedArrayTest {

    @Test
    public void testBothArraysNull() {
        int[] out = MergeSortedArray.run(null, null);

        assertNull(out);
    }

    @Test
    public void testFirstArrayNull() {
        int[] array2 = {1, 2, 3};

        int[] out = MergeSortedArray.run(null, array2);

        assertTrue(Arrays.equals(out, array2));
    }

    @Test
    public void testSecondArrayNull() {
        int[] array1 = {1, 2, 3};

        int[] out = MergeSortedArray.run(array1, null);

        assertTrue(Arrays.equals(out, array1));
    }

    @Test
    public void testArraysOneElement() {
        int[] array1 = {200};
        int[] array2 = {50};

        int[] actual = {50, 200};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testFirstArrayMinor() {
        int[] array1 = {1, 2, 3, 4, 5};
        int[] array2 = {6, 7, 8};

        int[] actual = {1,2, 3, 4, 5, 6, 7, 8};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testFirstArrayBigger() {
        int[] array1 = {6, 7, 8};
        int[] array2 = {1, 2, 3, 4, 5};

        int[] actual = {1,2, 3, 4, 5, 6, 7, 8};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMergeIntoArray1() {
        int[] array1 = {1, 2, 4, 5, 6};
        int[] array2 = {3};

        int[] actual = {1,2, 3, 4, 5, 6};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMergeIntoArray1B() {
        int[] array1 = {1, 2, 4, 5, 6};
        int[] array2 = {3, 7};

        int[] actual = {1,2, 3, 4, 5, 6, 7};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMergeIntoArray2() {
        int[] array1 = {3};
        int[] array2 = {1, 2, 4, 5, 6};

        int[] actual = {1,2, 3, 4, 5, 6};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMergeIntoArray2B() {
        int[] array1 = {3, 7};
        int[] array2 = {1, 2, 4, 5, 6};

        int[] actual = {1,2, 3, 4, 5, 6, 7};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMerge() {
        int[] array1 = {3, 7, 21, 100, 150};
        int[] array2 = {4, 23, 44, 130, 600};

        int[] actual = {3, 4, 7, 21, 23, 44, 100, 130, 150, 600};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }

    @Test
    public void testMergeWithDuplicates() {
        int[] array1 = {3, 7, 21, 23, 100, 150, 600};
        int[] array2 = {3, 4, 23, 44, 130, 600};

        int[] actual = {3, 3, 4, 7, 21, 23, 23, 44, 100, 130, 150, 600, 600};
        int[] out = MergeSortedArray.run(array1, array2);

        assertTrue(Arrays.equals(out, actual));
    }
}
