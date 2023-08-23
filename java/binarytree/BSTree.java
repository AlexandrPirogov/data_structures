import java.io.*;
import java.util.*;

class BSTNode<T> {
  public int NodeKey; // ключ узла
  public T NodeValue; // значение в узле
  public BSTNode<T> Parent; // родитель или null для корня
  public BSTNode<T> LeftChild; // левый потомок
  public BSTNode<T> RightChild; // правый потомок

  public BSTNode(int key, T val, BSTNode<T> parent) {
    NodeKey = key;
    NodeValue = val;
    Parent = parent;
    LeftChild = null;
    RightChild = null;
  }
}

// промежуточный результат поиска
class BSTFind<T> {
  // null если в дереве вообще нету узлов
  public BSTNode<T> Node;

  // true если узел найден
  public boolean NodeHasKey;

  // true, если родительскому узлу надо добавить новый левым
  public boolean ToLeft;

  public BSTFind() {
    Node = null;
  }
}

class BST<T> {
  BSTNode<T> Root; // корень дерева, или null

  public BST(BSTNode<T> node) {
    Root = node;
  }

  public BSTFind<T> FindNodeByKey(int key) {
    BSTFind<T> res = new BSTFind<T>();
    if (Root == null) {
      res.Node = null;
      res.NodeHasKey = false;
      return res;
    }

    BSTNode<T> curr = Root;
    BSTNode<T> tmp = curr;
    while (curr != null && curr.NodeKey != key) {
      tmp = curr;
      curr = curr.NodeKey > key ? curr.LeftChild : curr.RightChild;
    }

    if (curr == null) {
      res.Node = tmp;
      res.NodeHasKey = false;
      res.ToLeft = (key < tmp.NodeKey) ? true : false;
    } else {
      res.Node = curr;
      res.NodeHasKey = true;
    }

    return res;
  }

  public boolean AddKeyValue(int key, T val) {
    if (Root == null) {
      Root = new BSTNode<T>(key, val, null);
      return true;
    }

    BSTFind<T> find = FindNodeByKey(key);
    if (find.NodeHasKey) {
      return false;
    }

    BSTNode<T> add = new BSTNode<T>(key, val, find.Node);
    if (find.ToLeft) {
      find.Node.LeftChild = add;
      add.Parent = find.Node;
    } else {
      find.Node.RightChild = add;
      add.Parent = find.Node;
    }

    // добавляем ключ-значение в дерево
    return true; // если ключ уже есть
  }

  public BSTNode<T> FinMinMax(BSTNode<T> FromNode, boolean FindMax) {
    return Root == null ? null : (FindMax ? max(FromNode) : min(FromNode));
  }

  public boolean DeleteNodeByKey(int key) {
    if (Root == null) {
      Root = null;
      return false;
    }

    BSTFind<T> find = FindNodeByKey(key);
    if (!find.NodeHasKey) {
      return false;
    }

    BSTNode<T> curr = find.Node;
    if (curr.LeftChild == null && curr.RightChild == null) {
      if (Root == curr) {
        Root = null;
        return true;
      }

      curr = curr.Parent;
      if (curr.LeftChild.NodeKey == key) {
        curr.LeftChild = null;
      } else {
        curr.RightChild = null;
      }

      return true;
    }

    if (curr.LeftChild != null && curr.RightChild == null) {
      curr = curr.Parent;
      curr.LeftChild = curr.LeftChild.LeftChild;
      return true;
    }

    BSTNode<T> toReplace = FinMinMax(curr.RightChild, false);
    curr.NodeKey = toReplace.NodeKey;
    curr.NodeValue = toReplace.NodeValue;

    BSTNode<T> tmp = curr.RightChild;
    if (tmp.LeftChild == null) {
      curr.RightChild = curr.RightChild.RightChild;
      return true;
    }

    while (tmp.LeftChild.LeftChild != null) {
      tmp = tmp.LeftChild;
    }

    tmp.LeftChild = null;
    return true;

  }

  public int Count() {
    if (Root == null) {
      return 0;
    }

    int count = 0;
    Stack<BSTNode<T>> stack = new Stack<BSTNode<T>>();
    stack.add(Root);

    while (!stack.empty()) {
      BSTNode<T> curr = stack.pop();
      count++;

      if (curr.LeftChild != null) {
        stack.add(curr.LeftChild);
      }

      if (curr.RightChild != null) {
        stack.add(curr.RightChild);
      }
    }

    return count; // количество узлов в дереве
  }

  private BSTNode<T> min(BSTNode<T> FromNode) {
    BSTNode<T> tmp = FromNode;
    while (tmp.LeftChild != null) {
      tmp = tmp.LeftChild;
    }

    return tmp;
  }

  private BSTNode<T> max(BSTNode<T> FromNode) {
    BSTNode<T> tmp = FromNode;
    while (tmp.RightChild != null) {
      tmp = tmp.RightChild;
    }

    return tmp;
  }
}