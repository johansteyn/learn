import java.util.List;

public class User {
	private String name;
	private List<Order> orders;
	
	public User(String name, List<Order> orders) {
		this.name = name;
		this.orders = orders;
	}
	
	public String getName() {
		return name;
	}
	
	public void setName(String name) {
		this.name = name;
	}
	
	public List<Order> getOrders() {
		return orders;
	}
	
	public void setOrders(List<Order> orders) {
		this.orders = orders;
	}

	public static void main(String[] args) {
		User user = new User("Johan", null);
		System.out.println("Created user: " + user.getName());
		user.setName("Steyn");
		System.out.println("Changed user: " + user.getName());
		Order order = new Order(123, null);
		System.out.println("Created order #" + order.getId());
		//order.setId(-1);
		Product product = new Product(234, "Furniture");
		System.out.println("Created product #" + product.getId() + " (" + product.getCategory() + ")");
	}
}

