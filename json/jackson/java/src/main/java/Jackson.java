import java.io.IOException;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

public class Jackson {
	public static void main(String[] args) throws IOException {
		ObjectMapper mapper = new ObjectMapper();

		String carString = "{ \"brand\" : \"Mercedes\", \"doors\" : 5 }";
		// NOTE: Using a string for the doors works the same...
		//String carString = "{ \"brand\" : \"Mercedes\", \"doors\" : \"5\" }";
		Car car = mapper.readValue(carString, Car.class);
		System.out.println("" + car);

		String modernString = "{\"status\":200, \"data\": {\"decodingTime\":10, \"translation\":\"hello\", \"sourceWordcount\":1, \"targetWordcount\":1}}";
		ModernMT modern = mapper.readValue(modernString, ModernMT.class);
		System.out.println("" + modern);
	}
}

class Car {
	private String brand;
	private int doors;

	public String getBrand() { 
		return brand;
	}
	
	public void setBrand(String brand) {
		this.brand = brand;
	}

	public int getDoors() {
		return doors;
	}
	
	public void setDoors (int doors) {
		this.doors = doors;
	}

	public String toString() {
		StringBuilder sb = new StringBuilder();
		return sb.append("Car: brand=").append(brand).append(", doors=").append(doors).toString();	
	}
}

class ModernMT {
	private String status;
	private Data data;

	public String getStatus() { 
		return status;
	}
	
	public void setStatus(String status) {
		this.status = status;
	}

	public Data getData() {
		return data;
	}
	
	public void setData (Data data) {
		this.data = data;
	}

	public String toString() {
		StringBuilder sb = new StringBuilder();
		return sb.append("ModernMT: status=").append(status).append(", data=").append(data).toString();
	}
}

class Data {
	private int decodingTime;
	private String translation;
	private int sourceWordcount;
	private int targetWordcount;

	public int getDecodingTime() {
		return decodingTime;
	}
	
	public void setDecodingTime (int decodingTime) {
		this.decodingTime = decodingTime;
	}

	public String getTranslation() { 
		return translation;
	}
	
	public void setTranslation(String translation) {
		this.translation = translation;
	}

	public int getSourceWordcount() {
		return sourceWordcount;
	}
	
	public void setSourceWordcount (int sourceWordcount) {
		this.sourceWordcount = sourceWordcount;
	}

	public int getTargetWordcount() {
		return targetWordcount;
	}
	
	public void setTargetWordcount (int targetWordcount) {
		this.targetWordcount = targetWordcount;
	}

	public String toString() {
		StringBuilder sb = new StringBuilder();
		return sb.append("Data: decodingTime=").append(decodingTime).append(", translation=").append(translation).append(", sourceWordcount=").append(sourceWordcount).append(", targetWordcount=").append(targetWordcount).toString();
	}
}
