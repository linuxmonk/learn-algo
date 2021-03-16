## Sorting

### Selection Sort

This sorting algorithm iterate through the list and tries to determine the lowest / highest value in the list. When found, the element is swapped with the element from the unsorted part of the list. When the entire array has been traversed looking for the next smallest / largest element and swapped with it's position the process completes yielding a sorted array. The two main operations of this algorithm are -

- Comparison
- Swapping

*Comparisons*

In the worst case the number of comparisons we make is same as bubble sort. We compare every element to its next value. 

*Swaps*

However the number of swaps per pass is only one. Where as in bubble sort every comparison was also accompanied by a swap. 
SelectionSort is slightly better than bubble sort.

In the worst case scenario when all the elements need sorting the numbers below indicate the number of ops performed -


| UnsortedIndex | Array        | Lowest Index | Compare       | Swaps |
|---------------|--------------|--------------|---------------|-------|
| 0             | [5,4,3,2,1]  |  4           | 4             | 1     |
| 1             | [1,4,3,2,5]  |  3           | 3             | 1     |
| 2             | [1,2,3,4,5]  |  2           | 2             | 0     |
| 3             | [1,2,3,4,5]  |  3           | 1             | 0     |
| 4             | [1,2,3,4,5]  |  4           | 0             | 0     |
|               |              |              |               |       |
|               |    TOTAL     |              | 10            | 02    |

Total of 12 operations.
