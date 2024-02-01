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

## Setup using Helm Charts

### docker image

### configs

## Contribute

If you find a suitable issue, express your interest in working on it and proceed to fork the project's repository to your GitHub account.
From there, create a new branch to make your changes, ensuring they adhere to the project's coding style and guidelines.
Thoroughly test your modifications to confirm they work as intended before committing and pushing them to your forked repository.
Finally, submit a detailed pull request to the original project's repository, including an explanation of your changes and their benefits.
Engage in constructive discussions with project maintainers and other contributors, respecting their guidelines and feedback throughout the process.
