Method sets
============

Method sets determine what methods attach to a TYPE. It is exactly what the name says: `What is the set of methods for a given type? That is its method set.`

### IMPORTANT: 
    “The method set of a type determines the INTERFACES that the type implements.....”
                    ~ golang spec


#### a NON-POINTER RECEIVER
        works with values that are POINTERS or NON-POINTERS.

#### a POINTER RECEIVER
        only works with values that are POINTERS.

### Receivers       Values
   -----------------------
    (t T)           T and *T
    (t *T)          *T
