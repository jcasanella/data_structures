package org.data.algorithm;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BubbleSortTest {
    @DisplayName("An empty array should return an empty array")
    @Test
    public void testEmptyArray() {
        int[] actual = {};
        int[] expected = BubbleSort.of(actual);

        assertTrue(expected.length == 0);
    }

    @DisplayName("An array of one element should return the same array")
    @Test
    public void testSingleElement() {
        int[] actual = { 5 };
        int[] expected = BubbleSort.of(actual);

        assertArrayEquals(expected, actual);
    }

    @DisplayName("A sorted array should return the same array")
    @Test
    public void testSortedArray() {
        int[] actual = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        int[] expected = BubbleSort.of(actual);

        assertArrayEquals(expected, actual);
    }

    @DisplayName("An array with positive numbers should be sorted")
    @Test
    public void testUnsortedArrayPositiveNumbers() {
        int[] source = { 2, 1, 5, 3, 1, 5, 0, 12, 9 };
        int[] expected = BubbleSort.of(source);
        int[] actual = { 0, 1, 1, 2, 3, 5, 5, 9, 12 };

        assertArrayEquals(expected, actual);
    }

    @DisplayName("An array with negative numbers should be sorted")
    @Test
    public void testUnsortedArrayNegativeNumbers() {
        int[] source = { -2, -12, -17, -4, -19, -3, -1 };
        int[] expected = BubbleSort.of(source);
        int[] actual = { -19, -17, -12, -4, -3, -2, -1 };

        assertArrayEquals(expected, actual);
    }

    @DisplayName("An array with positive and negative numbers should be sorted")
    @Test
    public void testUnsortedArrayNumbers() {
        int[] source = { -2, 2, 1, 5, -12, -17, 3, 1, 5, -4, -19, -3, 0, 12, 9, -1 };
        int[] expected = BubbleSort.of(source);
        int[] actual = { -19, -17, -12, -4, -3, -2, -1, 0, 1, 1, 2, 3, 5, 5, 9, 12 };

        assertArrayEquals(expected, actual);
    }
}