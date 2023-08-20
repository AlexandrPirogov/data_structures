# Dynamic Array

Unlike lists, which have O(n) element access, not counting the beginning and end of the list, dynamic arrays can be used. 
In this case, we will have constant access to the element. When adding or removing an element, it will be necessary to shift all the elements in O(n).
It is worth paying attention to when we should expand the array: when the array is completely filled or by a certain percentage? 
The same question goes for narrowing an array.
When expanding an array, you will need to copy all the elements from the old array to the new one. 
This means our insert operation takes O(n) on average. The situation is similar with the update.

## Properties

Dynamic array can be:
1) more than `k%` full. This greatly affects the number of collisions.

## Implementation

Be carefull with:
1) Choosing hash-function
2) Collision resolution

## Advantages & Disadvantages

:heavy_plus_sign: Accesing elements `o(1)`

:heavy_plus_sign: Insertion elements `o(1)`

:heavy_plus_sign: Deleting elements `o(1)`


:heavy_minus_sign: dictionary can contains lot of collision and all operations become `O(n)`

## Complexety 

### Memory
Memory complexety: `O(n)`. Depends on implementation

### Time complexety:

Average:
| Insert | Delete | Search | Index access| 
|:------:|:------:|:------:|:-----------:|
|`o(1)`  | `o(1)` | `o(1)` |  `o(1)`     |

Words (if we have a lot of collisions):
| Insert | Delete | Search | Index access| 
|:------:|:------:|:------:|:-----------:|
|`O(n)`  | `O(n)` | `O(n)` |  `O(n)`     |
