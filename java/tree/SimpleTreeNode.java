import java.util.*;

public class SimpleTreeNode<T> {
  public T NodeValue; // значение в узле
  public SimpleTreeNode<T> Parent; // родитель или null для корня
  public List<SimpleTreeNode<T>> Children; // список дочерних узлов или null
  public boolean Visited;

  public SimpleTreeNode(T val, SimpleTreeNode<T> parent) {
    NodeValue = val;
    Parent = parent;
    Children = null;
    Visited = false;
  }

  public void AddChild(SimpleTreeNode<T> child) {
    if (Children == null) {
      Children = new ArrayList<SimpleTreeNode<T>>();
    }
    child.Parent = this;
    Children.add(child);
  }
}

class SimpleTree<T> {
  public SimpleTreeNode<T> Root; // корень, может быть null

  public SimpleTree(SimpleTreeNode<T> root) {
    Root = root;
  }

  public void AddChild(SimpleTreeNode<T> ParentNode, SimpleTreeNode<T> NewChild) {
    // ваш код добавления нового дочернего узла существующему ParentNode
    if (Root == null) {
      Root = NewChild;
      Root.Children = new ArrayList<SimpleTreeNode<T>>();
      return;
    }

    if (ParentNode == null) {
      return;
    }

    Optional<SimpleTreeNode<T>> res = recSearch(ParentNode);
    res.ifPresent(acc -> acc.AddChild(NewChild));
  }

  public void DeleteNode(SimpleTreeNode<T> NodeToDelete) {
    // ваш код удаления существующего узла NodeToDelete
    if (Root == null) {
      return;
    }

    if (Root.equals(NodeToDelete)) {
      Root = null;
      return;
    }

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();

      if (Node.equals(NodeToDelete)) {
        Node.Parent.Children.remove(Node);
        Node = null;
        return;
      }

      if (Node.Children == null || Node.Children.size() == 0) {
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }

  }

  public List<SimpleTreeNode<T>> GetAllNodes() {
    List<SimpleTreeNode<T>> res = new ArrayList<SimpleTreeNode<T>>();
    // ваш код выдачи всех узлов дерева в определённом порядке
    if (Root == null) {
      return res;
    }

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();
      res.add(Node);

      if (Node.Children == null || Node.Children.size() == 0) {
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }

    return res;
  }

  public List<SimpleTreeNode<T>> FindNodesByValue(T val) {
    // ваш код поиска узлов по значению
    List<SimpleTreeNode<T>> res = new ArrayList<SimpleTreeNode<T>>();
    // ваш код выдачи всех узлов дерева в определённом порядке
    if (Root == null) {
      return res;
    }

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();

      if (Node.NodeValue.equals(val)) {
        res.add(Node);
      }

      if (Node.Children == null || Node.Children.size() == 0) {
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }

    return res;
  }

  public void MoveNode(SimpleTreeNode<T> OriginalNode, SimpleTreeNode<T> NewParent) {
    if (Root == null) {
      return;
    }

    Optional<SimpleTreeNode<T>> from = recSearch(OriginalNode);
    Optional<SimpleTreeNode<T>> to = recSearch(NewParent);

    if (from.isPresent() && to.isPresent()) {
      OriginalNode.Parent.Children.remove(from.get());
      NewParent.AddChild(from.get());
    }

  }

    public ArrayList<T> EvenTrees()
    {
      ArrayList<T> res = new ArrayList<T>();
      if (Root == null) {
        return res;
      }
      PostOrder(Root, 0, res);
      return res;
    }

  public int PostOrder(SimpleTreeNode<T> from, int count, ArrayList<T> res) {
    if (from.Children == null) {
      return 1;
    }
    for (SimpleTreeNode<T> node : from.Children) {
      if (!node.Visited) {
        from.Visited = true;
        
        count += PostOrder(node, 0, res);
      }
    }
    if ((count+1)%2==0 && from != Root) {
      res.add(from.Parent.NodeValue);
      res.add(from.NodeValue);
    }
    return count+1;
  }

  public int Count() {
    if (Root == null) {
      return 0;
    }
    int count = 0;

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();
      count++;

      if (Node.Children == null || Node.Children.size() == 0) {
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }

    return count;
  }

  public int LeafCount() {
    if (Root == null) {
      return 0;
    }

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);
    int leafs = 0;

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();

      if (Node.Children == null || Node.Children.size() == 0) {
        leafs++;
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }

    return leafs;
  }

  private Optional<SimpleTreeNode<T>> recSearch(SimpleTreeNode<T> Desire) {
    if (Root == null) {
      return Optional.empty();
    }

    Stack<SimpleTreeNode<T>> stack = new Stack<SimpleTreeNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      SimpleTreeNode<T> Node = stack.pop();
      if (Node.equals(Desire)) {
        return Optional.of(Node);
      }

      if (Node.Children == null || Node.Children.size() == 0) {
        continue;
      }

      for (SimpleTreeNode<T> child : Node.Children) {
        stack.add(child);
      }
    }
    return Optional.empty();
  }
}