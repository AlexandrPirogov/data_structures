import static org.junit.Assert.assertArrayEquals;
import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class BBST_test {
      @Test
    public void TestGenerateBSTEmpty() {
        int[] a = new int[]{};
        // 1,2,3,4,5,6,7,8,9
        int[] expected = new int[]{};
        int[] sorted = AlgorithmsDataStructures2.GenerateBBSTArray(a);

        assertArrayEquals(sorted, expected);
    }

          @Test
    public void TestGenerateBSTSingle() {
        int[] a = new int[]{1};
        // 1,2,3,4,5,6,7,8,9
        int[] expected = new int[]{1};
        int[] sorted = AlgorithmsDataStructures2.GenerateBBSTArray(a);

        assertArrayEquals(sorted, expected);
    }

    @Test
    public void TestGenerateBST() {
        int[] a = new int[]{9,2,6,3,1,4,7,8,5};
        // 1,2,3,4,5,6,7,8,9
        int[] expected = new int[]{5,3,7,2,4,6,8,1,9};
        int[] sorted = AlgorithmsDataStructures2.GenerateBBSTArray(a);

        assertArrayEquals(sorted, expected);
    }

    @Test
    public void TestGenerateBST_1() {
        int[] a = new int[]{3,2,1};
        //1,2,3
        int[] expected = new int[]{2,1,3};
        int[] sorted = AlgorithmsDataStructures2.GenerateBBSTArray(a);

        assertArrayEquals(sorted, expected);
    }

       @Test
    public void TestGenerateBST_2() {
        int[] a = new int[]{5,4,3,2,1};
        //1,2,3,4,5
        int[] expected = new int[]{3,2,4,1,5};
        int[] sorted = AlgorithmsDataStructures2.GenerateBBSTArray(a);

        assertArrayEquals(sorted, expected);
    }
}
