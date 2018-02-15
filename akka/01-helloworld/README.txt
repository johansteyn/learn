Separate dirs for Java and Scala implementations.

Java
====

Build and run with Maven:

  $ mvn clean compile exec:exec

Build & package with Maven, and run with Java on the command line:

  $ mvn clean package
  $ java -jar target/helloworld-1.0.jar

Note: 
The alternative implementation extends AbstractLoggingActor.
To run it, use:

  $ java -classpath target/helloworld-1.0.jar HelloLogger


Scala
=====

Build and run with SBT:

  $ sbt run

NOTE: SBT leaves the terminal in some weird state, so rather use:

  $ sbt run; stty sane

TODO: Figure out how to make SBT less verbose...

To run on the command line, use the "sbt-assembly" plugin to build a JAR containing all the dependencies.
- Create a file called project/assembly.sbt with single line:
    addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "0.14.5")
- Build and package with:
    $ sbt assembly
- Run using:
    $ scala target/scala-2.12/helloworld.jar
  Or even:
    $ java -jar target/scala-2.12/helloworld.jar


