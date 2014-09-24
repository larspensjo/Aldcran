#Principles
Use Aldcran as a control program, to optimize on parameters.

##Memory

The complete memory can also be used as indirect registers.

##Instructions

Only one instruction is needed, but it can do all arithmetic operations, flow control and memory movements. That way, one instruction can't mutate into another.

##Subroutines

Identify sequences of instructions that can be replaced by an already existing sub routine. This can be done based on hashing of the current subroutine, or hashing on the original sub routine, or both.

Spontaneous creation of sub routines, with at least N (5?) common instructions.

To determine if two subroutines can be merged, use the diff algorithm and compute how different they are.
