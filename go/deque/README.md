# Deque

Deque is a simple data structre that works simultaneously as FIFO and as LIFO.
Used for situation when we want to pick fast not only first element in the container, but it's last 
Easy to implement with double linked list.

## Properties

Dequeue can be:
1) Empty or filled

## Implementation

Be carefull with:
1) Nodes' next/prev pointers values 

## Advantages & Disadvantages
:heavy_plus_sign: `O(1)` push and pop complexetiy

:heavy_plus_sign: widely used and simple to implement with double linked list

:heavy_minus_sign: searching `O(n)`

## Complexetiy 

### Memory
Memory complexery: `O(n)`

# Time:

| Insert | Delete | Search |
|:------:|:------:|:------:|
| `O(1)` | `O(1)` |`O(n)`  |
