name := "etl-scala"

version := "1.0"

libraryDependencies += "com.datastax.cassandra" % "cassandra-driver-core" % "3.8.0"
libraryDependencies += "com.google.code.gson" % "gson" % "2.8.5"
libraryDependencies += "com.google.guava" % "guava" % "28.2-jre"

assemblyMergeStrategy in assembly := {
 case PathList("META-INF", xs @ _*) => MergeStrategy.discard
 case x => MergeStrategy.first
}


