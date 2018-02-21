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

Build JAR (with all dependent JARs):

  $ mkdir project
  $ cp assembly.sbt project
  $ sbt assembly

Run:

  $ scala target/scala-2.12/pingpong.jar

  or

  $ java -jar target/scala-2.12/pingpong.jar



