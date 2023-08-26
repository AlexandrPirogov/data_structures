import java.util.*;

class Heap
{
    public int [] HeapArray; // хранит неотрицательные числа-ключи
    
    public Heap() { HeapArray = null; }
    private int cursor;
    public void MakeHeap(int[] a, int depth)
    {
        int newlen = (int)Math.pow(2, depth+1)-1;
        HeapArray = new int[newlen];
        cursor = 0;
        for (int i = 0; i < a.length; i++) {
            Add(a[i]);
        }

        for (int i = cursor; i < HeapArray.length; i++) {
            HeapArray[i] = -1;
        }
    }
    
    public int GetMax()
    {
      if (HeapArray[0] < 0 || cursor <= 0) {
        return -1;
      }

      int res = HeapArray[0]; 
      cursor--;
      HeapArray[0] = HeapArray[cursor];
      HeapArray[cursor] = -1;
      bubbleDown();


      return res;
    }

    public boolean Add(int key)
    {
        if (cursor > HeapArray.length-1) {
            return false;
        }

        HeapArray[cursor] = key;
        bubbleUp();
        cursor++;
        return true; 
    }

    private void bubbleDown() {
        int curr = 0;
        int leftMax = 1;
        int rightMax = 2;
        while (curr < cursor && (HeapArray[curr] < HeapArray[leftMax] || HeapArray[curr] < HeapArray[rightMax])) {
            int max = HeapArray[leftMax];
            int maxI = leftMax;
            if (max < HeapArray[rightMax]) {
                max = HeapArray[rightMax];
                maxI = rightMax;
            }

            int tmp = HeapArray[curr];
            HeapArray[curr] = HeapArray[maxI];
            HeapArray[maxI] = tmp;

            curr = maxI;
            leftMax = 2 * curr+1;
            rightMax = 2 * curr+2;
        }
    }

    private void bubbleUp() {
        int curr = cursor;
        while (curr > 0 && (HeapArray[curr] > HeapArray[(curr-1)/2])) {
            int tmp = HeapArray[curr];
            HeapArray[curr] = HeapArray[(curr-1)/2];
            HeapArray[(curr-1)/2] = tmp;
            curr = (curr-1)/2;
        }
    }
  
}