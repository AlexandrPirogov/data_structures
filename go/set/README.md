# Set

Sets are data structure that contains only distinct elements. It can be easily implemented via dictionary.

## Implementation

Be carefull with:
1) Choosing hash-function
2) Choosing method for collision resolution

## Advantages & Disadvantages
:heavy_plus_sign: Low time complexety to determine if element exists in set

:heavy_minus_sign: Other operations have time complexety greater than `O(1)`

## Complexetiy 

### Memory

Memory complexety: `O(n)`

# Time:


Worst time when collisions are met often:
| Insert | Delete | Search | Determine if exists element |
|:------:|:------:|:------:|:---------------------------:|
| `O(n)` | `O(n)` | `O(n)` |         `O(1)`              |
