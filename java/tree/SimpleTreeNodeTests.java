import static org.junit.Assert.*;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class SimpleTreeNodeTests {

    @Test
    public void Create() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        assertTrue(sut.Count() == 0);
    }

    @Test
    public void AddEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        assertTrue(sut.Count() == 1);
        assertTrue(sut.Root == root);
    }

    @Test
    public void AddChildrenToExistingRoot() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 1; i < 11; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        assertTrue(sut.Count() == 11);
    }

    @Test
    public void AddChildrenToChild() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);
        sut.AddChild(root, new SimpleTreeNode<Integer>(1, null));

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i - 1, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        assertTrue(sut.Count() == 12);
    }

    @Test
    public void AddChildToNotExisting() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        SimpleTreeNode<Integer> notExisting = new SimpleTreeNode<Integer>(-1, null);
        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(-1, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(notExisting, node);
        }

        assertTrue(sut.Count() == 1);
    }

    @Test
    public void AddChildToNull() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(-1, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(null, node);
        }

        assertTrue(sut.Count() == 1);
    }

    @Test
    public void GetNodesFromEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        List<SimpleTreeNode<Integer>> res = sut.GetAllNodes();

        assertTrue(res.size() == 0);
    }

    @Test
    public void GetAllNodesFromFilled() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);
        sut.AddChild(root, new SimpleTreeNode<Integer>(1, null));

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i - 1, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        List<SimpleTreeNode<Integer>> res = sut.GetAllNodes();

        assertTrue(res.size() == sut.Count() && res.size() > 6);
    }

    @Test
    public void FindAllNodesFromEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        List<SimpleTreeNode<Integer>> res = sut.FindNodesByValue(-1);

        assertTrue(res.size() == 0);
    }

    @Test
    public void FindAllNodesFromFilled() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);
        sut.AddChild(root, new SimpleTreeNode<Integer>(1, null));

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            List<SimpleTreeNode<Integer>> res = sut.FindNodesByValue(node.NodeValue);
            assertTrue(res.size() == 1);
        }
    }

    @Test
    public void FindAllNodesFromFilledNotExisted() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);
        sut.AddChild(root, new SimpleTreeNode<Integer>(1, null));

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        for (int i = 0; i < 12; i++) {
            List<SimpleTreeNode<Integer>> res = sut.FindNodesByValue(-1);
            assertTrue(res.size() == 0);
        }
    }

    @Test
    public void LeafsCountInEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);

        assertTrue(sut.LeafCount() == 0);
    }

    @Test
    public void LeafsCountInSingle() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        assertTrue(sut.LeafCount() == 1);
    }

    @Test
    public void LeafsCountInFilled_1() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 1; i < 11; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i, null));
        }

        for (SimpleTreeNode<Integer> node : nodes) {
            sut.AddChild(root, node);
        }

        assertTrue(sut.LeafCount() == 10);
    }

    @Test
    public void DeleteFromEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> delete = new SimpleTreeNode<Integer>(1, null);

        int before = sut.Count();
        sut.DeleteNode(delete);
        int after = sut.Count();

        assertTrue(before == after);
    }

    @Test
    public void DeleteRoot() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        sut.AddChild(null, root);
        sut.AddChild(root, new SimpleTreeNode<Integer>(1, null));

        List<SimpleTreeNode<Integer>> nodes = new ArrayList<SimpleTreeNode<Integer>>();
        for (int i = 2; i < 12; i++) {
            nodes.add(new SimpleTreeNode<Integer>(i, null));
        }

        int before = sut.Count();
        sut.DeleteNode(root);
        int after = sut.Count();

        int leafs = sut.LeafCount();
        assertTrue(before - after == before);
        assertTrue(leafs == 0);
    }

    @Test
    public void DeleteExisting() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        SimpleTreeNode<Integer> l1 = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> l2 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> r2 = new SimpleTreeNode<Integer>(3, null);

        sut.AddChild(null, root);
        sut.AddChild(root, l1);
        sut.AddChild(l1, l2);
        sut.AddChild(l1, r2);

        sut.DeleteNode(l1);
        int after = sut.Count();
        int leafsA = sut.LeafCount();

        assertTrue(after == 1);
        assertEquals(1, leafsA);
    }

    @Test
    public void DeleteNotExisting() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        SimpleTreeNode<Integer> l1 = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> l2 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> r2 = new SimpleTreeNode<Integer>(3, null);
        SimpleTreeNode<Integer> notExists = new SimpleTreeNode<Integer>(3, null);

        sut.AddChild(null, root);
        sut.AddChild(root, l1);
        sut.AddChild(l1, l2);
        sut.AddChild(l1, r2);

        int before = sut.Count();
        sut.DeleteNode(notExists);
        int after = sut.Count();

        assertTrue(after == before);
    }

    @Test
    public void MoveInEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);

        SimpleTreeNode<Integer> from = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> to = new SimpleTreeNode<Integer>(1, null);

        sut.MoveNode(from, to);

        assertTrue(from.Parent == null);
    }

    @Test
    public void MoveInFilled1() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        SimpleTreeNode<Integer> l1 = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> l2 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> r2 = new SimpleTreeNode<Integer>(3, null);

        sut.AddChild(null, root);
        sut.AddChild(root, l1);
        sut.AddChild(l1, l2);
        sut.AddChild(l1, r2);

        int before = root.Children.size();
        sut.MoveNode(r2, root);
        int after = root.Children.size();

        assertTrue(after - 1 == before);
        assertTrue(r2.Parent.equals(root));
    }

    @Test
    public void MoveToNewParentWithNullChildren() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(0, null);
        SimpleTreeNode<Integer> l1 = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> r1 = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> ll2 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> lr2 = new SimpleTreeNode<Integer>(3, null);

        sut.AddChild(null, root);
        sut.AddChild(root, l1);
        sut.AddChild(root, r1);
        sut.AddChild(l1, ll2);
        sut.AddChild(l1, lr2);

        int leafsB = sut.LeafCount();
        sut.MoveNode(l1, r1);
        int after = r1.Children.size();
        int leafsA = sut.LeafCount();

        assertTrue(after == 1);
        assertTrue(l1.Parent.equals(r1));
        assertFalse(root.Children.contains(l1));
        assertEquals(leafsB - 1, leafsA);
    }

    @Test
    public void Testr() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> n1 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> n2 = new SimpleTreeNode<Integer>(3, null);
        SimpleTreeNode<Integer> n3 = new SimpleTreeNode<Integer>(6, null);

        SimpleTreeNode<Integer> n4 = new SimpleTreeNode<Integer>(5, null);
        SimpleTreeNode<Integer> n5 = new SimpleTreeNode<Integer>(7, null);
        SimpleTreeNode<Integer> n6 = new SimpleTreeNode<Integer>(4, null);
        SimpleTreeNode<Integer> n7 = new SimpleTreeNode<Integer>(8, null);

        SimpleTreeNode<Integer> n8 = new SimpleTreeNode<Integer>(9, null);
        SimpleTreeNode<Integer> n9 = new SimpleTreeNode<Integer>(10, null);

        sut.AddChild(null, root);
        sut.AddChild(root, n1);
        sut.AddChild(root, n2);
        sut.AddChild(root, n3);

        sut.AddChild(n1, n4);
        sut.AddChild(n1, n5);

        sut.AddChild(n2, n6);
        sut.AddChild(n2, n1);

        sut.AddChild(n3, n7);

        sut.AddChild(n7, n8);
        sut.AddChild(n7, n9);

        ArrayList<Integer> res = sut.EvenTrees();
        assertTrue(res.size() == 4);
    }

    @Test
    public void TestEvenTreesInEmpty() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        ArrayList<Integer> res = sut.EvenTrees();

        assertEquals(0, res.size());
    }

     @Test
    public void TestEvenTreesWithoutEvenTrees() {
        SimpleTree<Integer> sut = new SimpleTree<Integer>(null);
        SimpleTreeNode<Integer> root = new SimpleTreeNode<Integer>(1, null);
        SimpleTreeNode<Integer> n1 = new SimpleTreeNode<Integer>(2, null);
        SimpleTreeNode<Integer> n2 = new SimpleTreeNode<Integer>(3, null);
        SimpleTreeNode<Integer> n3 = new SimpleTreeNode<Integer>(6, null);


        sut.AddChild(null, root);
        sut.AddChild(root, n1);
        sut.AddChild(root, n2);
        sut.AddChild(root, n3);

        ArrayList<Integer> res = sut.EvenTrees();
        assertTrue(res.size() == 0);
    }
}