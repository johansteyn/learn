import java.nio.ByteBuffer
import java.nio.file.{Files, Path, Paths}
import java.util.logging.{Level, Logger}
import scala.collection.JavaConverters._
import com.google.api.gax.paging.Page
import com.google.cloud.WriteChannel
import com.google.cloud.storage.{Blob, BlobId, BlobInfo}
import com.google.cloud.storage.{Bucket, BucketInfo, CopyWriter, Storage, StorageOptions}

object GoogleStorage extends App {
  val logger: Logger = Logger.getLogger("GoogleStorage")
  val storage: Storage = StorageOptions.getDefaultInstance().getService()
  val bucketName: String = args(0)
  listBuckets()
  createBucket()
  upload1KB(1, "KiloByte")
  upload1KB(1024, "MegoByte")
  upload1KB(1024 * 1024, "GigaByte")
  listObjects()
  renameObject("MegoByte", "MegaByte")
  downloadObject("MegaByte")
  
  def log(message: String) = logger.log(Level.INFO, message)

  def listBuckets() = {
    println("Buckets:")
    val buckets = storage.list().iterateAll().asScala
    for (bucket <- buckets) println("  " + bucket.getName)
    println()
  }

  def createBucket() = {
    log("Creating new bucket: " + bucketName)
    val bucket: Bucket = storage.create(BucketInfo.of(bucketName))
    println("Created bucket with metadata:")
    println("                   Name: " + bucket.getName())
    println("  DefaultEventBasedHold: " + bucket.getDefaultEventBasedHold())
    println("      DefaultKmsKeyName: " + bucket.getDefaultKmsKeyName())
    println("            GeneratedId: " + bucket.getGeneratedId())
    println("              IndexPage: " + bucket.getIndexPage())
    println("               Location: " + bucket.getLocation())
    println("           LocationType: " + bucket.getLocationType())
    println("         Metageneration: " + bucket.getMetageneration())
    println("           NotFoundPage: " + bucket.getNotFoundPage())
    println(" RetentionEffectiveTime: " + bucket.getRetentionEffectiveTime())
    println("        RetentionPeriod: " + bucket.getRetentionPeriod())
    println("RetentionPolicyIsLocked: " + bucket.retentionPolicyIsLocked())
    println("          RequesterPays: " + bucket.requesterPays())
    println("               SelfLink: " + bucket.getSelfLink())
    println("           StorageClass: " + bucket.getStorageClass().name())
    println("            TimeCreated: " + bucket.getCreateTime())
    println("      VersioningEnabled: " + bucket.versioningEnabled())
    if (bucket.getLabels() != null) {
      println("                 Labels:")
      val labels = bucket.getLabels().asScala
      for ((k,v) <- labels) println("  " + k + "=" + v)
    }
    if (bucket.getLifecycleRules() != null) {
      println("     Lifecycle Rules:")
      val rules = bucket.getLifecycleRules().asScala
      for (rule <- rules) println("  " + rule)
    }
  }

  def upload1KB(n: Int, objectName: String) = { 
    val filePath = "1KB"
    log("Uploading contents of file '" + filePath + "' to bucket '" + bucketName + "' " + n + " times as object '" + objectName + "'...")
    val bytes: Array[Byte] = Files.readAllBytes(Paths.get(filePath))
    val blobId: BlobId = BlobId.of(bucketName, objectName)
    val blobInfo: BlobInfo = BlobInfo.newBuilder(blobId).build()
    val writer: WriteChannel = storage.writer(blobInfo)
    for (i <- 0 until n) writer.write(ByteBuffer.wrap(bytes, 0, bytes.length))
    writer.close
    log("Done")
  }

  def listObjects() = {
    val bucket: Bucket = storage.get(bucketName)
    val blobs = bucket.list().iterateAll().asScala
    println("Objects:")
    for (blob <- blobs) println("  " + blob.getName)
    println()
  }

  def renameObject(from: String, to: String) = {
    val fromBlob: Blob = storage.get(bucketName, from)
    val copyWriter: CopyWriter = fromBlob.copyTo(bucketName, to)
    val toBlob: Blob = copyWriter.getResult()
    fromBlob.delete()
    println("Renamed object '" + fromBlob.getName() + "' to '" + toBlob.getName() + "'")
  }

  def downloadObject(objectName: String) = {
    log("Downloading object '" + objectName + "'...")
    val blob: Blob = storage.get(bucketName, objectName)
    blob.downloadTo(Path.of(objectName))
    log("Done.")
  }
}

