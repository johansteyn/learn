name := "gson-scala"

version := "1.0"

libraryDependencies += "com.google.code.gson" % "gson" % "2.8.5"
libraryDependencies += "com.google.guava" % "guava" % "28.2-jre"

assemblyMergeStrategy in assembly := {
 case PathList("META-INF", xs @ _*) => MergeStrategy.discard
 case x => MergeStrategy.first
}


