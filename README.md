# gostats : efficient statistics array for golang

Gostats is a package that attempts to create an efficient float64 array for statistics aggregation. The main goal being to optimise elemental operations (mean, variance, skewness, quantiles) to an O(1) complexity and piggy-back off said complexity to optimise higher order statistics operations. 

The `Arrayf64` is ultimately a wrapper around a `[]float64` which, at the cost of an `O(nlog(n))` insert complexity and the order of said insert, optimises the computation of all statistics.

## Usage


```go
a := Arrayf64{}
a.Init(Optionf64{
    Degree:    4,       // Degree of summation to store
    Harmonic:  true,    // Optimise harmonic mean computation
    Geometric: true,    // Optimise geometric mean computation
    // The more optimisation and the higher the degree chosen
    // the larger the array size (as an indicator, 1 extra float64
    // added to the array wrapper for each option)
})

a.Insert(5)
a.InsertSlice([]float64{5.2,6,54.,21,.22})

a.Summary()             // Prints statistics with O(n) garantee 
```


## Reference

### Statistics operations

|        operation        |   complexity  |        algorithm / method       |
|:-----------------------:|:-------------:|:-------------------------------:|
|          Mean()         |      O(1)     |         iterative insert        |
|         Stddev()        |      O(1)     |         iterative insert        |
|          Min()          |      O(1)     |           sorted array          |
|          Max()          |      O(1)     |           sorted array          |
|  Quantile() or Median() |      O(1)     |           sorted array          |
|      GeometricMean()    |      O(1)     |         iterative insert        |
|       HarmonicMean()    |      O(1)     |         iterative insert        |
|        Entropy()        |      O(n)     |                                 |
|      Skewness(Yule)     |      O(1)     |                                 |
| Skewness(PearsonSecond) |      O(1)     |                                 |
|  Skewness(PearsonFirst) | O(card(mode)) |                                 |
|        Kurtosis()       |      O(n)     |                                 |
|      ShapiroWilk()      |      O(n)     | ROYSTON, Patrick. Remark AS R94 |

### Elementary array operations

|  operation  |  complexity |
|:-----------:|:-----------:|
|    Insert   |  O(nlog(n)) |
| InsertSlice | O(mnlog(n)) |
|      At     |     O(1)    |
|    Remove   |  O(nlog(n)) |
|   DeepCopy  |     O(n)    |
|    Center   |     O(n)    |
|    Reduce   |     O(n)    |
|  apply(op)  |  O(n)*O(op) |
