## Sorting

### Bubble Sort

This sorting algorithm compares an element in a position with the next element and performs a swap if the first element is less/greater than the second one. There are two main operations performed by the algorithm -

- Comparison
- Swapping

For this array example array - [5, 4, 3, 2, 1]

*Comparisons*

Number of comparisons in the first iteration 	= 4
Number of comparisons in the second iteration 	= 3
Number of comparisons in the third iteration 	= 2
Number of comparisons in the fourth iteration 	= 1
===================================================
Total number of comparisons			= 10

*Swaps*

Total number of swaps depend on the numbers/data. But in the worst case when the list is in the opposite order then number of swaps would be equal to number of comparisons.

1 + 2 + 3 .... N = ((N) * (N - 1)) / 2

| # of elements (N) | Ops (Comparisons + Swaps) |  N ^ 2 |
| -----------------:| -------------------------:| ------:|
| 5                 | 20                        |  25    |
| 10                | 180                       |  100   |
| 20                | 380                       |  400   |

The algorithms complexity tends to be N^2 (N-Squared).




