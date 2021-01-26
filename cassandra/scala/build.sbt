name := "cassandra-scala"

version := "1.0"

libraryDependencies += "com.datastax.cassandra" % "cassandra-driver-core" % "3.8.0"

assemblyMergeStrategy in assembly := {
 case PathList("META-INF", xs @ _*) => MergeStrategy.discard
 case x => MergeStrategy.first
}


