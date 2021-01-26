import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import com.google.api.gax.paging.Page;

import com.google.cloud.WriteChannel;
import com.google.cloud.storage.Blob;
import com.google.cloud.storage.BlobId;
import com.google.cloud.storage.BlobInfo;
import com.google.cloud.storage.Bucket;
import com.google.cloud.storage.BucketInfo;
import com.google.cloud.storage.CopyWriter;
import com.google.cloud.storage.Storage;
import com.google.cloud.storage.StorageOptions;

public class GoogleStorage {
  private static Logger logger = Logger.getLogger(GoogleStorage.class.getName());
  private static Storage storage;
  private static String bucketName;

  private static void log(String message) {
    logger.log(Level.INFO, message);
  }

  private static void listBuckets()   { 
    Page<Bucket> buckets = storage.list();
    System.out.println("Buckets:");
    for (Bucket b : buckets.iterateAll()) {
      System.out.println("  " + b.getName());
    }
    System.out.println();
  }

  private static void createBucket() { 
    log("Creating new bucket: " + bucketName);
    Bucket bucket = storage.create(BucketInfo.of(bucketName));
    System.out.println("Created bucket with metadata:");
    System.out.println("                   Name: " + bucket.getName());
    System.out.println("  DefaultEventBasedHold: " + bucket.getDefaultEventBasedHold());
    System.out.println("      DefaultKmsKeyName: " + bucket.getDefaultKmsKeyName());
    System.out.println("            GeneratedId: " + bucket.getGeneratedId());
    System.out.println("              IndexPage: " + bucket.getIndexPage());
    System.out.println("               Location: " + bucket.getLocation());
    System.out.println("           LocationType: " + bucket.getLocationType());
    System.out.println("         Metageneration: " + bucket.getMetageneration());
    System.out.println("           NotFoundPage: " + bucket.getNotFoundPage());
    System.out.println(" RetentionEffectiveTime: " + bucket.getRetentionEffectiveTime());
    System.out.println("        RetentionPeriod: " + bucket.getRetentionPeriod());
    System.out.println("RetentionPolicyIsLocked: " + bucket.retentionPolicyIsLocked());
    System.out.println("          RequesterPays: " + bucket.requesterPays());
    System.out.println("               SelfLink: " + bucket.getSelfLink());
    System.out.println("           StorageClass: " + bucket.getStorageClass().name());
    System.out.println("            TimeCreated: " + bucket.getCreateTime());
    System.out.println("      VersioningEnabled: " + bucket.versioningEnabled());
    if (bucket.getLabels() != null) {
      System.out.println("                 Labels:");
      for (Map.Entry<String, String> label : bucket.getLabels().entrySet()) {
        System.out.println("  " + label.getKey() + "=" + label.getValue());
      }
    }
    if (bucket.getLifecycleRules() != null) {
      System.out.println("     Lifecycle Rules:");
      for (BucketInfo.LifecycleRule rule : bucket.getLifecycleRules()) {
        System.out.println("  " + rule);
      }
    }
  }

  private static void upload1KB(int n, String objectName) throws IOException { 
    String filePath = "1KB";
    log("Uploading contents of file '" + filePath + "' to bucket '" + bucketName + "' " + n + " times as object '" + objectName + "'...");
    byte[] bytes = Files.readAllBytes(Paths.get(filePath));
    BlobId blobId = BlobId.of(bucketName, objectName);
    BlobInfo blobInfo = BlobInfo.newBuilder(blobId).build();
    try (WriteChannel writer = storage.writer(blobInfo)) {
      for (int i = 0; i < n; i++) {
        writer.write(ByteBuffer.wrap(bytes, 0, bytes.length));
      }
    }
    log("Done");
  }

  private static void listObjects() {
    Bucket bucket = storage.get(bucketName);
    Page<Blob> blobs = bucket.list();
    System.out.println("Objects:");
    for (Blob blob : blobs.iterateAll()) {
      System.out.println("  " + blob.getName());
    }
    System.out.println();
  }

  private static void renameObject(String from, String to) {
    Blob fromBlob = storage.get(bucketName, from);
    CopyWriter copyWriter = fromBlob.copyTo(bucketName, to);
    Blob toBlob = copyWriter.getResult();
    fromBlob.delete();
    System.out.println("Renamed object '" + fromBlob.getName() + "' to '" + toBlob.getName() + "'");
  }

  private static void downloadObject(String objectName) {
    log("Downloading object '" + objectName + "'...");
    Blob blob = storage.get(bucketName, objectName);
    blob.downloadTo(Path.of(objectName));
    log("Done.");
  }

  public static void main(String[] args) throws IOException {
    storage = StorageOptions.getDefaultInstance().getService();
    bucketName = args[0];
    listBuckets();
    createBucket();
    upload1KB(1, "KiloByte");
    upload1KB(1024, "MegoByte");
    upload1KB(1024 * 1024, "GigaByte");
    listObjects();
    renameObject("MegoByte", "MegaByte");
    downloadObject("MegaByte");
  }
}

