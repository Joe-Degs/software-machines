## Implementing Machines in Software
This is a code along of Eleanor McHugh's talk on building components of virtual
machines in software(how vms are built). Basically abstract virtual machines. I do not really know anything about stuff like this, but reading and listening to stuff about the intricate
parts of machines always fascinates me, so i'm trying out some of the code to
understand more of it and how all this things work.

The codes are in c and go, go I am a little bit comfortable with but c i do not
know anything about. But because go is close to c i feel like i know some things
about it. So i'm testing the c code too.

Lets start with the components of computer machines;

### Memory(Heap and Stack)
There are few models of memory around. Most i have not heard of before. One
i actually know is the Von-Neumann memory model but turns out there is a whole
lot more of memory. The basic function of the memory is storing data and
instructions.
- Von Neumannn model
This one is just like an array. One big contigous chunk of addressable space,
separated/partitioned so we have different parts of the chunk for stack, heap
anbd the rest. The basic one everybody who hasnt taken a computer science course
on the subject knows about. 
- Harvard model
Separate ROM and RAM, this means instructions and data are not stored in the
same contigous chunk of memory. This potentially prevents buffer overruns
attacks. so code will be stored in the ROM and data in the RAM.
- Forth model
Complex model with two diffent stacks, not really a standard heap but some
lookup map. And some other complicated details
- Hybrid model
This one is combines both forth and von neumann model to create a more
complicated model that have probably more than two stacks.
#### Heap
- the heap is word aligned, byte addressable chunk of contigous memory. It is
  basically the same accross all the memory models. And the implementation is go
might not be that different the implementation in c
#### Stack
- The stack is a slightly different beast, depending on the memory model, it
  might be trivial to implement. The implementation will be off the type linux
processes use.
- There are a few types of stacks. We have the `spaghetti/cactus stack` and the
  `array stack`(regular linux process stack)
- The stack is fixed array of memory, can be grown and shrinked and has
  a program counter that is an offset into the array.
- Whereas the array stack is just one contigous chunk, the cactus/spaghetti
  stack is list of memory cells linked together, basically a linked list. It is
a functional data type especially if you make it immutable.
