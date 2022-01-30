wart1: Empty slice
------------------

A slice is similar to an array, but the size does not form part of its type
Using [n] or [...] declares an array, while using [] declares a slice

Since slices can grow and shrink, you can have an empty slice (but not an empty array).
The default value for a slice is nil
But Go's nil is not quite like C or Java's null
It's an identifier that represents the lack of a value.
Like literals and untyped constants, nil has no type so it can be assigned or
compared to values of different types.
But slices aren't comparable, ie. you can't use == or !=, but you can compare a slice to nil.

Creating a slice using an empty slice literal results in an empty but non-nil slice.

