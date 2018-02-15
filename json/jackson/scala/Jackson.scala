import java.io.IOException

import scala.beans._
import scala.reflect._

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.scala.DefaultScalaModule
import com.fasterxml.jackson.module.scala.experimental.ScalaObjectMapper

object Main extends App {
	val mapper = new ObjectMapper()

	var carString = """{ "brand" : "Mercedes", "doors" : 5 }"""
	val car = mapper.readValue(carString, classOf[Car])
	println("" + car)

	val modernString = """{"status":200, "data": {"decodingTime":10, "translation":"hello", "sourceWordcount":1, "targetWordcount":1}}"""
	val modern = mapper.readValue(modernString, classOf[ModernMT])
	println("" + modern)

	// A simpler approach, that uses plain maps instead of model objects...
	// https://coderwall.com/p/o--apg/easy-json-un-marshalling-in-scala-with-jackson
	val scalaMapper = new ObjectMapper() with ScalaObjectMapper
	scalaMapper.registerModule(DefaultScalaModule)
	val carMap = scalaMapper.readValue[Map[String, Object]](carString)
	println(carMap)
	val modernMap = scalaMapper.readValue[Map[String, Object]](modernString)
	println(modernMap)
}

class Car {
	@BeanProperty var brand: String = null
	@BeanProperty var doors: Int = 0

	override def toString(): String = {
		val sb = new StringBuilder()
		return sb.append("Car: brand=").append(brand).append(", doors=").append(doors).toString()
	}
}

class ModernMT {
	@BeanProperty var status: String = null
	@BeanProperty var data: Data = null

	override def toString(): String = {
		val sb = new StringBuilder()
		return sb.append("ModernMT: status=").append(status).append(", data=").append(data).toString()
	}
}

class Data {
	@BeanProperty var decodingTime: Int = 0
	@BeanProperty var translation: String = null
	@BeanProperty var sourceWordcount: Int = 0
	@BeanProperty var targetWordcount: Int = 0

	override def toString(): String = {
		val sb = new StringBuilder()
		return sb.append("Data: decodingTime=").append(decodingTime).append(", translation=").append(translation).append(", sourceWordcount=").append(sourceWordcount).append(", targetWordcount=").append(targetWordcount).toString()
	}
}

