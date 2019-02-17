// https://www.journaldev.com/9585/scala-variances-covariant-invariant-contravariant

object Main extends App {
  class Animal { print("Animal.") }
  class Mammal extends Animal { print("Mammal.") }
  class Cat extends Mammal { print("Cat.") }
  class Burmese extends Cat { print("Burmese.") }
  class Ginger extends Cat { print("Ginger.") }
  class Dog extends Mammal { print("Dog.") }
  class Husky extends Dog { print("Husky.") }
  class Boxer extends Dog { print("Boxer.") }
  class Reptile extends Animal { print("Reptile.") }
  class Lizard extends Reptile { print("Lizard.") }
  class Snake extends Reptile { print("Snake.") }

  //                       ____ Animal_____
  //                      /                \
  //            ___Mammal____               Reptile
  //           /             \              /     \
  //        Cat              Dog       Lizard     Snake
  //       /  \             /  \
  // Burmese  Ginger    Husky  Boxer

  val mammal = new Mammal
  println()
  val dog = new Dog
  println()
  val husky = new Husky
  println()
  val boxer = new Boxer
  println()

  println("=======Covariance=======")
  // Pet does not form part of the "Animal" Class hierarchy.
  // It uses a "has-a" rather than an "is-a" relationship.
  // It is "generic" in that it can hold a value of any type - not just Animal
  // (it just happens to name the parameter "animal").
  // The type parameter T means that when we declare a Pet to hold type T,
  // then the parameter we pass MUST be of type T - nothing else.
  // But if we use +T instead, then the parameter can be of type T or any subclass of T.
  class Pet[+T](val animal:T) { print("Pet[" + animal.getClass + "]") }

  val m:Pet[Mammal] = new Pet[Mammal](mammal)
  println()
  val fido:Pet[Dog] = new Pet[Dog](dog)
  println()
  val bruno:Pet[Boxer] = new Pet[Boxer](boxer)
  println()
  val sky:Pet[Husky] = new Pet[Husky](husky)
  println()
  // Pet[Dog] is a subclass of Pet[Mammal], because Dog extends Mammal
  // Similarly, Pet[Boxer] and Pet[Husky] are subclasses of Pet[Dog]
  // and hence they are also subclasses of Pet[Mammal]

  //          Pet[Mammal]
  //               |
  //            Pet[Dog]
  //           /       \
  //  Pet[Boxer]       Pet[Husky]


  println("------------------------")
  // A dog owner can have a Pet of type Dog or any of it's subclasses, 
  // thanks to the +T type parameter of Pet above.
  class DogOwner(val dog:Pet[Dog]) { print("Dogowner") }

  // Works when Pet has type parameter T or +T
  // In this case, fido has exactly the required type: Pet[Dog]
  val dogOwner = new DogOwner(fido)
  println()
  // Only works when Pet has type parameter +T
  // In this case, sky has sub-type: Pet[Husky]
  val huskyOwner = new DogOwner(sky)
  println()
  // Doesn't work at all, because m has type: Pet[Mammal]
  // which is above type Pet[Dog] in the hierarchy.
//  val mammalOwner = new DogOwner(m)
//  println()

  println("=====Contravariance=====")
  class Wild[-T] { print("Wild.") }
  class WildMammal extends Wild[Mammal] { print("WildMammal.") }
  class WildDog extends Wild[Dog] { print("WildDog.") }
  class WildHusky extends Wild[Husky] { print("WildHusky.") }
  class WilderDog extends WildDog { print("WilderDog.") }
  class WildCat extends Wild[Cat] { print("WildCat.") }


  def shoot(feral: Wild[Dog]) {
    println("Shot feral of type: " + feral.getClass)
  }

  val bigfoot = new WildMammal
  println()
  val dingo = new WildDog
  println()
  val wolf = new WildHusky
  println()
  val wilder = new WilderDog
  println()

  // We can shoot an instance of WildDog, because it "is-a" Wild[Dog]
  shoot(dingo)
  println()
  // We can also shoot an instance of WildMammal, because Wild[Mammal] is a superclass of Wild[Dog]
  shoot(bigfoot)
  println()
  // But we cannot shoot an instance of WildHusky, because it is a sub-class of Wild[Dog]
  //shoot(wolf)
  //println()
  // Yet we can shoot an instance of WilderDog, because...
  shoot(wilder)
  println()

//  // TODO...
//  val mammalOwner = new DogOwner(mammal)
//  println()

}


