# NOTE

**UPDATE**

> Changes from this fork were more or less merged to the upstream repository (https://github.com/mcuadros/terraform-provider-helm) and the upstream repository is becoming more stable. I'm going to contribute to upstream from now on, so master branch of this repo won't be maintained. You probably will have to modify your state file if you started using this fork, and now will switch to the upstream, although, you can just remove the releases from your state file, since this provider will just create them in the state file, if release exists in your cluster.

This is a quick and dirty fork of original mcuadros/terraform-provider-helm. 

These are the changes made: 

- Rename `helm_chart` to `helm_release` since that is the correct wording used by Helm.
- Add `reuse_values` option to the release.
- Use `dep` instead of `glide` and update Terraform and Helm dependencies.

Tests may not pass, but everything is working fine. This is a quick fork for getting the job done. The changes may will be submitted to the upstream.

Terraform Provider for Helm
===========================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)
<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/source/assets/images/logo-text.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-helm`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-helm
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-helm
$ make build
```

Using the provider
------------------

- [Provider](docs/index.html.md)
- [Resource: helm_chart](docs/chart.html.md)
- [Resource: helm_repository](docs/repository.html.md)

Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-helm
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
