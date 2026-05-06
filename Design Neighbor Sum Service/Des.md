# Design Neighbor Sum Service

**Difficulty:** Easy

## Problem Statement

You are given an `n x n` 2D array `grid` containing distinct elements in the range `[0, n^2 - 1]`.

Implement the `NeighborSum` class.

## API

```java
NeighborSum(int[][] grid)
```

Initializes the object.

```java
int adjacentSum(int value)
```

Returns the sum of elements that are adjacent neighbors of `value`, meaning the elements directly:

- above
- below
- left
- right

```java
int diagonalSum(int value)
```

Returns the sum of elements that are diagonal neighbors of `value`, meaning the elements at:

- top-left
- top-right
- bottom-left
- bottom-right

## Example 1

**Input**

```text
["NeighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"]
[[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]
```

**Output**

```text
[null, 6, 16, 16, 4]
```

**Explanation**

- The adjacent neighbors of `1` are `0`, `2`, and `4`.
- The adjacent neighbors of `4` are `1`, `3`, `5`, and `7`.
- The diagonal neighbors of `4` are `0`, `2`, `6`, and `8`.
- The diagonal neighbor of `8` is `4`.

## Example 2

**Input**

```text
["NeighborSum", "adjacentSum", "diagonalSum"]
[[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15], [9]]
```

**Output**

```text
[null, 23, 45]
```

**Explanation**

- The adjacent neighbors of `15` are `0`, `10`, `7`, and `6`.
- The diagonal neighbors of `9` are `4`, `12`, `14`, and `15`.

## Constraints

- `3 <= n == grid.length == grid[0].length <= 10`
- `0 <= grid[i][j] <= n^2 - 1`
- All `grid[i][j]` are distinct.
- `value` in `adjacentSum` and `diagonalSum` will be in the range `[0, n^2 - 1]`.
- At most `2 * n^2` calls will be made to `adjacentSum` and `diagonalSum`.
