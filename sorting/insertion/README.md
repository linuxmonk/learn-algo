## Sorting

### Insertion Sort

This sorting algorithm starts at index 1 of the list as the pivot. The value at the pivot is held in a temporary variable. It then compares all the elements to the left of the pivot with the value in temporary variable (pivot). If the value on the left is greater / lesser than pivot then that value is right shifted towards the pivot. When there are no more elements on the left or when the element on the list on the left is no longer greater than pivot then the shifting stops. The temporary variable at the pivot is inserted back into the list where the shifting stops. The main operations of the algorithms are -

- Comparison
- Shifting
- Inserting

In the worst case scenario when all the elements need sorting the numbers below indicate the number of ops performed -


| Pivot | Array        | Compare       | Shift | Insert |
|-------|--------------|---------------|-------|--------|
| 1     | [5,4,3,2,1]  | 1             | 1     | 1      |
| 2     | [4,5,3,2,1]  | 2             | 2     | 1      |
| 3     | [3,4,5,2,1]  | 3             | 3     | 1      |
| 4     | [2,3,4,5,1]  | 4             | 4     | 1      |
| 5     | [1,2,3,4,5]  |               |       |        |
|       |              |               |       |        |
|       |    TOTAL     | 10            | 10    | 4      |

Total of 24 operations.
