name := "Jackson"

version := "1.0"

scalaVersion := "2.12.3"

resolvers += "Typesafe Repository" at "http://repo.typesafe.com/typesafe/releases/"

libraryDependencies += "com.fasterxml.jackson.core" % "jackson-core" % "2.9.4"
libraryDependencies += "com.fasterxml.jackson.core" % "jackson-annotations" % "2.9.4"
libraryDependencies += "com.fasterxml.jackson.core" % "jackson-databind" % "2.9.4"
libraryDependencies += "com.fasterxml.jackson.module" %% "jackson-module-scala" % "2.9.4"

assemblyJarName in assembly := "johansteyn-jackson-scala.jar"

