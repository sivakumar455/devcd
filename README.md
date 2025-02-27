# devCD

devCD is a comprehensive CI/CD tool designed to streamline the deployment and testing of microservices. It supports deploying applications using Helm charts or Docker Compose files, catering to both Kubernetes and containerized environments. 
With its simple and intuitive command-line interface, devCD allows developers to efficiently manage the deployment, scaling, and maintenance of their microservices and associated dependencies. By reducing the complexity of setup and orchestration, devCD enables teams to focus on development and innovation while ensuring consistency and reliability in their CI/CD workflows.

## Key Components

devCD is structured into the following sections:

1. **devcd**
   - The core module containing the primary logic for installing and deleting Helm charts.

2. **devcd-ext**
   - A customizable extension module where users can:
     - Add custom commands.
     - Include pre-setup configurations needed for the development CI/CD environment.

3. **Deployment Modes**
   devCD supports two deployment modes:

   - **Helm Charts Mode (devcd-helm)**
     - Contains all Helm charts required for deploying both backing services and microservices.
     - Ideal for deploying applications in a Kubernetes environment.

   - **Docker Compose Mode (devcd-compose)**
     - Contains Docker Compose files for deploying applications as containers.
     - Suitable for local testing or environments using Docker containers.

4. **devcd-runtime**
   - A runtime environment to:
     - Add microservice JARs (to be mounted in containers).
     - Store logs or any runtime files required during the CI/CD process.

## Installation Steps

### Pre-Setup
- Ensure your system meets the following requirements:
  - **Docker/Podman**: CPU: 4 cores, RAM: 8/12 GB
  - **Docker Desktop**: Kubernetes enabled
  - **Go**: 1.22

### Step 1: Build Executable

```bash
# Build the devCD executable
$ make build

# Run the following command for help
$ ./devc run help

```

### Step 2: Run devCD
```bash
# Start/stop all backing services and application microservices at once
$ ./devc run cd [start/stop]

# Start/stop application microservices only
$ ./devc run ms [start/stop]

# Start/stop application backing services only
$ ./devc run bs [start/stop]

```

### Step 3: Adding a New Microservice Helm Chart
```bash
# Navigate to the devcd-helm/ms-helm directory for ms:
$ cd devcd-helm/ms-helm

#Create a new folder for your microservice Helm chart:
#Add your Helm chart files (Chart.yaml, values.yaml, templates, etc.) inside the newly created folder.

$ mkdir <microservice-name-service>/mschart

# Add chart path in helm_ms_config.yaml file under helmMs section to include in cd deployment
$ vi helm_ms_config.yaml

# Start application microservice
$ ./devc run testms [start/stop]

# Refer existing charts for reference in the folder devcd-helm
```

## Appendix : 

- Kubernetes DashBoard Login
```bash
# Start the Kubernetes proxy
kubectl proxy

# Access the Kubernetes Dashboard UI
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy

# Retrieve the login token
kubectl get secret admin-user -n kubernetes-dashboard -o jsonpath={".data.token"} | base64 -d
```
- Setting insecure registries
 ```bash
# Add the following insecure registry in Docker settings
"insecure-registries": [
    "reg.company.com:1234"
]

# Perform Docker login

docker login reg.company.com:1234
# Username: user
# Password: <nexus password>
```

```bash
# Micro services example github repo
https://github.com/sivakumar455/e-market


```

