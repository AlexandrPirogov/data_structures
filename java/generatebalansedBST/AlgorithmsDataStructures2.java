import java.util.*;

public class AlgorithmsDataStructures2
{
    public static int[] GenerateBBSTArray(int[] a)
    {
        if (a.length == 0) {
          return a;
        } 
        Arrays.sort(a);
        int[] res = new int[a.length];
        generate(0, a.length-1, a, res, 0);
        return res;
    }

    private static void generate(int l, int r, int[] source, int[] res, int pos) {
      if (l < r){
        int mid = (l+r)/2;
        res[pos] = source[mid];
        generate(l, mid, source, res, pos*2+1);
        generate(mid+1, r, source, res, pos*2+2);
      }
      res[source.length-1] = source[source.length-1];
    }

}