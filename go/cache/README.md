# Cache

Cache is a data structure that used to improve reading operations if you read often the same resources.
This data structure widely used in hardware, OS and other software.  CPU uses L1 cache to fetch most used instructions for example.
Cache usually has a size less than count of stored items. 
If you found needed resource in cache so you got a "hit". Otherwise it called "miss".
To build good cache you have to predict which items will be requested most of time.

## Properties

Cache has the main property: relation between hits and misses.
The more hits you've got -- the more effective cache is.

## Implementation

There are vague types of cache implementation from small parts of hardware such as caches for CPU to database systems like Redis.

## Advantages & Disadvantages
:heavy_plus_sign: Improve perfomance of reading operations


:heavy_minus_sign: Highly depends on data set and imlpementation

## Complexetiy 

Time complexety depends on relation between hits and misses.
