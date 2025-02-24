Slices
A slice can be created and initialized using the following ways:
- Using slice literal
- Using an array
- Using already existing slice
- Using make() funcion

You cat iterate over slice using the following ways:
- Using for loop
- Using range in for loop
- Using a blank identifier in for loop

Slice is a reference type it can refer an underlying array. If we change some elements in the slice, then the changes should also take place in the referenced array. 

You can use == operator to check the given slice is nill or not.
If you try to compare two slices with the == operator then it will give you an error. 

If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function: 
- Multi-Dimensional Slice
- Sorting of Slice 

Multi-D slice are just like the multi-D array, except that slice does not contain the size

In Go we are allowed to sort the elements present in the slice. The standard library of Go provides the sort package which contains different types of sorting methods for sorting the slices of ints, float64s, and strings. These functions always sort the elements available in slice in ascending order


// https://www.geeksforgeeks.org/slices-in-golang/
