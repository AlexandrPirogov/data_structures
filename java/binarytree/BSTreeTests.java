import static org.junit.Assert.*;

import org.junit.Test;

public class BSTreeTests {

    @Test
    public void TestCreate() {
        BST<Integer> sut = new BST<Integer>(null);
        // assert that count == 0 and root is null

        assertEquals(0, sut.Count());
    }

    @Test
    public void TestAddKeyInEmpty() {
        BSTNode<Integer> root = new BSTNode(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        // assert that root == added key and count == 1
        assertEquals(1, sut.Count());
        assertEquals(root, sut.Root);
    }

    @Test
    public void TestAddOnlyLeftKeys() {
        int k = 10;
        int v = 10;

        BSTNode<Integer> root = new BSTNode(k, v, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = k; i > 0; i--) {
            sut.AddKeyValue(k, v);
            k--;
            v--;
        }

        // assert that root-> right is null and count == bunch of nodes
        assertNull(sut.Root.RightChild);
        assertEquals(10, sut.Count());
    }

    @Test
    public void TestAddOnlyRightKeys() {
        int k = 10;
        int v = 10;

        BSTNode<Integer> root = new BSTNode(k, v, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = k; i > 0; i--) {
            sut.AddKeyValue(k, v);
            k++;
            v++;
        }

        // assert that root-> right is null and count == bunch of nodes
        assertNull(sut.Root.LeftChild);
        assertEquals(10, sut.Count());
    }

    @Test
    public void AddAllEqualsNodes() {
        int k = 0;
        int v = 0;

        BSTNode<Integer> root = new BSTNode(k, v, null);
        BST<Integer> sut = new BST<Integer>(root);
        for (int i = 0; i < 10; i++) {
            sut.AddKeyValue(k, v);
        }

        assertNull(sut.Root.LeftChild);
        assertNull(sut.Root.RightChild);
        assertEquals(1, sut.Count());
    }

    @Test
    public void TestFindMinRoot() {
        BSTNode<Integer> root = new BSTNode(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = -5; i < 6; i++) {
            sut.AddKeyValue(i, i);
        }

        BSTNode<Integer> min = sut.FinMinMax(root, false);
        assertEquals(-5, min.NodeKey);
    }

    @Test
    public void TestFindMax() {
        BSTNode<Integer> root = new BSTNode(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = -5; i < 6; i++) {
            sut.AddKeyValue(i, i);
        }

        BSTNode<Integer> min = sut.FinMinMax(root, true);
        assertEquals(5, min.NodeKey);
    }

    @Test
    public void TestFindMinEmpty() {
        BST<Integer> sut = new BST<Integer>(null);
        BSTNode<Integer> res = sut.FinMinMax(sut.Root, false);
        assertEquals(null, res);
    }

    @Test
    public void TestFindMaxEmpty() {
        BST<Integer> sut = new BST<Integer>(null);
        BSTNode<Integer> res = sut.FinMinMax(sut.Root, true);
        assertEquals(null, res);
    }

    @Test
    public void TestFindValInEmpty() {
        BST<Integer> sut = new BST<Integer>(null);
        BSTFind<Integer> res = sut.FindNodeByKey(-1);

        assertNull(res.Node);
        assertFalse(res.NodeHasKey);
        // assert that we can't find val in empty tree
    }

    @Test
    public void TestFindExistingVal() {

        BSTNode<Integer> root = new BSTNode(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = -5; i < 6; i++) {
            sut.AddKeyValue(i, i);
        }

        for (int i = -5; i < 6; i++) {
            BSTFind<Integer> res = sut.FindNodeByKey(i);
            assertTrue(res.NodeHasKey);
        }

    }

    @Test
    public void TestFindNotExistingVal() {
        BSTNode<Integer> root = new BSTNode(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        for (int i = -5; i < 6; i++) {
            sut.AddKeyValue(i, i);
        }

        for (int i = 20; i < 30; i++) {
            BSTFind<Integer> res = sut.FindNodeByKey(i);
            assertFalse(res.NodeHasKey);
        }
    }

    @Test
    public void TestDeleteFromEmpty() {
        BST<Integer> sut = new BST<Integer>(null);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(0);
        int after = sut.Count();

        assertFalse(res);
        assertEquals(before, after);
    }

    @Test
    public void TestDeleteRootSingle() {
        BSTNode<Integer> root = new BSTNode<Integer>(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(0);
        int after = sut.Count();

        assertTrue(res);
        assertNull(sut.Root);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteRootWithRightChild() {
        BSTNode<Integer> root = new BSTNode<Integer>(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);
        sut.AddKeyValue(1, 1);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(0);
        int after = sut.Count();

        assertTrue(res);
        assertNotNull(sut.Root);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteSubtreeWithoutRight() {
        BSTNode<Integer> root = new BSTNode<Integer>(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        sut.AddKeyValue(-1, -1);
        sut.AddKeyValue(-2, -2);
        sut.AddKeyValue(-3, -3);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(-2);
        int after = sut.Count();
        BSTFind<Integer> find = sut.FindNodeByKey(-2);

        assertTrue(res);
        assertFalse(find.NodeHasKey);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteSubtreeWithRightWithLeft() {
        BST<Integer> sut = new BST<Integer>(null);

        sut.AddKeyValue(8, 8);
        sut.AddKeyValue(4, 4);
        sut.AddKeyValue(6, 6);
        sut.AddKeyValue(5, 5);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(4);
        int after = sut.Count();

        assertTrue(res);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteSubtreeWithRightWithoutLeft() {
        BSTNode<Integer> root = new BSTNode<Integer>(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        sut.AddKeyValue(8, 8);
        sut.AddKeyValue(4, 4);
        sut.AddKeyValue(6, 6);
        sut.AddKeyValue(7, 7);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(4);
        int after = sut.Count();

        assertTrue(res);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteSubtreeWithRightWithLeft2() {
        BSTNode<Integer> root = new BSTNode<Integer>(0, 0, null);
        BST<Integer> sut = new BST<Integer>(root);

        sut.AddKeyValue(8, 8);
        sut.AddKeyValue(4, 4);
        sut.AddKeyValue(6, 6);
        sut.AddKeyValue(7, 7);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(4);
        int after = sut.Count();

        assertTrue(res);
        assertEquals(before - 1, after);
    }

    @Test
    public void TestDeleteCustom() {
        BST<Integer> sut = new BST<Integer>(null);
        int[] arr = new int[] { 8, 4, 12, 2, 1, 3, 6, 5, 7, 10, 9, 11, 14, 13, 15 };
        for (Integer num : arr) {
            sut.AddKeyValue(num, num);
        }

        for (Integer num : arr) {

            int before = sut.Count();
            boolean res = sut.DeleteNodeByKey(num);
            int after = sut.Count();

            BSTFind<Integer> find = sut.FindNodeByKey(num);

            assertFalse(find.NodeHasKey);
            assertTrue(res);
            assertEquals(before - 1, after);
        }
    }

    @Test
    public void TestDeleteMin() {
        BST<Integer> sut = new BST<Integer>(null);
        int[] arr = new int[] { 8, 4, 12, 2, 1, 3, 6, 5, 7, 10, 9, 11, 14, 13, 15 };
        for (Integer num : arr) {
            sut.AddKeyValue(num, num);
        }

        BSTNode<Integer> min = sut.FinMinMax(sut.Root, false);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(min.NodeKey);
        int after = sut.Count();

        BSTFind<Integer> find = sut.FindNodeByKey(min.NodeKey);

        assertFalse(find.NodeHasKey);
        assertTrue(res);
        assertEquals(before - 1, after);

    }

    @Test
    public void TestDeleteMax() {
        BST<Integer> sut = new BST<Integer>(null);
        int[] arr = new int[] { 8, 4, 12, 2, 1, 3, 6, 5, 7, 10, 9, 11, 14, 13, 15 };
        for (Integer num : arr) {
            sut.AddKeyValue(num, num);
        }

        BSTNode<Integer> min = sut.FinMinMax(sut.Root, true);

        int before = sut.Count();
        boolean res = sut.DeleteNodeByKey(min.NodeKey);
        int after = sut.Count();

        BSTFind<Integer> find = sut.FindNodeByKey(min.NodeKey);

        assertFalse(find.NodeHasKey);
        assertTrue(res);
        assertEquals(before - 1, after);

    }
}
