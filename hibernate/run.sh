#!/bin/bash

echo "Running..."

java -cp \
ojdbc7-12.1.0.2.0.jar:\
hibernate-release-5.2.11.Final/lib/required/antlr-2.7.7.jar:\
hibernate-release-5.2.11.Final/lib/required/classmate-1.3.0.jar:\
hibernate-release-5.2.11.Final/lib/required/dom4j-1.6.1.jar:\
hibernate-release-5.2.11.Final/lib/required/hibernate-commons-annotations-5.0.1.Final.jar:\
hibernate-release-5.2.11.Final/lib/required/hibernate-core-5.2.11.Final.jar:\
hibernate-release-5.2.11.Final/lib/required/hibernate-jpa-2.1-api-1.0.0.Final.jar:\
hibernate-release-5.2.11.Final/lib/required/jandex-2.0.3.Final.jar:\
hibernate-release-5.2.11.Final/lib/required/javassist-3.20.0-GA.jar:\
hibernate-release-5.2.11.Final/lib/required/jboss-logging-3.3.0.Final.jar:\
hibernate-release-5.2.11.Final/lib/required/jboss-transaction-api_1.2_spec-1.0.1.Final.jar:\
target/hibernate-learn.jar Hibernate

