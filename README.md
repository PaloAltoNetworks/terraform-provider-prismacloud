Terraform Provider for Palo Alto Networks Prisma Cloud
======================================================

<img src="https://raw.githubusercontent.com/hashicorp/terraform-website/master/public/img/logo-hashicorp.svg" width="600px">

- Website: https://www.terraform.io
- Documentation: https://www.terraform.io/docs/providers/prismacloud/index.html
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.11.x or higher
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-prismacloud`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-prismacloud
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-prismacloud
$ make build
```

Using the provider
------------------

If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

See the [Palo Alto Networks Prisma Cloud Provider documentation](https://www.terraform.io/docs/providers/prismacloud/index.html) to get started using the provider.

Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly set up a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-prismacloud
...
```

To test the provider, you can run `make test`.

```sh
$ make test
```

To run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources and often cost money to run.

```sh
$ make testacc
```
