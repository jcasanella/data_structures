package org.data.problem;

import org.junit.jupiter.api.Test;

import static org.hamcrest.CoreMatchers.*;
import static org.hamcrest.MatcherAssert.assertThat;

public class FirstRecurringTest {

    @Test
    public void shouldReturnNullEmptyArray() {
        int[] numbers = {};
        Integer res = FirstRecurring.run(numbers);

        assertThat(res, is(nullValue()));
    }

    @Test
    public void shouldReturnNullNotRepeatedElement() {
        int[] numbers = {2, 3, 4, 5};
        Integer res = FirstRecurring.run(numbers);

        assertThat(res, is(nullValue()));
    }

    @Test
    public void shouldReturnElement() {
        int[] numbers = {2, 1, 1, 2, 3, 5, 1, 2, 4};
        Integer res = FirstRecurring.run(numbers);

        assertThat(res, is(equalTo(1)));
    }

    @Test
    public void shouldReturnElement2() {
        int[] numbers = {2, 5, 1, 2, 3, 5, 1, 2, 4};
        Integer res = FirstRecurring.run(numbers);

        assertThat(res, is(equalTo(2)));
    }
}
