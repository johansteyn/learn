class User(var name: String, var orders: List[Order])

// To override a field setter:
// - Rename the constructor parameter (usually by prepending an underscore)
// - Provide a getter and a setter using the original name
class Order (private var _id: Int, var products: List[Product]) {
	def id = _id // Getter
	def id_=(value: Int) = {
		// Setter
		if (value <= 0) throw new IllegalArgumentException("Order ID must be larger than zero!")
		_id = value
	}
}

class Product (var id: Int, var category: String)

object Main extends App {
	val user = new User("Johan", null)
	println("Created user: " + user.name)
	user.name = "Steyn"
	println("Changed user: " + user.name)
	val order = new Order(123, null)
	println("Created order #" + order.id)
	//order.id = -1
	val product = new Product(234, "Furniture")
	println("Created product #" + product.id + " (" + product.category + ")")
}

