package org.data.structure;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

public class ArrayDynamicTest {

    private ArrayDynamic<Integer> ad;

    @BeforeEach
    public void initialize() {
        ad = new ArrayDynamic<>();
    }

    @Test
    @DisplayName("Testing an empty array")
    public void testEmptyArray() {
        assertEquals(0, ad.getLength(), "Length should be empty");
        assertEquals(1, ad.getCapacity(), "Capacity should be 1");
    }

    @Test
    @DisplayName("Testing push operation")
    public void testPush() {
        Integer value = ad.push(1);

        assertEquals(1, ad.getLength(), "Length should be 1");
        assertEquals(1, ad.getCapacity(), "Capacity should be 1");
        assertEquals(1, value, "Value should be 1");

        value = ad.push(2);
        assertEquals(2, ad.getLength(), "Length should be 2");
        assertEquals(2, ad.getCapacity(), "Capacity should be 2");
        assertEquals(2, value, "Value should be 2");

        value = ad.push(3);
        assertEquals(3, ad.getLength(), "Length should be 3");
        assertEquals(4, ad.getCapacity(), "Capacity should be 4");
        assertEquals(3, value, "Value should be 3");

        value = ad.push(4);
        assertEquals(4, ad.getLength(), "Length should be 4");
        assertEquals(4, ad.getCapacity(), "Capacity should be 4");
        assertEquals(4, value, "Value should be 4");

        ad.push(5);
        ad.push(6);
        value = ad.push(7);
        assertEquals(7, ad.getLength(), "Length should be 7");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        assertEquals(7, value, "Value should be 7");
    }

    @Test
    @DisplayName("Testing get operation")
    public void testGet() {
        Integer value = ad.get(-1);
        assertNull(value, "Object should be null when index not valid");

        Integer expected = ad.push(1);
        Integer actual = ad.get(0);
        assertNotNull(actual);
        assertEquals(expected, actual, "Should be the same value");

        ad.push(2);
        ad.push(34);
        ad.push(70);
        ad.push(45);

        actual = ad.get(4);
        assertEquals(45, actual, "Should be the same value");

        actual = ad.get(3);
        assertEquals(70, actual, "Should be the same value");

        actual = ad.get(2);
        assertEquals(34, actual, "Should be the same value");

        actual = ad.get(1);
        assertEquals(2, actual, "Should be the same value");

        value = ad.get(55);
        assertNull(value, "Object should be null when index not valid");
    }

    @Test
    @DisplayName("Testing pop operation")
    public void testPop() {
        Integer value = ad.pop();
        assertNull(value, "Object should be null when index not valid");

        Integer expected = ad.push(1);
        value = ad.pop();
        assertEquals(value, expected, "Should be the same value");
        assertEquals(0, ad.getLength(), "Length should be 0");
        assertEquals(1, ad.getCapacity(), "Capacity should be 1");

        expected = ad.push(2);
        value = ad.get(0);
        assertEquals(1, ad.getCapacity(), "Capacity should be 1");
        assertEquals(1, ad.getLength(), "Length should be 1");
        assertEquals(value, expected, "Should be the same value");

        ad.push(4);
        ad.push(8);
        ad.push(16);
        ad.push(32);
        ad.push(64);
        ad.push(128);
        assertEquals(7, ad.getLength(), "Length should be 7");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");

        List<Integer> values = new ArrayList<>(Arrays.asList(2, 4, 8, 16, 32, 64, 128));
        for(int idx = ad.getLength() - 1; idx >= 0; idx--) {
            value = ad.pop();
            Integer value2 = ad.get(idx);
            assertNull(value2, "Should be null");
            assertEquals(values.get(idx), value, "Should be the same value");

            String msg = String.format("Length should be %d", idx);
            assertEquals(idx, ad.getLength(), msg);
            assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        }

        value = ad.pop();
        assertNull(value, "Should be null");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        assertEquals(0, ad.getLength(), "Length should be 0");
    }

    @Test
    @DisplayName("Testing deleteAtIndex operation")
    public void testDeleteAtIndex() {
        Integer value = ad.deleteAtIndex(0);
        assertNull(value, "Should be null");

        ad.push(2);
        ad.push(4);
        ad.push(8);
        ad.push(16);
        ad.push(32);
        ad.push(64);
        ad.push(128);
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        assertEquals(7, ad.getLength(), "Length should be 7");

        value = ad.get(2);
        Integer value2 = ad.deleteAtIndex(2);
        assertEquals(value, value2, "Should be the same value");
        assertEquals(6, ad.getLength(), "Length should be 6");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        value = ad.get(2);
        assertEquals(16, value, "Should be the same value");

        value = ad.get(0);
        value2 = ad.deleteAtIndex(0);
        assertEquals(value, value2, "Should be the same value");
        assertEquals(5, ad.getLength(), "Length should be 5");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        value = ad.get(0);
        assertEquals(4, value, "Should be the same value");

        value = ad.get(3);
        value2 = ad.deleteAtIndex(3);
        assertEquals(value, value2, "Should be the same value");
        assertEquals(4, ad.getLength(), "Length should be 4");
        assertEquals(8, ad.getCapacity(), "Capacity should be 8");
        value = ad.get(3);
        assertEquals(128, value, "Should be the same value");

        ad.deleteAtIndex(3);
        assertEquals(3, ad.getLength(), "Length should be 3");
        ad.deleteAtIndex(2);
        assertEquals(2, ad.getLength(), "Length should be 2");
        ad.deleteAtIndex(1);
        assertEquals(1, ad.getLength(), "Length should be 1");
        ad.deleteAtIndex(0);
        assertEquals(0, ad.getLength(), "Length should be 0");

        value = ad.deleteAtIndex(50);
        assertNull(value, "Should be a null value");
    }

    @Test
    @DisplayName("Testing shiftingElements operation")
    public void testShiftingElements() {
        ad.shiftElements(0);
        assertEquals(0, ad.getLength(), "Length should be 0");
        Integer value = ad.get(0);
        assertNull(value, "Should be null");

        ad.push(1);
        ad.push(2);
        ad.push(3);
        ad.push(4);
        ad.push(5);
        ad.push(6);
        ad.push(7);
        assertEquals(7, ad.getLength(), "Length should be 7");

        ad.shiftElements(0);
        assertEquals(6, ad.getLength(), "Length should be 6");
        value = ad.get(0);
        assertEquals(2, value, "Should be the same value");

        ad.shiftElements(20);
        assertEquals(6, ad.getLength(), "Length should be 6");

        ad.shiftElements(3);
        assertEquals(5, ad.getLength(), "Length should be 5");
        value = ad.get(3);
        assertEquals(6, value, "Should be the same value");

        ad.shiftElements(3);
        assertEquals(4, ad.getLength(), "Length should be 4");
        value = ad.get(3);
        assertEquals(7, value, "Should be the same value");
    }
}
