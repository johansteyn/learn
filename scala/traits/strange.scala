trait Logger {
	def log(message: String)
}


trait ConsoleLogger extends Logger {
	def log(message: String) { 
		println(message)
	}
}

trait TimestampLogger extends ConsoleLogger {
	override def log(message: String) {
		super.log(s"${java.time.Instant.now()} $message")
	}
}

trait FileLogger extends Logger {
	import java.io.{FileWriter, PrintWriter}
	val pw = new PrintWriter(new FileWriter("strange.log", true))

	def log(message: String) {
		pw.println(message);
		pw.flush()
	}
}

trait FileLogger2 extends ConsoleLogger {
	import java.io.{FileWriter, PrintWriter}
	val pw = new PrintWriter(new FileWriter("strange.log", true))

	override def log(message: String) {
		pw.println(message);
		pw.flush()
	}
}

import java.io.File
val file = new File("strange.log")
if (file.exists()) {
	file.delete()
}

object logger1 extends ConsoleLogger
logger1.log("logger1: This message should appear on the console.")

object logger2 extends ConsoleLogger with TimestampLogger
logger2.log("logger2: This message should also appear on the console, with a timestamp.")

object logger3 extends FileLogger
logger3.log("logger3: This message should appear in the file.")

object logger4 extends FileLogger2
logger4.log("logger4: This message should also appear in the file - even though the trait extends ConsoleLogger.")

object logger5 extends FileLogger with TimestampLogger
logger5.log("logger5: This message should also appear in the file, with a timestamp, but it doesn't...")
//              Logger
//                |   \
//                |    \
//                |     ConsoleLogger
//                |          |
// logger5---FileLogger---TimestampLogger

object logger6 extends FileLogger2 with TimestampLogger
logger6.log("logger6: Yet this message appears in the file, even though the trait extends ConsoleLogger!")
//                    Logger
//                      | 
//                 ConsoleLogger
//                    /   \
// logger6---FileLogger2---TimestampLogger

// -------------------------------------------------------------------------------------------------------
// Some extra experiments...

object logger7 extends FileLogger with ConsoleLogger {
//	override def log(message: String) {
//		pw.println(message);
//		pw.flush()
//	}
	override def log = super[FileLogger].log
}
logger7.log("logger7: This message appears in the file, but only after overriding the 'log' method.")

//object logger8 extends ConsoleLogger with FileLogger
//logger8.log("logger8: ???")


