<p align="center">
  <img src="https://static.wikia.nocookie.net/disney_mirrorverse/images/6/64/Hades.png/revision/latest?cb=20220708113837" width=300" />
</p>

# Hades

This project is a __Kubernetes Operator__ designed to automate the creation of jobs for a cronjob until a successful one is achieved.
In the event of a cronjob failure, this operator initiates the generation of new jobs at regular intervals until a successful job is obtained.
This operator is activated following the identification of a cronjob failure by Kubernetes.

## How it works?

This system includes a monitoring mechanism that oversees cronjobs within a specific namespace.
Upon detecting a failed cronjob, it records this occurrence in a MySQL database.
Subsequently, an agent routinely retrieves these entries from the database and generates Jobs until a successful job is achieved.
The timing interval and list of cronjobs can be configured within the operator's configmap.

## Why Hades?

__Hades__ is a prominent figure in Greek mythology, known as the ruler of the underworld and the god of the dead.
He is one of the three major Olympian gods, alongside his brothers Zeus and Poseidon. Hades is often depicted as stern and unyielding,
residing in his realm of the dead where souls go after death. He is typically portrayed as a somber figure, holding a scepter or key that symbolizes
his control over the underworld. Despite his ominous role, Hades is not considered inherently evil; rather, he maintains order and balance in the realm of the dead.

The operator functions akin to the mythical figure of __Hades__.
In a similar fashion, this operator is responsible for managing failed cronjobs, which can be likened to souls in the underworld.
Just as Hades maintains order and control over the deceased, this operator ensures that failed cronjobs are appropriately handled and managed within the system.
It embodies a sense of authority and regulation, akin to the role played by Hades in Greek mythology.

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/operator:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/operator:tag
```

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

UnDeploy the controller to the cluster:

```sh
make undeploy
```

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/)
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster

### Test It Out

1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Contribute

If you find a suitable issue, express your interest in working on it and proceed to fork the project's repository to your GitHub account.
From there, create a new branch to make your changes, ensuring they adhere to the project's coding style and guidelines.
Thoroughly test your modifications to confirm they work as intended before committing and pushing them to your forked repository.
Finally, submit a detailed pull request to the original project's repository, including an explanation of your changes and their benefits.
Engage in constructive discussions with project maintainers and other contributors, respecting their guidelines and feedback throughout the process.
