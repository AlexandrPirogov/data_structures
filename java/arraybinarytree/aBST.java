import java.util.*;

 class aBST
{
    public Integer Tree []; // массив ключей
  
    public aBST(int depth)
    {
      // правильно рассчитайте размер массива для дерева глубины depth:
      int tree_size = (int) Math.pow(2, depth+1)-1;

      Tree = new Integer[ tree_size ];
      for(int i=0; i<tree_size; i++) Tree[i] = null;
    }
  
    public Integer FindKeyIndex(int key)
    {
        int pos = 0;
        while (pos < Tree.length && Tree[pos] != null && Tree[pos] != key) {
            if (Tree[pos] > key) {
                pos = pos * 2 +1;
            } else {
                pos = pos * 2 + 2;
            }
        }

        if (pos >= Tree.length) {
            return null;
        }

        if (Tree[pos] == null) {
            return pos * -1;
        }
      // ищем в массиве индекс ключа
      return pos; // не найден
    }
  
    public int AddKey(int key)
    {
        Integer pos = FindKeyIndex(key);
        if (pos == null) {
            return -1;
        }

        if (pos <= 0) {
            Tree[-1 * pos] = key;
            return -1*pos;
        }


      // добавляем ключ в массив
      return -1; 
      // индекс добавленного/существующего ключа или -1 если не удалось
    }
  
}