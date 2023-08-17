# LinkedList

A linked\connected recursive data structure that is represented as a head and tail (another list). 
In the imperative version, it is implemented using pointers.
Unlike an array, where the data is arranged in a contiguous block, the elements of a list can be "scattered" in memory.

## Advantages:
1) Fast adding/removing elemts from begining and end
2) Flexible memory allocation

## Disadvantages:
1) O(n) searching

## Complexetiy 

If we have tail pointer:

| Insert | Delete | Search |
|:------:|:------:|:------:|
| O(1)   | O(1)   | O(n)   |


If we haven't tail pointer:

| Insert Head | Insert Tail | Delete | Search |
|:-----------:|:-----------:|:------:|:------:|
| O(1)        |      O(n)   | O(n)   | O(n)   |
