Separate dirs for Java and Scala implementations.

Java
====

Build and run with Maven:

  $ mvn clean compile exec:exec

Build & package with Maven, and run with Java on the command line:

  $ mvn clean package
  $ java -jar target/pingpong-1.0.jar


Scala
=====

Build and run with SBT:

  $ sbt run


