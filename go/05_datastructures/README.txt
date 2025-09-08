Implementing datastructures myself...

Done:
- LinkedList
- Queue
- Stack
- BinaryTree
- Set
  Uses a map of structs as per:
  https://emersion.fr/blog/2017/sets-in-go
	https://www.scaler.com/topics/golang/golang-set/

Todo:
- Generics (ie. not only ints)

Expectations for binary trees...

- Maximum depths for n elements:
            10 =>  4
           100 =>  9
         1_000 => 15
        10_000 => 18
       100_000 => 22
     1_000_000 => 26
    10_000_000 => 31
   100_000_000 => 38
 1_000_000_000 => 42

- Time to add n elements sequentially:
            10 =>           2µs
           100 =>          14µs
         1_000 =>         140µs
        10_000 =>       1_700µs
       100_000 =>      24_000µs
     1_000_000 =>     225_000µs
    10_000_000 =>   2_400_000µs
   100_000_000 =>  28_000_000µs
 1_000_000_000 => 450_000_000µs

- Time to remove n elements sequentially
            10 =>           1µs
           100 =>           7µs
         1_000 =>          70µs
        10_000 =>       1_000µs
       100_000 =>      13_000µs
     1_000_000 =>     130_000µs
    10_000_000 =>   1_400_000µs
   100_000_000 =>  16_000_000µs
 1_000_000_000 => 250_000_000µs

TODO: Times for random inserts and removals...

