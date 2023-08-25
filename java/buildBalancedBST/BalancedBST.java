import java.util.*;

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
    if (Root == null) {
      return true;
    }
    return BFS(root_node.LeftChild)-BFS(root_node.RightChild) < 2; // сбалансировано ли дерево с корнем root_node
  }

  private int BFS(BSTNode from) {
    if (from == null) {
      return 0;
    }
    int len = 0;
    ArrayDeque<BSTNode> q = new ArrayDeque<BSTNode>();
    q.add(from);
    int level = 0;
    int lastAdded = 1;
    while(q.size() > 0) {
      int added = 0;
    
      int i = lastAdded;
      while (i > 0) {
        BSTNode item = q.poll();
        
        if (item.LeftChild != null) {
          added++;
          q.add(item.LeftChild);
        }

        if (item.RightChild != null) {
          added++;
          q.add(item.RightChild);
        }

        i--;
      }
      lastAdded = added;
      level++;
    }
    return level;
  }
}