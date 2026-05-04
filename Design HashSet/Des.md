# 705. Design HashSet

**Difficulty:** Easy

## Description

Design a HashSet without using any built-in hash table libraries.

Implement the `MyHashSet` class:

- `void add(key)`: Inserts the value `key` into the HashSet.
- `bool contains(key)`: Returns whether the value `key` exists in the HashSet.
- `void remove(key)`: Removes the value `key` from the HashSet. If `key` does not exist, do nothing.

## Example 1

**Input**

```text
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
```

**Output**

```text
[null, null, null, true, false, null, true, null, false]
```

**Explanation**

```java
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return true
myHashSet.contains(3); // return false, not found
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return true
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return false, already removed
```
