import java.io.IOException;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

public class Jackson {
	public static void main(String[] args) throws IOException {
		ObjectMapper objectMapper = new ObjectMapper();

		String string = "{ \"brand\" : \"Mercedes\", \"doors\" : 5 }";
		// NOTE: Using a string for the doors works the same...
		//String string = "{ \"brand\" : \"Mercedes\", \"doors\" : \"5\" }";
		Car car = objectMapper.readValue(string, Car.class);
		System.out.println("" + car);

		string = "{\"status\":200, \"data\": {\"decodingTime\":10, \"translation\":\"hello\", \"sourceWordcount\":1, \"targetWordcount\":1}}";
		ModernMT modern = objectMapper.readValue(string, ModernMT.class);
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
		sb.append("Car: brand=");
		sb.append(brand);
		sb.append(", doors=");
		sb.append(doors);
		return sb.toString();
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
		sb.append("ModernMT: status=");
		sb.append(status);
		sb.append(", data=");
		sb.append(data);
		return sb.toString();
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
		sb.append("Data: decodingTime=");
		sb.append(decodingTime);
		sb.append(", translation=");
		sb.append(translation);
		sb.append(", sourceWordcount=");
		sb.append(sourceWordcount);
		sb.append(", targetWordcount=");
		sb.append(targetWordcount);
		return sb.toString();
	}
}
