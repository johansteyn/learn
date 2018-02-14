import java.math.BigDecimal;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

@Entity
@Table( name = "ORCHESTRATIONS" )
public class Orchestration {
	private BigDecimal id; // Using BigDecimal instead of "long" results in NUMBER(19,2) instead of NUMBER(19,0)
	private String parameters;

	public Orchestration() {
		// Used by Hibernate, otherwise we get exception:
		//   javax.persistence.PersistenceException: org.hibernate.InstantiationException: 
		//     No default constructor for entity:  : Orchestration
	}

	public Orchestration(BigDecimal id, String parameters) {
		// Used by application
		this.id = id;
		this.parameters = parameters;
	}

	@Id
	public BigDecimal getId() {
		return id;
	}

	private void setId(BigDecimal id) {
		this.id = id;
	}

	public String getParameters() {
		return parameters;
	}

	public void setParameters(String parameters) {
		this.parameters = parameters;
	}
}

