import java.util.Arrays;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class MainApp {

    public static boolean hasCommonElements(char[] array1, char[] array2) {
        // Create set
        Set<Character> characters = IntStream.range(0, array1.length)
                .mapToObj(posic -> array1[posic])
                .collect(Collectors.toSet());

        return IntStream.range(0, array2.length)
                .mapToObj(posic -> array2[posic])
                .anyMatch(characters::contains);
    }

    public static void main(String[] args) {
        // Given 2 arrays, create a function than let's user know (true, false) whether these 2 arrays contains
        // any common items
        // For example:
        // array1 = ['a', 'b', 'c', 'x']
        // array2 = ['z', 'y', 'i']
        // should return false

        // array1 = ['a', 'b', 'c', 'x']
        // array2 = ['z', 'y', 'x']
        // should return true

        char[] array1 = {'a', 'b', 'c', 'x'};
        char[] array2 = {'z', 'y', 'i'};

        boolean b1 = hasCommonElements(array1, array2);
        System.out.println(b1);

        char[] array3 = {'z', 'x', 'y'};
        boolean b2 = hasCommonElements(array1, array3);
        System.out.println(b2);
    }
}
