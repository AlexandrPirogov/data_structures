import java.util.*;

import javax.swing.tree.TreeNode;

class BSTNode {
  public int NodeKey; // ключ узла
  public BSTNode Parent; // родитель или null для корня
  public BSTNode LeftChild; // левый потомок
  public BSTNode RightChild; // правый потомок
  public int Level; // глубина узла

  public BSTNode(int key, BSTNode parent) {
    NodeKey = key;
    Parent = parent;
    LeftChild = null;
    RightChild = null;
  }
}

class BalancedBST {
  public BSTNode Root; // корень дерева

  public BalancedBST() {
    Root = null;
  }

  public void GenerateTree(int[] a) {
    Arrays.sort(a);
    if (a.length == 1) {
      Root = new BSTNode(a[0], null);
      return;
    }
    // создаём дерево с нуля из неотсортированного массива a
    // ...
    Root = generate(0, a.length - 1, a, null);

  }

  private BSTNode generate(int l, int r, int[] source, BSTNode parent) {
    if (l > r) {
      return null;
    }

    int mid = (l + r) / 2;
    BSTNode curr = new BSTNode(source[mid], parent);
    curr.Level = parent == null ? 0 : parent.Level+1;
    curr.LeftChild = generate(l, mid-1, source, curr);
    curr.RightChild = generate(mid + 1, r, source, curr);
    return curr;
  }

  public boolean IsBalanced(BSTNode root_node) {
    if (root_node == null) {
      return true;
    }
    return balanceHeight(root_node) != -1; // сбалансировано ли дерево с корнем root_node
  }

 private int balanceHeight (BSTNode from)
    {
        if (from == null)
        {
            return 0;
        }

        // checking left subtree
        int leftSubtreeHeight = balanceHeight (from.LeftChild);
        if (leftSubtreeHeight == -1) return -1;
        // if left subtree is not balanced then the entire tree is also not balanced

        //checking right subtree
        int rightSubtreeHeight = balanceHeight (from.RightChild);
        if (rightSubtreeHeight == -1) return -1;
        // if right subtree is not balanced then the entire          tree is also not balanced

        //checking the difference of left and right subtree for current node
        if (Math.abs(leftSubtreeHeight - rightSubtreeHeight) > 1)
        {
            return -1;
        }
        //if it is balanced then return the height
        return (Math.max(leftSubtreeHeight, rightSubtreeHeight) + 1);
    }
}