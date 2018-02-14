import java.math.BigDecimal;

import java.util.Date;
import java.util.List;
import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

public class Hibernate {
	private static EntityManagerFactory entityManagerFactory;

	public static void main(String[] args) {
		System.out.println("============ Hibernate ============");
		// The name matches the name we gave the persistence-unit in persistence.xml!
		System.out.println("--- Creating an Entity Manager...");
		entityManagerFactory = Persistence.createEntityManagerFactory("hibernate.learn");
		EntityManager entityManager = entityManagerFactory.createEntityManager();

		System.out.println("--- Persisting some events...");
		entityManager.getTransaction().begin();
		entityManager.persist(new Event("Our very first event!", new Date()));
		entityManager.persist(new Event("A follow up event", new Date()));
		entityManager.getTransaction().commit();
		entityManager.close();

		System.out.println("--- Persisting some orchestrations...");
		entityManager = entityManagerFactory.createEntityManager();
		entityManager.getTransaction().begin();
		entityManager.persist(new Orchestration(new BigDecimal(1), "Parameters (which should be JSON...)"));
		entityManager.persist(new Orchestration(new BigDecimal(2), "More parameters (which should also be JSON...)"));
		entityManager.persist(new Orchestration(new BigDecimal(3), "Even more parameters (which, again, should also be JSON...)"));
		entityManager.getTransaction().commit();
		entityManager.close();

		System.out.println("--- Retrieving the events...");
		entityManager = entityManagerFactory.createEntityManager();
		entityManager.getTransaction().begin();
		List<Event> events = entityManager.createQuery( "from Event", Event.class ).getResultList();
		for (Event event : events) {
			System.out.println( "--- Event (" + event.getDate() + ") : " + event.getTitle() );
		}
		entityManager.getTransaction().commit();
		entityManager.close();

		System.out.println("--- Retrieving the orchestrations...");
		entityManager = entityManagerFactory.createEntityManager();
		entityManager.getTransaction().begin();
		List<Orchestration> orchestrations = entityManager.createQuery( "from Orchestration", Orchestration.class ).getResultList();
		for (Orchestration orchestration : orchestrations) {
			System.out.println( "--- Orchestration (" + orchestration.getId() + ") : " + orchestration.getParameters() );
		}
		entityManager.getTransaction().commit();
		entityManager.close();

		System.out.println("--- Closing the Entity Manager...");
		entityManagerFactory.close();
		System.out.println("--- Done.");
	}
}



