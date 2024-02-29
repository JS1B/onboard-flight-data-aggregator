Testing Documentation
=====================

Introduction
------------

This document outlines the testing procedures for our web application project, focusing on ensuring the reliability and functionality of our microservices and scripts, including install.sh and run.sh. Our testing strategy encompasses various test types to maintain high-quality standards in our software development lifecycle.
Testing Frameworks and Tools

- Go Testing Framework: Utilized for unit and integration testing of Go-based microservices.
- Docker: For containerized testing environments, ensuring consistency across platforms.
- ShellCheck: A tool for linting and static analysis of shell scripts.

Types of Tests
--------------

### Unit Tests

- Purpose: Test individual components for correctness.
- Execution: Run within the Go testing framework.

### Integration Tests

- Purpose: Ensure that different modules of the application work together as expected.
- Execution: Conducted after unit tests, using Docker containers to simulate production-like environments.

### Script Tests

- Purpose: Validate the functionality of install.sh and run.sh.
- Execution: Tested in different Docker containers to ensure compatibility across environments like Ubuntu and Arch.

Test Environment Setup
----------------------

- Docker Installation: Ensure Docker is installed on your system.
- Building Containers:
  - Ubuntu: `docker build -t ubuntu-test -f Dockerfile.ubuntu .`

Running Tests
-------------

### Script Tests

- Ubuntu Environment:
  - Run docker run -v $(pwd):/app ubuntu-test /app/install.sh
  - Followed by docker run -v $(pwd):/app ubuntu-test /app/run.sh
- Arch Environment:
  - Run docker run -v $(pwd):/app arch-test /app/install.sh
  - Followed by docker run -v $(pwd):/app arch-test /app/run.sh
- More OS distors can be added if necessary

Writing Tests
-------------

- Unit Tests: For backend, follow the Go testing conventions. Include a variety of input scenarios.
- Script Tests: Test for both successful execution and handling of failure cases.

Continuous Integration (CI)
---------------------------

TODO
<!-- - CI pipelines are configured to run tests on every commit.
- Integration with Docker to test scripts in different environments. -->

Code Coverage
-------------

- Aim for at least 80% code coverage.
- Coverage reports are generated after each test run.

Troubleshooting Common Issues
-----------------------------

- Ensure Docker containers are built correctly.
- Check for environment-specific issues in the scripts.

Contributing Tests
------------------

- Contributions should include tests for new features or bug fixes.
- Follow the established patterns and update this document if necessary.

Contact for Queries
-------------------

- For any testing-related queries, please contact [p.krzeminski7@gmail.com].
