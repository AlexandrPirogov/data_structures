# Dynamic Array

Unlike lists, which have O(n) element access, not counting the beginning and end of the list, dynamic arrays can be used. 
In this case, we will have constant access to the element. When adding or removing an element, it will be necessary to shift all the elements in O(n).
It is worth paying attention to when we should expand the array: when the array is completely filled or by a certain percentage? 
The same question goes for narrowing an array.
When expanding an array, you will need to copy all the elements from the old array to the new one. 
This means our insert operation takes O(n) on average. The situation is similar with the update.

## Properties

Dynamic array can be:
1) Empty or filled
2) Filled to a certain percentage

## Implementation

Be carefull with:
1) Accesing elemets by index
2) Moving elements after insertion/delete
3) Reallocation when array if fulfilled/almost empty

## Advantages & Disadvantages

:heavy_plus_sign: Accesing elements by index `O(1)`


:heavy_minus_sign: Searching elements takes `O(n)`

:heavy_minus_sign: Insertion and deleting elements takes `O(n)`

:heavy_minus_sign: Hard to choose the best reallocation scheme

## Complexety 

### Memory
Memory complexety: `O(n)`

### Time complexety:
| Insert | Delete | Search | Index access| 
|:------:|:------:|:------:|:-----------:|
|`O(n)`  | `O(n)` | `O(n)` |  `O(1)`     |
