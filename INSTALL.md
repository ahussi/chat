# Installing Tinode

The config file [`tinode.conf`](./server/tinode.conf) contains extensive instructions on configuring the server.

## Installing from Binaries

1. Visit the [Releases page](https://github.com/tinode/chat/releases/), choose the latest or otherwise the most suitable release. From the list of binaries download the one for your database (supported: MySQL, PostgreSQL, MongoDB, RethinkDB) and platform (Linux ARM or Intel, Windows, Mac ARM or Intel). Once the binary is downloaded, unpack it to a directory of your choosing, `cd` to that directory.

2. Make sure your database is running. Make sure it's configured to accept connections from `localhost`. In case of MySQL, Tinode will try to connect as `root` without the password. In case of PostgreSQL, Tinode will try connect as `postgres` with the password `postgres`. See notes below (_Building from Source_, section 4) on how to configure Tinode to use a different user or a password. MySQL 5.7 or above is required (use InnoDB, not MyISAM storage engine). MySQL 5.6 or below **will not work**, use of MyISAM **will cause problems**. PostgreSQL 13 or above is required. PostgreSQL 12 or below **will not work**. MongoDB 4.4 or above is required. MongoDB 4.2 and below **will not work**.

3. Run the database initializer `init-db` (or `init-db.exe` on Windows):
	```
	./init-db -data=data.json
	```

4. Run the `tinode` (or `tinode.exe` on Windows) server. It will work without any parameters.
	```
	./tinode
	```

5. Test your installation by pointing your browser to http://localhost:6060/


## Docker

See [instructions](./docker/README.md)


## Building from Source

1. Install [Go environment](https://golang.org/doc/install). The installation instructions below are for Go 1.18 and newer. Building with the latest Go environment is recommended.

2. OPTIONAL only if you intend to modify the code: Install [protobuf](https://developers.google.com/protocol-buffers/) and [gRPC](https://grpc.io/docs/languages/go/quickstart/) including [code generator](https://developers.google.com/protocol-buffers/docs/reference/go-generated) for Go.

3. Make sure one of the following databases is installed and running:
 * MySQL 5.7 or above configured with `InnoDB` engine (8.x preferred). MySQL 5.6 or below **will not work**.
 * PostgreSQL 13 or above. PostgreSQL 12 or below **will not work**.
 * MongoDB 4.4 or above (8.x preferred). MongoDB 4.2 and below **will not work**.
 * RethinkDB (deprecated, support will be dropped in 2027 unless RethinkDB team resumes development).

> **Personal note:** I use PostgreSQL 15 locally. If you're on macOS, `brew install postgresql@15` works well and the default `postgres` user/password combo matches what Tinode expects out of the box.

4. Fetch, build Tinode server and tinode-db database initializer:
  - **MySQL**:
	```
	go install -tags mysql github.com/tinode/chat/server@latest
	go install -tags mysql github.com/tinode/chat/tinode-db@latest
	```
  - **PostgreSQL**:
	```
	go install -tags postgres github.com/tinode/chat/server@latest
	go install -tags postgres github.com/tinode/chat/tinode-db@latest
	```
  - **MongoDB**:
	```
	go install -tags mongodb github.com/tinode/chat/server
```
