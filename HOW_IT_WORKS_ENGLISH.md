# Algorithm Description

## General Description
The main principle of the algorithm is building a tree. Each expression consists of a left and right part. If a value is not yet calculated, a channel is passed to it, where the value will be transmitted once the expression is calculated. When the application starts, all expressions are launched in goroutines. If an expression lacks a variable for calculation, it remains in a blocked state. As soon as a calculated value is received in the channel, the expression is evaluated. Expressions not used in the tree are not calculated.

## Interfaces
* **SentenceArgument** - argument interface

## Entities
* **Number** - a wrapper for int64. It's needed to implement the SentenceArgument interface. This is necessary because an argument can be either a string or a number.
* **Variable** - a variable designed to store values. It can be printed, plus it implements an observer pattern - as soon as a value is calculated, the variable notifies all other expressions about it.
* **Sentence** - an expression. It can be calculated. Contains references to the left and right parts.