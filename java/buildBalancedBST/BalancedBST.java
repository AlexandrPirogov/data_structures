import java.util.*;

class BSTNode {
  public int NodeKey; 
  public BSTNode Parent; 
  public BSTNode LeftChild;
  public BSTNode RightChild;
  public int Level;

  public BSTNode(int key, BSTNode parent) {
    NodeKey = key;
    Parent = parent;
    LeftChild = null;
    RightChild = null;
  }
}

class BalancedBST {
  public BSTNode Root;

  public BalancedBST() {
    Root = null;
  }

  public void GenerateTree(int[] a) {
    Arrays.sort(a);
    if (a.length == 1) {
      Root = new BSTNode(a[0], null);
      return;
    }

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
    return balanceHeight(root_node) != -1; 
  }

 private int balanceHeight (BSTNode from)
    {
        if (from == null)
        {
            return 0;
        }

        int leftSubtreeHeight = balanceHeight (from.LeftChild);
        if (leftSubtreeHeight == -1) return -1;

        int rightSubtreeHeight = balanceHeight (from.RightChild);
        if (rightSubtreeHeight == -1) return -1;

        if (Math.abs(leftSubtreeHeight - rightSubtreeHeight) > 1)
        {
            return -1;
        }

        return (Math.max(leftSubtreeHeight, rightSubtreeHeight) + 1);
    }
}