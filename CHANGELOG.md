# Table of Contents

<!---
Adding a new version? You'll need three changes:
* Add the ToC link, like "[v1.2.3](#v123)".
* Add the section header, like "## [v1.2.3]".
* Add the diff link, like "[v2.7.0]: https://github.com/kong/kubernetes-ingress-controller/compare/v1.2.2...v1.2.3".
  This is all the way at the bottom. It's the thing we always forget.
--->
- [1.0.0 and prior](#100-and-prior)

## Unreleased

> Release date: TBA

## [v1.0.0]

> Release date: 2024-11-26

This is an initial stable release of the Kong Kubernetes configuration module.
It includes CRDs for the Kong Ingress Controller and the Kong Gateway Operator
in the `ingress-controller`, `ingress-controller-incubator` and `gateway-operator`
channels.

Go bindings for the CRDs are available in the [`api/`][api] directory and the
[`pkg/clientset/`][clientset] directory contains the clientset for interacting
with the CRDs.

[v1.0.0]: https://github.com/kong/kubernetes-configuration/compare/{}...v1.0.0
