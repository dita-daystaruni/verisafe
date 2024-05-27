# VeriSafe

VeriSafe is an authentication service tailored for academia backend services. It offers robust security measures and streamlined authentication processes to ensure safe and efficient access to academic resources.

## Key Features

- **Multi-Factor Authentication (MFA)**: Adds an extra layer of security by requiring users to authenticate via multiple methods.
- **Single Sign-On (SSO)**: Allows users to access various systems with one set of credentials, improving user experience.
- **Role-Based Access Control (RBAC)**: Manages user permissions based on their roles, ensuring access to only relevant resources.
- **Audit Logging**: Records login attempts, successful authentications, and system accesses for security monitoring and incident analysis.
- **Easy Integration**: Seamlessly integrates with existing academic platforms and services to enhance authentication capabilities.

## Quick Start Guide

### Requirements

- Go installed.
- Basic knowledge of Go programming.
- Familiarity with Docker.

### Setup Instructions

1. Clone the repository:

```bash
git clone https://github.com/dita-daystaruni/verisafe.git && cd verisafe
```


2. Build the Docker image:

```bash
docker build -t verisafe.
```


3. Run VeriSafe using Docker:
```bash
docker run -p 8080:8080 verisafe
```


Access VeriSafe at `http://localhost:8080`.

## Further Information

For detailed guides and documentation, visit our official documentation site at `https://docs.verisafe.academia`.

## Contribution Guidelines

Contributions are welcome Please review our contributing guidelines for more details.

## License

VeriSafe operates under the MIT License. Check the LICENSE file for full details.

