# Tinode Instant Messaging Server

<img src="docs/logo.svg" align="left" width=128 height=128> Instant messaging full stack. Backend in pure [Go](http://golang.org) (license [GPL 3.0](http://www.gnu.org/licenses/gpl-3.0.en.html)), clients for Android (Java), iOS (Swift), and web (ReactJS), as well as [gRPC](https://grpc.io/) client support for C++, C#, Go, Java, Node, PHP, Python, Ruby, Objective-C, etc (all clients licensed under [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0)). Wire transport is JSON over websocket (long polling is also available) or [protobuf](https://developers.google.com/protocol-buffers/) with gRPC.

This is beta-quality software: feature-complete and stable but probably with a few bugs or missing features. Follow [instructions](INSTALL.md) to install and run or use one of the cloud services below. Read [API documentation](docs/API.md).

Tinode is *not* XMPP/Jabber. It is *not* compatible with XMPP. It's meant as a replacement for XMPP. On the surface, it's a lot like open source WhatsApp or Telegram.

<a href="https://apps.apple.com/us/app/tinode/id1483763538"><img src="docs/app-store.svg" height=36></a> <a href="https://play.google.com/store/apps/details?id=co.tinode.tindroidx"><img src="docs/play-store.svg" height=36></a> <a href="https://web.tinode.co/"><img src="docs/web-app.svg" height=36></a>

> **Personal fork note:** I'm using this project to learn Go and explore self-hosted messaging. My local setup notes are in [docs/local-setup.md](docs/local-setup.md).
>
> **Tested with:** Go 1.21, MongoDB 6.0, running on Ubuntu 22.04. See [docs/local-setup.md](docs/local-setup.md) for the exact steps I followed to get a working local instance.

## Why?

The promise of [XMPP](http://xmpp.org/) was to deliver federated instant messaging: anyone would be able to spin up an IM server capable of exchanging messages with any other XMPP server in the world. Unfortunately, XMPP never delivered on this promise. Instant messengers are still a bunch of incompatible walled gardens, similar to what AoL of the late 1990s was to the open Internet.

The goal of this project is to deliver on XMPP's original vision: create a modern open platform for federated instant messaging with an emphasis on mobile communication. A secondary goal is to create a decentralized IM platform that is much harder to track and block by the governments.

An explicit NON-goal: we are not building yet another Slack replacement.

## Installing and running

See [general instructions](./INSTALL.md) or [docker-specific instructions](./docker/README.md).

## Getting support

* Read [API documentation](docs/API.md) and [FAQ](docs/faq.md). Read configuration instructions contained in the [`tinode.conf`](./server/tinode.conf) file.
* For support, general questions, discussions post to [https://groups.google.com/d/forum/tinode](https://groups.google.com/d/forum/tinode).
* For bugs and feature requests [open an issue](https://github.com/tinode/chat/issues/new/choose).
* Use https://tinode.co/contact for commercial inquiries.

## Helping out

* If you appreciate our work, please help spread the word! Sharing on Reddit, HN, and other communities
