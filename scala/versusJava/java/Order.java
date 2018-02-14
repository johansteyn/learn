import java.util.List;

public class Order {
	private int id;
	private List<Product> products;
	
	public Order(int id, List<Product> products) {
		this.id = id;
		this.products = products;
	}
	
	public int getId() {
		return id;
	}
	
	public void setId(int id) {
		if (id <= 0) {
			throw new IllegalArgumentException("Order ID must be larger than zero!");
		}
		this.id = id;
	}
	
	public List<Product> getProducts() {
		return products;
	}
	
	public void setProducts(List<Product> products) {
		this.products = products;
	}
}

