package org.data.algorithm;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.is;
import static org.hamcrest.Matchers.equalTo;

class BinarySearchTest {
    @DisplayName("Search in a empty array")
    @Test
    public void testSearchInEmptyArray() {
        int[] values = {};
        int actual = BinarySearch.of(values, 5);
        assertThat(actual, is(equalTo(-1)));
    }

    @DisplayName("Should not find element in an array of 1 element")
    @Test
    public void testSearchNotFindInASingleArray() {
        int[] values = { 5 };
        int actual = BinarySearch.of(values, 4);
        assertThat(actual, is(equalTo(-1)));
    }

    @DisplayName("Should find element in an array of 1 element")
    @Test
    public void testSearchFindInASingleArray() {
        int[] values = { 4 };
        int actual = BinarySearch.of(values, 4);
        assertThat(actual, is(equalTo(0)));
    }

    @DisplayName("Should find element in an array of 2 elements")
    @Test
    public void testSearchFindInArray() {
        int[] values = { 4, 6 };
        int actual = BinarySearch.of(values, 4);
        assertThat(actual, is(equalTo(0)));
    }

    @DisplayName("Should not find element in an array of 2 elements")
    @Test
    public void testSearchNotFindInArray() {
        int[] values = { 4, 6 };
        int actual = BinarySearch.of(values, 7);
        assertThat(actual, is(equalTo(-1)));
    }

    @DisplayName("Should find 2nd element in an array of 2 elements")
    @Test
    public void testSearch2FindInArray() {
        int[] values = { 4, 6 };
        int actual = BinarySearch.of(values, 6);
        assertThat(actual, is(equalTo(1)));
    }

    @DisplayName("Should find element in a big array")
    @Test
    public void testSearchFindInBigArray() {
        int[] values = { 4, 6, 8, 12, 20, 40, 70, 90 };
        int actual = BinarySearch.of(values, 70);
        assertThat(actual, is(equalTo(6)));
    }

}