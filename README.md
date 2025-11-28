# gontract
design-by-contract-like pre- and postconditions  for golang

## Design by Contract

Around 1988, [Bertrand Meyer](https://en.wikipedia.org/wiki/Bertrand_Meyer) has introduced the  ["Design by Contract" (DbC) paradigm](https://en.wikipedia.org/wiki/Design_by_contract) for object oriented programming (OOP) along with his programming language Eiffel (see [wikipedia](https://en.wikipedia.org/wiki/Eiffel_%28programming_language%29) and [eiffel.org](https://www.eiffel.org/) ).

The core of DbC consists  mainly of three concepts:

*  preconditions
* postconditions
* (class) invariants

while class invariants are indeed specific to OOP,  preconditions and postconditions
apply more generally to routines (functions/methods):

## Preconditions and Postconditions

**preconditions** are conditions that are  typically  imposed  on  the input of a funcion (parameters/arguments) and they  have to be satisfied in order for the function to be executed at all. It is the caller's responsibiliy to make sure that preconditions are satisfied.

**postconditions** are usually imposed on the result of a function  (return values)
and they have to be satisfied in order for the function to complete regularly.
It is the responsibility of the function implementation  to make sure postconditions are satisfied
if the function ran, i. e. if it was provided with valid input.

preconditions and ppostconditions form the essence of the **contract** between the
contractor (the function) and the contractee (the caller of the function).

Pre- and postconditions are essentially means for checking the input and result of a function, just a bit more intrinsically and idiomatically than with the usual conditional (if-else) 
constructs. -- if built into the programming language proper.

It is however an important point to note that preconditions and postconditions serve a very different purpose than usual input and result  validation:

While input validation and result validation are mostly meant to catch runtime errors of a program (like invalid user input), preconditions and postconditions are mostly  meant to catch bugs in the software:

*  preconditions catch bugs in the caller.
*  postconditions catch bugs in the fubction implementation.
  





The above  description should  have made it evident that the concepts of preconditions and postconditions are applicable to non-OOP languages as well.

## This project 

This project offers mechanisms for writing  preconditions and postconditions in golang.

Since golang does not have any such mechanisms   built into the language, 
they are implemented as separate functions intended  to be called at the beginning
and at the end of a function definition, respectively.

gontract implements conditions in an assertion-like manner using `panic()`.

This approach effectively prevents a function to run or complete at all when conditions are not satisfied.

The implementation of gontract was partly inspired by

[stone.code/assert](https://pkg.go.dev/gitlab.com/stone.code/assert)













