# gcp

These are experiments exploring the Google Cloud Java Client API in both Java and Scala, starting with Google Storage:

  https://cloud.google.com/storage/docs/reference/libraries

## Build and Run Instructions

Both the Java and Scala builds result in "fat" JARs that include all dependent JARs.

### Java

```$ cd java```

```$ mvn clean package```

```$ ./run.sh <bucket>```
  
### Scala

```$ cd scala```

```$ sbt clean assembly```

```$ ./run.sh <bucket>```

NOTE: Running these apps requires GCP credentials, which are not included here as they do not belong in Git.

