import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;

import org.junit.Test;

public class aBST_tests {
    
    @Test
    public void TestSize() {
        aBST sut = new aBST(3);

        assertEquals(15, sut.Tree.length);
    }

    @Test
    public void TestFindInEmpty() {
        aBST sut = new aBST(5);

        int index = sut.FindKeyIndex(-1);

        assertEquals(index, 0);
    }

     @Test
    public void TestAddInEmpty() {
        aBST sut = new aBST(5);

        int index = sut.AddKey(0);
        int pos = sut.FindKeyIndex(0);

        assertEquals(index, pos);
    }

    @Test
    public void TestAddIndexesCorrectness() {
        aBST sut = new aBST(2);

        int i1 = sut.AddKey(0);
        int i2 = sut.AddKey(-1);
        int i3 = sut.AddKey(1);
        
        assertEquals(i1*2+1, i2);
        assertEquals(i1*2+2, i3);
    }

    @Test
    public void TestFindExisting() {
        aBST sut = new aBST(2);

        int i1 = sut.AddKey(0);
        int i2 = sut.AddKey(-1);
        int i3 = sut.AddKey(1);
        
        int p1 = sut.FindKeyIndex(0);
        int p2 = sut.FindKeyIndex(-1);
        int p3 = sut.FindKeyIndex(1);

        assertEquals(i1, p1);
        assertEquals(i2, p2);
        assertEquals(i3, p3);
    }


    @Test
    public void TestFindNotExistingInFullFilled() {
        aBST sut = new aBST(1);

        
        sut.AddKey(0);
        sut.AddKey(-1);
        sut.AddKey(1);
        
        Integer p1 = sut.FindKeyIndex(10);
        Integer p2 = sut.FindKeyIndex(-11);
        Integer p3 = sut.FindKeyIndex(11);

        assertNull(p1);
        assertNull(p2);
        assertNull(p3);
    }

     @Test
    public void TestFindNotExistingButCanInsert() {
        aBST sut = new aBST(2);

        
        int root = sut.AddKey(0);
        
        Integer left = sut.FindKeyIndex(-1);
        Integer right = sut.FindKeyIndex(1);

        assertEquals(root * 2 + 1, -1 * left.intValue());
        assertEquals(root * 2 + 2, -1 * right.intValue());
    }

    @Test
    public void TestAddInFullfilled() {
        aBST sut = new aBST(1);

        
        sut.AddKey(0);
        sut.AddKey(-1);
        sut.AddKey(1);
        
        int fail = sut.AddKey(2);

        assertEquals(-1, fail);
    }
}
