name := "Word Counter"

version := "1.0"

scalaVersion := "2.12.3"

resolvers += "Typesafe Repository" at "http://repo.typesafe.com/typesafe/releases/"

libraryDependencies += "com.typesafe.akka" % "akka-actor_2.12" % "2.5.3"

assemblyJarName in assembly := "wordcounter.jar"

