println("For comprehension")

val list = List(1, 2, 3)
val doubles = for (x <- list) yield x * 2
println("list: " + list)
println("doubles: " + doubles)
// The above is equivalent to:
val doubles2 = list.map(x => x * 2)
println("doubles2: " + doubles2)

val list2 = List(11, 12)
val sums = for (x <- list; y <- list2) yield x + y
println("list2: " + list2)
println("sums: " + sums)
// The above is equivalent to:
val sums2 = list flatMap (x => list2 map (y => x + y))
println("sums2: " + sums2)

val list3 = List(1, 2)
val products = for (x <- list; y <- list2; z <- list3) yield x * y * z
println("list3: " + list3)
println("products: " + products)
// The above is equivalent to: 
val products2 = list flatMap (x => list2 flatMap (y => list3 map (z => x * y * z)))
println("products2: " + products2)
// Note that all except the last use flatMap. The last uses map.
// Good explanation here:
//   https://stackoverflow.com/questions/14598990/confused-with-the-for-comprehension-to-flatmap-map-transformation


// https://scala-tuts.net/2014/07/08/the-versatile-scala-for-comprehension
// Assuming both host and port are not None, you will get a valid InetSocketAddress.
// If either one is None, addr will also be None
import java.net.InetSocketAddress
val host: Option[String] = Some("localhost")
//val host: Option[String] = None
val port: Option[Int] = Some(8080)
val addr: Option[InetSocketAddress] = for {
    h <- host
    p <- port
} yield new InetSocketAddress(h, p)
println("addr: " + addr)
// The above is shorthand for:
val addr2 = host flatMap { 
  h => port map {
    p => new InetSocketAddress(h, p)
  }
}
println("addr2: " + addr2)
// Note that the above 2 are possible because host and port are "Options".
// If host were a String, then it would iterate over the chars,
// and try to construct instances of InetSocketAddress with chars,
// for which no constructor exists.
// Instead, it iterates over Options, which contain either one element,
// of type Some or one of type None, ie. it "iterates" over a single value.
// ie. variables h and p are effectively the strings inside of Options host and port:
println("host: " + host)
println("port: " + port)
val addr3 = host flatMap { 
  h => {
    println("h: " + h)
    port map {
      p => {
        println("p: " + p)
        new InetSocketAddress(h, p)
      }
    }
  }
}
println("addr3: " + addr3)
// In summary: The code iterates over two sequences containing only one element each,
// (some string and some int) and yields a sequence that also contains only one element
// (some address).

