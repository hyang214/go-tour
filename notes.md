# A Tout of Go

## 0x00. Brief

A simple class for go from Go official web.

[SITE](https://tour.golang.org/list)

With this class, may I could learn the basic use and knowledge of Go language.

## 0x01. Welcome

### 1. Hello world

Just write the classical program in go.

[Code](Welcome/Hello.go)

### 2. Go local

No need to read.

### 3. Go offline (optional)

Skip.

### 4. The Go Playground

Go program run in Go playground.

And introduce the ```time``` package.

[Code](Welcome/Time.go)

### 5. Congratulations

Well, the and of Welcome.

## 0x02. Basics

The starting point, learn all the basics of the language.

Declaring variables, calling functions, and all the things you need to know before moving to the next lessons.

### 1. Packages, variables and functions

#### a. Packages

The entrance of program in the ``main`` function in the go file with package ```main```

like
``` go

package main

func main() {
  // do soemthing
}

```

The package name is the same as the last element of the import path.

[Code](Basics/Math.go)

#### b. Imports

use ```import``` to introduce package into program.

The two methods are both work well and equal.

+ I.
```
import "fmt"
import "math"
```

+ II.
```
import (
    "fmt"
    "math
)
```

[Code](Basics/Imports.go)

#### c. Exported names

All name is exported if it begins with a capital letter, include function and variable.

[Code](Basics/Exported.go)

#### d. Functions

A function can take zero or more arguments and type comes after the variable name.

[Code](Basics/Functions.go)

#### e. Functions continued

To simplify the code, if your function has two or more consecutive named function parameters share a type, 
you only need define type of the last one.

Such as,
```
func add(x, y int) {
	return x + y
}
```
is equal with
```
func add(x int, y int) {
	return x + y
}
```

[Code](Basics/FunctionsContinued.go)

#### f. Multiple results

```=``` value assign
```:=``` variable define, infer and assign

A function can return any number of results.

[Code](Basics/MultipleResults.go)

#### g. Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

[Code](Basics/NamedReturn.go)

#### h. Variables

var can be defined at package, function level.

variable will be init with default value.

int: 0
bool: false
string: a blank space, just like " "

[Code](Basics/Variables.go)

#### i. Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

[Code](Basics/initializerVariables.go)

#### j. Short variable declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

[Code](Basics/Assignment.go)

#### k. Basic types

All basic type in Go is following:

+ bool
+ string
+ int  int8  int16  int32  int64
+ uint uint8 uint16 uint32 uint64 uintptr
+ byte 
+ rune
+ float32 float64
+ complex64 complex128

Attention:
+ uintptr is the integer type use for store pointer
+ byte is alias for uint8
+ rune is alias for int32, represents a Unicode code point
+ The length of int, uint and uintptr type depend on the system, usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems

[Code](Basics/BasicType.go)

#### l. Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

+ 0 for numeric types
+ false for the boolean type
+ "" (the empty string) for strings

[Code](Basics/Zero.go)

#### m. Type conversions

The expression T(v) converts the value v to the type T.

such as:

```
var i int = 42
f := float64(i)
var u uint = uint(f)
```

[Code](Basics/TypeConversions.go)

#### n. Type inference

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type
When the right hand side of the declaration is untyped, go will infer the type.

[Code](Basics/TypeInference.go)

#### o. Constants

Constants can be any basic types, but use ```const``` keyword to define.

The only way to define a constant variable is:

```
const i = 42
const str = "sun"
```

[Code](Basics/Constants.go)

#### p. Numeric Constants

Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

### 2. Flow control statement: for, if else, switch and defer

Learn how to control the flow of your code with conditionals, loops, switches and defers.

#### a. For

Simple!

[Code](Basics/For.go)

#### b. For continued

Simple!

[Code](Basics/ForContinued.go)

#### c. For is Go's "while"

use ```for``` instead of ```while```.

[Code](Basics/While.go)

#### d. Forever

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

[Code](Basics/Forever.go)

#### e. If

How if work, I think you know.

[Code](Basics/If.go)


#### f. If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

[Code](Basics/IfStatement.go)

#### g. If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

[Code](Basics/Else.go)

#### h. Exercise: Loops and Functions

Show [Code](Basics/LoopAndFunction.go)

#### i. Switch

DO NOT Need ```break``` in each case.
Go's switch cases need not be constants, and the values involved need not be integers.

Show [Code](Basics/Switch.go)

#### j. Switch evaluation order

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

Show [Code](Basics/SwitchEvaluationOrder.go)

#### k. Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

Show [Code](Basics/NoConditionSwitch.go)

#### l. Defer

Code show how ```defer``` work.

Show [Code](Basics/Defer.go)

#### m. Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

Show [Code](Basics/Defer.go)

### 3. More types: structs, slices, and maps

#### a. Pointers

Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

The & operator generates a pointer to its operand.

Show [Code](Basics/Pointers.go)

#### b. Structs

A struct is a collection of fields.

Show [Code](Basics/Structs.go)

#### c. Struct Fields

Struct fields are accessed using a dot.

Show [Code](Basics/StructFields.go)

#### d. Pointers to structs

Struct fields can be accessed through a struct pointer.

And ```(*p).X``` is equals woth ```p.X```

Show [Code](Basics/PointersOfStructs.go)

#### e. Struct Literals

struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

Show [Code](Basics/StructLiterals.go)

#### f. Arrays

The type [n]T is an array of n values of type T.

Show [Code](Basics/Arrays.go)

#### g, Slices

```slice``` is used to get a part of an array.

```a[low: high]``` generates a new array with length (high - low), and with oppose elements in the array ```a```.

Show [Code](Basics/Slices.go)

#### h. Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

Show [Code](Basics/SlicesReference.go)

#### i. Slice literals

Slice literals
A slice literal is like an array literal without the length.

This is an array literal:
```
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:

```
[]bool{true, true, false}
```

Show [Code](Basics/SliceLiterals.go)

#### j. Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

Show [Code](Basics/SliceDefault.go)

#### k. Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

Show [Code](Basics/SliceLengthAndCapacity.go)

#### l. Nil slices

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

Show [Code](Basics/NilSlice.go)

#### m. Creating a slice with make

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

```
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to make:

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

Show [Code](Basics/MakeSlice.go)

#### n. Slices of slices

Slices can contain any type, including other slices.

Show [Code](Basics/SlicesOfSlices.go)

#### o. Appending to a slice

```
func append(s []T, vs ...T) []T
```

Append a new element of T into the tail of the specific slice

Show [Code](Basics/SliceAppend.go)

#### p. Range

The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

Show [Code](Basics/SliceRange.go)

#### q. Range continued

You can skip the index or value by assigning to _.

```
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```
for i := range pow
```

Show [Code](Basics/SliceRangeContinued.go)

#### r. Exercise: Slices

Just coding.

Show [Code](Basics/ExerciseSlices.go)

#### s. Maps

A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.

Show [Code](Basics/Maps.go)

#### t. Map literals

Map literals are like struct literals, but the keys are required.

Show [Code](Basics/MapLiterals.go)

#### u. Map literals continued

If the top-level type is just a type name, you can omit it from the elements of the literal.

Show [Code](Basics/MapLiterals.go)

#### v. Mutating Maps

Some simple mutation of map, code show the usage.

Show [Code](Basics/MutatingMaps.go)

#### w. Exercise: Maps

Show [Code](Basics/ExerciseMaps.go)

#### x. Function values

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

Show [Code](Basics/FunctionValues.go)

#### y. Function closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

Closure belong to the function instance, if a type of function has two function instance, each instance has its own closure, and does not share with other.

Show [Code](Basics/FunctionClosures.go)

#### z. Exercise: Fibonacci closure

Coding!

Show [Code](Basics/ExerciseFibonacciClosure.go)

## 0x03. Methods and interfaces

Learn how to define methods on types, how to declare interfaces, and how to put everything together.

### 1. Methods and Interfaces

#### a. Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

Show [Code](MethodAndIntefaces/Methods.go)

#### b. Methods are functions

Remember: a method is just a function with a receiver argument.

Show [Code](MethodAndIntefaces/MethodsAreFunctions.go)

#### c. Methods continued

You can declare a method on non-struct types, too.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

Show [Code](MethodAndIntefaces/MethodsContinued.go)

#### d. Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Show [Code](MethodAndIntefaces/PointerReceivers.go)

#### e. Pointers and functions

Another way to change the value of argument.

Show [Code](MethodAndIntefaces/PointersAndFunctions.go)

#### f. Methods and pointer indirection

When the function has a pointer, you must fill it with a pointer. 

But when there is a pointer receiver, go will transform a value ```v``` to its pointer by ```&v``` automatically.

Show [Code](MethodAndIntefaces/MethodsAndPointerIndirection.go)

#### g. Methods and pointer indirection (2)

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

Show [Code](MethodAndIntefaces/MethodsAndPointerIndirection.go)

#### h. Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

Show [Code](MethodAndIntefaces/ChoosingReceiver.go)

#### i. Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Show [Code](MethodAndIntefaces/Interfaces.go)

#### j. Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

Show [Code](MethodAndIntefaces/InterfacesImplementedImplicitly.go)

#### k. Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

Show [Code](MethodAndIntefaces/InterfacesValue.go)

#### l. Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

Show [Code](MethodAndIntefaces/InterfacesNilValue.go)

#### m. Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

Show [Code](MethodAndIntefaces/nil-interface-values.go)

#### n. The empty interface

The interface type that specifies zero methods is known as the empty interface:

```
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

Show [Code](MethodAndIntefaces/empty-interface.go)

#### o. Type assertions

A type assertion provides access to an interface value's underlying concrete value.

```
t := i.(T)
```

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```
t, ok := i.(T)
```

If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

Show [Code](MethodAndIntefaces/type-assertions.go)

#### p. Type switches

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```
switch v := i.(type) {
case T:
// here v has type T
case S:
// here v has type S
default:
// no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

Show [Code](MethodAndIntefaces/type-switch.go)

#### q. Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

```
type Stringer interface {
    String() string
}
```
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

Show [Code](MethodAndIntefaces/stringers.go)

#### r. Errors

Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

```
type error interface {
    Error() string
}
```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
A nil error denotes success; a non-nil error denotes failure.

Show [Code](MethodAndIntefaces/errors.go)

#### s. Exercise: Errors

Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

```
type ErrNegativeSqrt float64
```

and make it an error by giving it a

```
func (e ErrNegativeSqrt) Error() string
```

method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

Show [Code](MethodAndIntefaces/exercise-errors.go)

#### t. Readers

The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The io.Reader interface has a Read method:

```
func (T) Read(b []byte) (n int, err error)
```

Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

The example code creates a strings.Reader and consumes its output 8 bytes at a time.

Show [Code](MethodAndIntefaces/readers.go)

#### u. Exercise: Readers

Just do it.

Show [Code](MethodAndIntefaces/exercise-readers.go)

#### v. Exercise: rot13Reader

Coding, coding, coding!

Show [Code](MethodAndIntefaces/coding.go)

#### w. Images

Package image defines the Image interface:

```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

(See the documentation for all the details.)

The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package

Show [Code](MethodAndIntefaces/image.go)

#### x. Exercise: Images

Coding

Show [Code](MethodAndIntefaces/exercise-image.go)

## 0x04 Concurrency

Go provides concurrency features as part of the core language.

This module goes over goroutines and channels, and how they are used to implement different concurrency patterns.

### 1. Concurrency

Go provides concurrency constructions as part of the core language. This lesson presents them and gives some examples on how they can be used.

#### a. Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

starts a new goroutine running

```
f(x, y, z)
```

The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.

Show [Code](Concurrency/goroutine.go)

#### b. Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
          // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

Show [Code](Concurrency/channels.go)

#### c. Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

```
ch := make(chan int, 100)
```

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

Show [Code](Concurrency/buffered-channels.go)

#### d. Range and Close

A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```
v, ok := <-ch
```

ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Show [Code](Concurrency/range-and-close.go)

#### e. Select

The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Show [Code](Concurrency/select.go)

#### f. Default Selection

The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

Show [Code](Concurrency/default-selection.go)

#### g. Exercise: Equivalent Binary Trees

Coding exercise.

Show [Code](Concurrency/exercise-equivalent-binary-tree.go)

#### h. sync.Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

```
Lock
Unlock
```

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.

We can also use defer to ensure the mutex will be unlocked as in the Value method.

Show [Code](Concurrency/mutex-conter.go)

#### i. Exercise: Web Crawler

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

Show [Code](Concurrency/exercise-web-crawler.go)

#### j.