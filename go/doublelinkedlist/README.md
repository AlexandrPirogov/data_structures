# Double linked list

Adding an extra pointer allows you to traverse the list both ways, which can help when deleting nodes. 
There is also such a unique ability of doubly linked lists as efficient reordering of nodes without additional memory allocation.

## Properties

Linked list can be:
1) Empty or filled
2) Contains duplicates or not

## Implementation

Be carefull with:
1) Value of head and tile while inserting/deleting elements
2) Shuffle elemenets
3) Deleting all elements with value V 
4) Value of next/prev pointers

## Advantages & Disadvantages

:heavy_plus_sign: Deleting nodes much easier than in linked list

:heavy_plus_sign: Easy to shuffle nodes in double linked list


:heavy_minus_sign: Still searching elements takes `O(n)`

## Complexety 

### Memory
Memory complexety: `O(n)`

### Time complexety:
If we have tail pointer:

| Insert | Delete | Search |
|:------:|:------:|:------:|
|`O(1)`  | `O(1)` | `O(n)` |


If we haven't tail pointer:

| Insert Head | Insert Tail | Delete | Search |
|:-----------:|:-----------:|:------:|:------:|
| `O(1)`      |    `O(n)`   | `O(n)` | `O(n)` |
