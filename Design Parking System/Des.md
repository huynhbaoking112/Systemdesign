# Design Parking System

**Difficulty:** Easy

## Problem Statement

Design a parking system for a parking lot.

The parking lot has three kinds of parking spaces:

- Big
- Medium
- Small

Each type of parking space has a fixed number of available slots.

## Implementation

Implement the `ParkingSystem` class:

### Constructor

```cpp
ParkingSystem(int big, int medium, int small)
```

Initializes the `ParkingSystem` object.

The number of slots for each parking space type is given as part of the constructor.

### Method

```cpp
bool addCar(int carType)
```

Checks whether there is an available parking space for a car of type `carType`.

`carType` can be one of the following:

- `1`: Big car
- `2`: Medium car
- `3`: Small car

A car can only park in a parking space of its own type.

Return:

- `true` if there is an available slot, then park the car in that slot.
- `false` if there is no available slot.

## Example 1

### Input

```text
["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
[[1, 1, 0], [1], [2], [3], [1]]
```

### Output

```text
[null, true, true, false, false]
```

### Explanation

```cpp
ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);

parkingSystem.addCar(1); // returns true because there is 1 available slot for a big car
parkingSystem.addCar(2); // returns true because there is 1 available slot for a medium car
parkingSystem.addCar(3); // returns false because there is no available slot for a small car
parkingSystem.addCar(1); // returns false because the big car slot is already occupied
```

## Constraints

- `0 <= big, medium, small <= 1000`
- `carType` is `1`, `2`, or `3`
- At most `1000` calls will be made to `addCar`
