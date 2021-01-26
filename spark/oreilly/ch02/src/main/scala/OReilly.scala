import org.apache.spark.sql.SparkSession

object OReilly {
  def main(args: Array[String]): Unit = {
    val spark = SparkSession.builder().appName("OReilly").master("local[1]").getOrCreate()

    val myRange = spark.range(1000).toDF("number")
    val divisBy2 = myRange.where("number % 2 = 0")
    println(s"divisBy2=" + divisBy2)
    val count = divisBy2.count()
    println("count=" + count)

    val flightData2015 = spark
      .read
      .option("inferSchema", "true")
      .option("header", "true")
//      .csv("2015-summary.csv")
      .csv(args(0))
    val firstTen = flightData2015.take(10)
    firstTen.foreach(row => println(row))
//    val sorted = flightData2015.sort("count")
    val sorted = flightData2015.sort("count", "ORIGIN_COUNTRY_NAME")
    sorted.explain()
    val firstTenSorted = sorted.take(10)
    firstTenSorted.foreach(row => println(row))

    flightData2015.createOrReplaceTempView("flight_data_2015")
    val sqlWay = spark.sql("""
      SELECT DEST_COUNTRY_NAME, count(1)
      FROM flight_data_2015
      GROUP BY DEST_COUNTRY_NAME
      """)
    sqlWay.explain

    val dfWay = flightData2015
      .groupBy("DEST_COUNTRY_NAME") // NOTE: Book has a mistake on page 25...
      .count()
    dfWay.explain

    val maxSQL = spark.sql("""
      SELECT DEST_COUNTRY_NAME, sum(count) as destination_total
      FROM flight_data_2015
      GROUP BY DEST_COUNTRY_NAME
      ORDER BY sum(count) DESC
      LIMIT 5
      """)
    maxSQL.show()

    import org.apache.spark.sql.functions.desc
      flightData2015
        .groupBy("DEST_COUNTRY_NAME")
        .sum("count")
        .withColumnRenamed("sum(count)", "destination_total")
        .sort(desc("destination_total"))
        .limit(5)
        .show()
  }
}

