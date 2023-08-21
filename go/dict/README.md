# Dictionary

A dictionary is a data structure that stores objects as key-values.
Based on the hashing of the key, we get the position where the object will be stored and access to it will be calculated in constant time. 
There are many variations in the implementation of the dictionary: through trees, dynamic arrays, or multilevel hash tables.
The efficiency of operations depends on the choice of hash function.

## Properties

Dictionary can contains a lot of collisions. It highly affects complexety of operations

## Implementation

Be carefull with:
1) Choosing hash-function
2) Choosing method for collision resolution

## Advantages & Disadvantages
:heavy_plus_sign: Low average complexety

:heavy_minus_sign: Highly depends on hash-function and collision resolution.

## Complexetiy 

### Memory

Usualy memory complexery: `O(n)`. Highly depends on implementation

# Time:


Average time when collisions are rarely met:
| Insert | Delete | Search |
|:------:|:------:|:------:|
| `o(1)` | `o(1)` | `O(1)` |


Worst time when collisions are met often:
| Insert | Delete | Search |
|:------:|:------:|:------:|
| `O(n)` | `O(n)` | `O(n)` |
