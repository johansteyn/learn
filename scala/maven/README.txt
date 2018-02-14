Example of how to build a Scala "Hello World" program using Maven

Based on:
  http://fruzenshtein.com/scala-in-java-maven-project/

To compile:

  $ mvn clean scala:compile

This will create only .class files.

To run, either:

  $ cd target/classes
  $ scala Main

Or:

  $ scala -cp target/classes Main


To build the JAR:

  $ mvn clean package

To run the JAR, either:

  $ java -jar target/helloworld-1.0.0-jar-with-dependencies.jar

Or:

  $ scala target/helloworld-1.0.0-jar-with-dependencies.jar





