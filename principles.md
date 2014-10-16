# Principles
Use Aldcran as a control program, to optimize on parameters.

## Memory

The complete memory can also be used as indirect registers.

## Instructions

Only one instruction is needed, but it can do all arithmetic operations, flow control and memory movements. That way, one instruction can't mutate into another.

## Subroutines

Identify sequences of instructions that can be replaced by an already existing sub routine. This can be done based on hashing of the current subroutine, or hashing on the original sub routine, or both.

Spontaneous creation of sub routines, with at least N (5?) common instructions.

To determine if two subroutines can be merged, use the diff algorithm and compute how different they are.

One idea is to start with easy problems, store the results in a library, and let Aldcran use this library.
That way, it may be possible to gradually address more complex problems.

## Merging of programs
It is no good to serialize a program into bytes, and merge byte strings.
The reason is, e.g., that a sequence of 5 bytes can be matched to another sequence of 4 bytes.
So a merge can result in a sequence of bytes that would not translate to a full instruction when it
is deserialized.
Instead, merging has to be done on complete instructions.
However, merging of individual istructions can be done bit-wise.
