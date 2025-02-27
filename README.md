## Description

_"As a program evolves and acquires more features, it becomes complicated, with subtle dependencies between components. Over time, complexity accumulates, and it becomes harder and harder for programmers to keep all the relevant factors in their minds as they modify the system. This slows down development and leads to bugs, which slow development even more and add to its cost. Complexity increases inevitably over the life of any program. The larger the program, and the more people that work on it, the more difficult it is to manage complexity."_ - John Ousterhout  

The service starter kit is a starting point for building production grade scalable web service applications that leverage the power of a Domain Driven, Data Oriented Architecture that can run in Kubernetes. The goal of this project is to provide a proven starting point that reduces the repetitive tasks required for a new project to be launched into production. It uses minimal dependencies, implements idiomatic code and follows Go best practices. Collectively, the project lays out everything logically to minimize guess work and enable engineers to quickly maintain a mental model for the project.

The class behind this starter kit teaches how to build production-level software in Go leveraging the power of a Domain Driven, Data Oriented Architecture that can run in Kubernetes. From the beginning, you will pair program with the instructor walking through the design philosophies and guidelines for building software in Go. With each new feature that is added to the project, you will learn how to deploy to and manage the Kubernetes environment used to run the project. The core of this class is to teach you and your team how to handle and reduce the spread of complexity in the systems you are building.

Learn more about the project:

[Wiki](https://github.com/ardanlabs/service/wiki) | [Course Outline](https://github.com/ardanlabs/service/wiki/course-outline) | [Class Schedule](https://www.ardanlabs.com/events/)

## Licensing

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## Index

* [Installation](https://github.com/ardanlabs/service?tab=readme-ov-file#installation)
* [Create Your Own Version](https://github.com/ardanlabs/service?tab=readme-ov-file#create-your-own-version)
* [Running The Project](https://github.com/ardanlabs/service?tab=readme-ov-file#running-the-project)

## Installation

To clone the project, create a folder and use the git clone command. Then please read the [makefile](makefile) file to learn how to install all the tooling and docker images.

```
$ cd $HOME
$ mkdir code
$ cd code
$ git clone https://github.com/ardanlabs/service or git@github.com:ardanlabs/service.git
$ cd service
```

## Create Your Own Version

If you want to create a version of the project for your own use, use the new gonew command.

```
$ go install golang.org/x/tools/cmd/gonew@latest

$ cd $HOME
$ mkdir code
$ cd code
$ gonew github.com/ardanlabs/service github.com/mydomain/myproject
$ cd myproject
$ go mod vendor
```

Now you have a copy with your own module name. Now all you need to do is initialize the project for git.

## Running The Project

To run the project use the following commands.

```
# Install Tooling
$ make dev-gotooling
$ make dev-brew
$ make dev-docker

# Run Tests
$ make test

# Shutdown Tests
$ make test-down

# Run Project
$ make dev-up
$ make dev-update-apply
$ make token
$ export TOKEN=<COPY TOKEN>
$ make users

# Run Load
$ make load

# Run Tooling
$ make grafana
$ make statsviz

# Shut Project
$ make dev-down
```
