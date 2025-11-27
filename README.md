# gontract
design-by-contract-like pre- and postconditions  for golang


Bertrand Meyer has introduced the  "design by contract" (DbC) paradigm for object oriented programming (OOP) along with his language Eiffel.

DbC consists  mainly of three concepts:

*  preconditions
 * postconditions
* (class) invariants

while class invariants are indeed specific to OOP,  pre- and postconditions
apply more generally to routines (functions/methods):

**preconditions** are conditions that are  usually  imposed  on  the input of a funcion (parameters/arguments) and they  have to be satisfied in order for the function to be executed at all. It is the caller's responsibiliy to make sure that preconditions are satisfied.

**postconditions** are usually imposed on the result of a function  (return values)
and they have to be satisfied in order for the function to complete regularly.
It is the responsibility of the function to make sure postconditions are satisfied
if the function ran, i. e. if it was provided with valid input.

preconditions and ppostconditions form the essence of the **contract** between the
contractor (the function) and the contractee (the caller of the function).

Pre- and postconditions are essentially means for checking the input and result of a function, just a bit more intrinsically and idiomatically than with the usual conditional (if-else) constructs. -- if built into the programming language proper.

This description makes it plausible that the concept preconditions and postconditions are applicable to non-OOP languages as well.

This project offers mechanisms for writing  preconditions and postconditions in golang.

Since golang does not have any such mechanisms   built into the language, 
they are implemented as separate functions intended  to be called at the beginning
and at the end of a function definition, respectively.

gontract implements conditions in an assertion-like manner using `panic()`.

The implementation of gontract was partly inspired by

https://pkg.go.dev/gitlab.com/stone.code/assert













