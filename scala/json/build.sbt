name := "AOBMessage"
version := "1.0"
scalaVersion := "2.10.4"

libraryDependencies ++= Seq(
  "com.typesafe.play" %% "play-json" % "2.6.10",
  "com.lihaoyi" %% "fastparse" % "1.0.0",
  // We need the Macro paradise plugin as per:
  // https://github.com/lihaoyi/fastparse/issues/72
  compilerPlugin("org.scalamacros" % "paradise" % "2.0.1" cross CrossVersion.full)
)

