import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

import java.util.Arrays;
import java.util.Collections;

import org.junit.*;

public class Heap_test {
    

    @Test
    public void CreateHeapEmpty(){
        int depth = 10;
        Heap h = new Heap();
        int[] given = new int[10];
        h.MakeHeap(given, depth);

        assertTrue(h.HeapArray.length == Math.pow(2, depth+1)-1);
    }


    @Test
    public void CreateHeapEmptyAndAdd(){
        Heap h = new Heap();
        int[] given = new int[0];
        h.MakeHeap(given, 5);

        int expected = 10; 
        h.Add(expected);
        int actual = h.GetMax();
        assertEquals(expected, actual); 
    }

     @Test
    public void CreateHeapEmptyAndAddNewMax(){
        Heap h = new Heap();
        int[] given = new int[0];
        h.MakeHeap(given, 5);
        h.Add(5);

        int expected = 10; 
        h.Add(10);
        int actual = h.GetMax();
        assertEquals(expected, actual); 
    }

     @Test
    public void GetMaxFromEmpty(){
        Heap h = new Heap();
        int[] given = new int[0];
        h.MakeHeap(given, 0);
        int res = h.GetMax();
        assertTrue(-1 == res);
    }

    @Test
    public void AddInFull() {
        Heap h = new Heap();
        int[] given = new int[0];
        h.MakeHeap(given, 5);
        for (int expected = 0; expected < (int)Math.pow(2, 6)-1; expected++) {
            h.Add(expected);
        }   

        boolean actual = h.Add(1000);
        assertFalse(actual);
    }

    @Test
    public void CreateHeapEmptyAndAddMultiple(){
        Heap h = new Heap();
        int[] given = new int[0];
        h.MakeHeap(given, 5);
        for (int expected = 0; expected < (int)Math.pow(2, 6)-1; expected++) {
            h.Add(expected);
            int actual = h.GetMax();
            assertEquals(expected, actual);
        }   

        assertTrue(h.GetMax() == -1);
    }

 @Test
    public void TestWithSortedArray(){
        Heap h = new Heap();
        int[] given = new int[]{1,2,3,4,5,6,7,8,9,10};
        int[] e = new int[]{10,9,8,7,6,5,4,3,2,1};
        h.MakeHeap(given, 5);

    
        for (Integer expected : e) {
            int actual = h.GetMax();
            assertTrue(expected == actual);
        }
    assertTrue(h.GetMax() == -1);
    }


    @Test
    public void GetMaxFromFull(){
        Heap h = new Heap();
        int[] given = new int[]{7,6,2, 0, 9};
        h.MakeHeap(given, 1);
        for( Integer num : given){
             int res = h.GetMax();
            assertFalse(5 == res);

        }
    }


}
