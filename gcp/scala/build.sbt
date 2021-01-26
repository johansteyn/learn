name := "gcp-scala"

version := "1.0"

libraryDependencies += "com.google.cloud" % "google-cloud-storage" % "1.105.0"

assemblyMergeStrategy in assembly := {
 case PathList("META-INF", xs @ _*) => MergeStrategy.discard
 case x => MergeStrategy.first
}


