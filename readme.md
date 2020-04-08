# AWS Controller

| API Version | Website | Build | Coverage | Report | License | Release |
|:-----------:|:-------:|:-----:|:--------:|:------:|:-------:|:-------:|
| [![v1alpha1](https://img.shields.io/badge/apiversion-v1alpha1-red.svg)](https://github.com/awsctrl/manager/) | [![awsctrl.io](https://img.shields.io/badge/website-n/a-red.svg)](https://awsctrl.io) | [![TravisCI](https://travis-ci.org/awsctrl/manager.svg?branch=master)](https://travis-ci.org/awsctrl/manager) | [![CodeCov](https://codecov.io/gh/awsctrl/manager/branch/master/graph/badge.svg)](https://codecov.io/gh/awsctrl/manager) | [![GoReport Card](https://goreportcard.com/badge/github.com/awsctrl/manager)](https://goreportcard.com/report/github.com/awsctrl/manager) | [![License](https://img.shields.io/github/license/awsctrl/manager.svg)](https://github.com/awsctrl/manager/blob/master/LICENSE) | [![Release](https://img.shields.io/github/release/awsctrl/manager.svg)](https://github.com/awsctrl/manager/releases/latest) |

<!--ts-->
   * [AWS Controller](#aws-controller)
   * [Features](#features)
   * [Overview](#overview)
   * [Installation](#installation)
   * [Requirements](#requirements)

<!--te-->

AWS Controller is Kubernetes controller manager which contains controllers that manage AWS resources using [custom resource definitions (CRDs)](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/).

Features
========

-  Kubernetes Native; *using CRDs*
- 100% Open Source
- Infrastructure as Configuration

Overview
========

The AWS Controller is an open source Kubernetes native application that
runs in your Kubernetes cluster and listens to the Kubernetes API Server
for AWS related [Custom
Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/).
These custom resources allow you to create external Cloud Native
resources like [Amazon S3 Buckets](https://aws.amazon.com/s3/), [AWS ECR
Repositories](https://aws.amazon.com/ecr/), and many more.

By using the AWS Controller you enable a streamlined experience to
deploying your Cloud Native applications on AWS. No longer do you need
to manage complex deployment pipelines or manage separate deployment
tools for infrastructure as you do your application. Now you can deploy
everything using `kubectl apply -f`.

![](https://awsctrl.io/images/controller-flow.png)

Installation
============

With a Kubernetes cluster configured with [Amazon EKS Pod Identity
Webhook](https://github.com/aws/amazon-eks-pod-identity-webhook) and
your `KUBECONFIG` configured locally you can install the `awsctrl`
binary and then install the controller.

    go install go.awsctrl.io/manager/cmd/awsctrl
    awsctrl install manager | kubectl apply -f -
    awsctrl install config | kubectl apply -f -

For more details see [AWS Controller Website](https://awsctrl.io).

Requirements
============

-   Kubernetes &gt;1.15
