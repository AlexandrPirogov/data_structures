import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertTrue;

import org.junit.Test;

public class BalancedBST_test {
    
    @Test
    public void TestGenerate() {
        BalancedBST sut = new BalancedBST();
        int[] a = new int[]{1,2,3,4,5};
        sut.GenerateTree(a);

        assertEquals(0, sut.Root.Level);
    }

    @Test
    public void TestGenerateEmpty() {
        BalancedBST sut = new BalancedBST();
        int[] a = new int[]{};
        sut.GenerateTree(a);

        assertNull(sut.Root);
    }

    @Test
    public void TestGenerateSingle() {
        BalancedBST sut = new BalancedBST();
        int[] a = new int[]{1};
        sut.GenerateTree(a);

        BSTNode tmp = sut.Root;
        int level = 0;
        while(tmp.LeftChild != null) {
            assertEquals(level, tmp.Level);
            level++;
            tmp = tmp.LeftChild;
        }
    }

    @Test
    public void TestGenerate_1() {
        BalancedBST sut = new BalancedBST();
        int[] a = new int[]{1,2,3,4,5,6,7,8,9,10};
        sut.GenerateTree(a);

        BSTNode tmp = sut.Root;
        int level = 0;
        while(tmp.LeftChild != null) {
            assertEquals(level, tmp.Level);
            level++;
            tmp = tmp.LeftChild;
        }
    }

    @Test
    public void TestBalancedEmpty() {
         BalancedBST sut = new BalancedBST();
        int[] a = new int[]{};
        sut.GenerateTree(a);

        assertTrue(sut.IsBalanced(sut.Root));
    }

   @Test
    public void TestBalancedSingle() {
         BalancedBST sut = new BalancedBST();
        int[] a = new int[]{10};
        sut.GenerateTree(a);

        assertTrue(sut.IsBalanced(sut.Root));
    }
    
    @Test
    public void TestBalanced() {
         BalancedBST sut = new BalancedBST();
        int[] a = new int[]{1,2,3,4,5,6,7,8,9,10};
        sut.GenerateTree(a);

        assertTrue(sut.IsBalanced(sut.Root));
        assertTrue(sut.IsBalanced(sut.Root.LeftChild));
        assertTrue(sut.IsBalanced(sut.Root.RightChild));
    }
}
