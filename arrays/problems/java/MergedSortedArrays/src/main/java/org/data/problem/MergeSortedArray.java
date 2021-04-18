package org.data.problem;

public class MergeSortedArray {
    public static int[] run(int[] array1, int[] array2) {
        if (array1 == null && array2 != null) {
            return array2;
        } else if (array1 != null && array2 == null) {
            return array1;
        } else if (array1 == null && array2 == null) {
            return null;
        }

        int[] output = new int[array1.length + array2.length];

        int idx3 = 0;
        int idx2 = 0;
        for(int idx=0; idx<array1.length; idx++) {

            while (idx2<array2.length && array2[idx2] <= array1[idx]) {
                output[idx3++] = array2[idx2++];
            }

            output[idx3++] = array1[idx];
        }

        while (idx2 < array2.length) {
            output[idx3++] = array2[idx2++];
        }

        return output;
    }
}
