name := "spark-linkedin"

version := "1.0"

scalaVersion := "2.11.12"

libraryDependencies += "org.apache.spark" %% "spark-sql" % "2.2.0"

//assemblyMergeStrategy in assembly := {
// case PathList("META-INF", xs @ _*) => MergeStrategy.discard
// case x => MergeStrategy.first
//}



