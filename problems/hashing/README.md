## Hash Tables

Every programming language has hashtables implemented and is known under various names like map, dictionary, hash map, associative arrays etc. The time complexity of retriving an element in a hash map is O(1).

Hash tables rely on a hash function to get the O(1) value. For a hash function to be correct it must return the same hash for a given input string always. A basic implemenentation of a data structure to store the hash table would be an array 

```
| v | vx | vy | .... | .... | .... | vz | <-- CELLS
```

The values v .. vz are various values that are stored for a set of given keys that are hashed. So to retrieve a value for a given key (k) all the hash table has to do is 

```
index = hashfn(key)
return arr[index]
```

Thus having a constant time. There are further complexities to this sort of implementation -

* Sophistication of a hash fn. Does it ever give out same indexes for two or more keys (clashing/collision)
* Where to store the keys ?

If two keys give out the same index then we say there is a collision. One way to handle collision is to have a reference to an array at the collision index and store all the values in the referenced array.

```
| v | vx | .... | arr-ref | .... | .... | vz |
                     |
		     V
	|| kn : vn |,| km : vm ||
```

So when -

```
v := hash[km] 
```

Will result in a clash. The hash table lookup function searches the `arr-ref` in a loop until it finds `km` returns `vm`.

*NOTE* If all the keys end up in the same index then in such a case the worst case performance of a hash table will be O(N). So it is critical that the hash function is designed in such a way that there are little / no collisions.

### Efficiency

Hash table's efficiency depends on -

* Amount of data being stored
* Number of cells available in the hash table
* Hash function being used

If there is too much data and less number of cells then the efficiency drops. 
If the hash function results in many collisions the efficiency drops again.
Data must be uniformly distributed across the cells to avoid collision.

Having too many cells to avoid collisions leads to hogging memory. A good hash table implemenentation has to balance the number of cells to the amount of data. This percentage of number of cells to amount of data is called the *_load factor_*. A computer scientist determines the load factor for the hash table. As more data is added to the hash table the number of cells is increased so that the distribution of data is uniform across the cells.

An interesting application is searching. Hash table increases the performance of searches. Hash table works well with pairs. But it can also be used with lists/arrays. If we convert all the elements of an array into hash keys then searching the key becomes a O(1) instead of O(N).

### NOTE on arraysubset

At the outset the algorithm `SubsetV1` looks faster than `SubsetV2`. Examining the pprof cpu profile results it can be seen that the `SubsetV2` indeed takes less time.
