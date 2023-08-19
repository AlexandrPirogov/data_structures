# Ordered List

Ordered List gives as a fast way to retrieve max/min value from stream of data.
It kind of Priority Queue that implemented via linked list. 
There are also any different types of priority queue and implementations, for example via heap.
Priority queue highly used by OS for process scheduling and other internal tasks.


## Properties

Ordered List can be:
1) Empty or filled
2) Contain duplicates or not
3) Fast way to find if value exists in queue if value greater/less than queue's maximum/minimum

## Implementation

Be carefull with:
1) Nodes' next/prev pointers values

## Advantages & Disadvantages
:heavy_plus_sign: `O(1)` retrieve complexety 

:heavy_plus_sign: widely used and simple to implement with double linked list

:heavy_minus_sign: insertion `O(n)`. 

## Complexetiy 

### Memory
Memory complexery: `O(n)`

# Time:

| Insert | Delete | Search | Retrieve Max/Min|
|:------:|:------:|:------:|:----------------|
| `O(n)` | `O(n)` |`O(n)`  |    `O(1)`
