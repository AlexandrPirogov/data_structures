# Double linked list

Adding an extra pointer allows you to traverse the list both ways, which can help when deleting nodes. 
There is also such a unique ability of doubly linked lists as efficient reordering of nodes without additional memory allocation.

## Advantages:
1) Deleting nodes much easier than in linked list
2) Easy to shuffle nodes in double linked list

## Disadvantages:
1) Still searching elements takes `O(n)`

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
