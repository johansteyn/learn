import org.apache.spark.SparkConf
import org.apache.spark.SparkContext
import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.SQLContext

object LinkedIn {
  def main(args: Array[String]): Unit = {
    val conf = new SparkConf().setAppName("LinkedIn").setMaster("local[*]")
    val sc = new SparkContext(conf)

    val lines = sc.parallelize(List("This is line 1", "This is line 2", "This is third line"))
    lines.foreach(println _)
    val filteredLines = lines.filter(line => line.contains("2"))
    filteredLines.foreach(println _)

    val sqlContext = new SQLContext(sc)
    val df: DataFrame = sqlContext.read.json("employee.json")
    df.show()
    df.printSchema()
    df.select("name").show()
    df.filter(df("age") > 40).show()
  }
}

