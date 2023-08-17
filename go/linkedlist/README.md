# LinkedList

A linked\connected recursive data structure that is represented as a head and tail (another list). 
In the imperative version, it is implemented using pointers.
Unlike an array, where the data is arranged in a contiguous block, the elements of a list can be "scattered" in memory.

## Properties

Linked list can be:
1) Empty or filled
2) Contains duplicates or not

## Implementation

Be carefull with:
1) Value of head and tile while inserting/deleting elements
2) Shuffle elemenets
3) Deleting all elements with value V 

## Advantages & Disadvantages
:heavy_plus_sign: Fast adding/removing elemts from begining and end

:heavy_plus_sign: Flexible memory allocation

:heavy_minus_sign: `O(n)` searching

## Complexetiy 

### Memory
Memory complexery: `O(n)`

# Time:
If we have tail pointer:

| Insert | Delete | Search |
|:------:|:------:|:------:|
| `O(1)`   | `O(1)`   | `O(n)`   |


If we haven't tail pointer:

| Insert Head | Insert Tail | Delete | Search |
|:-----------:|:-----------:|:------:|:------:|
| `O(1)`      |    `O(n)`   | `O(n)` | `O(n)` |
