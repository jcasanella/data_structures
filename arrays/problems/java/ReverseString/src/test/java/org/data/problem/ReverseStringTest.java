package org.data.problem;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class ReverseStringTest {

    @Test
    public void testEmptyString() {
        String cad = "";

        String actual = ReverseString.calculate(cad);
        assertEquals(actual, cad, "We expect an empty string");
    }

    @Test
    public void testNullString() {
        String actual = ReverseString.calculate(null);
        assertNull(actual, "We expect a null string");
    }

    @Test
    public void testStringOneChar() {
        String cad = "a";

        String actual = ReverseString.calculate(cad);
        assertEquals(actual, cad, "We expect a string equal to a");
    }

    @Test
    public void testStringOneWord() {
        String cad = "skate";

        String actual = ReverseString.calculate(cad);
        assertEquals(actual, "etaks", "We expect a string equal to etaks");
    }

    @Test
    public void testStringOneWordWithSpaces() {
        String cad = " skate ";

        String actual = ReverseString.calculate(cad);
        assertEquals(actual, " etaks ", "We expect a string equal to etaks with spaces");
    }

    @Test
    public void testStringMultipleWords() {
        String cad = "Inline skates is cool";

        String actual = ReverseString.calculate(cad);
        assertEquals(actual, "looc si setaks enilnI", "We expect a string equal to etaks with spaces");
    }
}
