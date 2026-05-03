# 232. Implement Queue using Stacks

**Difficulty:** Easy

## Problem Statement

Implement a first-in-first-out (FIFO) queue using only two stacks.

The implemented queue should support all the functions of a normal queue:

- `push`
- `peek`
- `pop`
- `empty`

## Class Definition

Implement the `MyQueue` class:

| Method | Description |
| --- | --- |
| `void push(int x)` | Pushes element `x` to the back of the queue. |
| `int pop()` | Removes the element from the front of the queue and returns it. |
| `int peek()` | Returns the element at the front of the queue. |
| `boolean empty()` | Returns `true` if the queue is empty, otherwise returns `false`. |

## Notes

- You must use only standard stack operations.
- Valid stack operations include:
  - push to top
  - peek from top
  - pop from top
  - get size
  - check if empty
- Depending on your language, the stack may not be supported natively.
- You may simulate a stack using a list or deque, as long as you only use standard stack operations.

## Example 1

### Input

```text
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
```

### Output

```text
[null, null, null, 1, 1, false]
```

### Explanation

```java
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek();  // return 1
myQueue.pop();   // return 1, queue is [2]
myQueue.empty(); // return false
```

## Constraints

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `peek`, and `empty`.
- All calls to `pop` and `peek` are valid.

## Follow-up

Can you implement the queue so that each operation has **amortized O(1)** time complexity?

In other words, performing `n` operations should take overall `O(n)` time, even if one of those operations may take longer.
