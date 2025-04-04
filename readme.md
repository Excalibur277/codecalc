# Code Calc

CodeCalc is a simple signed 64-bit integer based language for linux x86-64 targets. It ignores whitespace and terminates lines with semi-colons (;). The language also contains arrays, but not as first class entities.

## Quickstart

To compile and run a .calc file:

```
./codecalc -run [pathto.calc] [output_binary]
i.e.
./codecalc -run ./speedtests/speed.calc ./speed
```

## Comments

\# A hash starts a line comment that continues until a newline

\` Back ticks are used to start and close multiline comments

## Expressions

Expressions consist of literals, variables and operation expressions. Variables must be assigned before use.

## Operations

There is a basic order of operations:

```
(,)
Unary: +,-,! (Positive, Negative, Not)
Binary: \*,/,% (Multiply, Divide, Modulus)
Binary: +,- (Add, Subtract)
Binary: <=,<,>=,> (Less than or Equal, Less than, Greater than or Equal, Greater than)
Binary: ==,!= (Equal, Not Equal)
Binary: &,| (And, Or)
```

Binary operators require a value on either side, while unary operators precede a single value.

Comparision operators evaluate to 0 if false and 1 if true. Division or modulus by zero causes the execution to fail, returning exit code 1.

& evaluates to 0 if either the value are 0 and 1 otherwise.

| evaluates to 0 if both the value are 0 and 1 otherwise.

Not evaluates to 1 if the value is 0 and 0 otherwise.

## Statements

There are 5 top level statements:

Statements can include sub codeblocks. Variables assigned for the first time in the code blocks will not exist outside of it, however the value of variables reassigned within them remain.

### Print Statement

An expression terminated by a semi-colon. This will be evaluated into a single number that will be printed to stdout.

```
1 + 1; # Evaluates and prints 2 to the stdout
```

### Array Declaration Statement

Arrays can be dynamically allocated to an identifier. They will be deallocated once the scope ends.

Arrays are not first class entities and cannot be passed around, only indexed and assigned to.

Arrays are 0 indexed, and are bounds checked in runtime, so cannot be used as a point equivilent (the language does not have pointers).

Arrays are zeroed out initially.

```
array myArray[5 + 12 * n];
a = myArray[8];
```

### Assign Statement

Variable identifiers must start with a letter or underscore and can be followed by any number of letter, underscores and numbers as defined by the Unicode standard.

```
myA = 5 + 2; # Sets myA to 7
myArray[0] = 5 + 2; # Sets the first value of myArray to 7
```

Variables can be reassigned any number of times.

### If Statement

If the expression provided to the if statement evaluates to anything other than 0 the subsequent code block will execute, otherwise if an there is a subsequent else or else if statement they will execute.

As a control statement semicolons are not required to terminate, rather the block should be opened and closed with curly braces.

The if statement can be chained with else if and else statements.

```
# If must be the first statement in the if-elseif-else chain.
if a > 2 {
    # Actions if a > 2
}
# Optional subsequent statements
else if a > 1 {
    # Actions if a > 1 and not a > 2
}
# Else must be the final statement in the if-elseif-else chain.
else {
    # Actions if not a > 1 and not > 2
}

# A solo if statement is valid
if a > 2 {
    # Actions if a > 2
}


if a > 2 {
    # Actions if a > 2
    b = 5;
}
b; # Does not compile as b is not accessible outside of the block it is first assigned in.
```

### While Statement

The while statement expression is evaluated the same as the if statement, however on the completion of the codeblock the while statement is repeated. This means while statements repeat the code block until the expression evaluates to 0.

As a control statement semicolons are not required to terminate, rather the block should be opened and closed with curly braces.

e.g.

```
a = 10;
while a > 2 {
    a = a - 1;
}
a; Prints 2
```

Additionally there are 2 more statements valid within the while loop.

```
continue; # Terminates this execution of the while codeblock and immediately goes to evaluate the while expression.
break; # Terminates this execution of the while codeblock and immediately ends the while statement without evaluation of the expression.
```

e.g.

```
a = 10;
b = 0;
while a > 2 {
    a = a - 1;
    b = b + 1;
    if b > 5 {
        break;
    }
}
a; Prints 4
```

## Compilation

### Run CodeCalc

To build the compiler run:

```

go build -o codecalc main.go

```

To compile a file run:

```

./codecalc [filename] [outfilename]

```

To run the binary run:

```

./[outfilename]

```

You can view the produced fasm code by adding the flag -fasm=[outputfasmfilename]

Additional debug info can be viewed with the flag -debug

Additionally the binary can be immediatly ran with the flag -run

### Parser

The parser uses antlr4 to parse the file and creates an abstract syntax tree through listener methods on the default tree walker.

### Backend

The abstract synatx tree is used to generate assembly for the fasm assembler. Fasm targets x86 instruction sets, however CodeCalc targets the x86-64 bit variant.

The assembly is not optimal, with high stack usage and no consideration for optimal register allocation.
