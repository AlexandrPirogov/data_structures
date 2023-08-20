# Hashing

To understand what hashing is we need some math. Imagine that we have set `U` of objects and set `S` of memory parts.
Hashinh is a function F: u -> s, where `u` belongs to `U` and `s` belongs to `S`. We can map elements of `U` to memory parts of `S`.
Having that function we can fast find `u` if it stored in `s`. So process to "convert" u to s is called hashing.

Usually we have memory parts of size `|S|` and `U > S`. There is a situation occurs when two distinct elements from `U`
after hashing them can point to the same memory parts of `S`. This sitation is called as collision.

There are many methods to resolve collision, but here two general:
1) Open addressing
2) Separate chaining
3) Cache-conscious collision resolution
