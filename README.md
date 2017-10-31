# Washout
Washout filter generates simulator's motions to simulate vehicle's motions.

![](https://travis-ci.org/shoarai/washout.svg?branch=master)

## Structure
![washoutstructure](https://cloud.githubusercontent.com/assets/5831786/26201462/0d654434-3c0e-11e7-8a44-e633e5290151.png)

## Installation
```sh
go get github.com/shoarai/washout
```
## Usage
```go
package main

import	"github.com/shoarai/washout/jaxfilter"

func main() {
    // Set the interval of processing in milliseconds.
    const interval = 10
    wash := jaxfilter.NewWashout(interval)

    // Pass vehicle's accelerations in meters per square second
    // and angular velocities in radians per second.
    position := wash.Filter(1, 1, 1, 3.14, 3.14, 3.14)

    // Position has simulator's displacements in meters and angles in radians.
    // type Position struct {
    //     X, Y, Z, AngleX, AngleY, AngleZ float64
    // }
}
```

## Preferences
[Evaluation of motion with washout algorithm for flight simulator using tripod parallel mechanism](http://ieeexplore.ieee.org/document/6484612/)
