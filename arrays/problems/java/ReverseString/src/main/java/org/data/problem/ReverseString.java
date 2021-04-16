package org.data.problem;

// Create a function that reverses a string
public class ReverseString {

    public static String calculate(String cad) {
        if (cad == null || cad.isEmpty()) {
            return cad;
        }

        char[] characs = cad.toCharArray();
        char[] output = new char[cad.length()];

        int idx2 = 0;
        for(int idx = cad.length() - 1; idx >= 0; idx--) {
            output[idx2++] = characs[idx];
        }

        return String.valueOf(output);
    }
}
