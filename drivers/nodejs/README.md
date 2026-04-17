nodejs-pg-age
===========



Example
-----
Initialize on make connection

```typescript
import {types, Client, QueryResultRow} from "pg";
import {setAGETypes} from "../src";

const config = {
    user: 'postgres',
    host: '127.0.0.1',
    database: 'postgres',
    password: 'postgres',
    port: 25432,
}

const client = new Client(config);
await client.connect();
await setAGETypes(client, types);

await client.query(`SELECT create_graph('age-first-time');`);
```

Query

```typescript
await client?.query(`
    SELECT *
    from cypher('age-first-time', $$
        CREATE (a:Part {part_num: '123'}),
            (b:Part {part_num: '345'}),
            (c:Part {part_num: '456'}),
            (d:Part {part_num: '789'})
    $$) as (a agtype);
`)

const results: QueryResultRow = await client?.query<QueryResultRow>(`
    SELECT *
    from cypher('age-first-time', $$
        MATCH (a) RETURN a
    $$) as (a agtype);
`)!
```
### TLS / Encrypted Connections

By default, the `pg` package does not enforce TLS.  Always use an encrypted
connection when the database is not on localhost.

Pass `ssl` options in the `Client` config object:

```typescript
import { Client } from "pg";
import { setAGETypes } from "../src";

// Recommended: verify-full validates the server certificate and hostname
const client = new Client({
    user: "app",
    host: "db.example.com",
    database: "postgres",
    password: "secret",
    port: 5432,
    ssl: {
        rejectUnauthorized: true,   // enforce certificate validation
        // ca: fs.readFileSync("/path/to/server-ca.pem").toString(),
        // cert: fs.readFileSync("/path/to/client-cert.pem").toString(),
        // key: fs.readFileSync("/path/to/client-key.pem").toString(),
    },
});
await client.connect();
await setAGETypes(client, types);
```

For a connection string, append `?sslmode=require` (or `verify-full`) to the
PostgreSQL URL and set `ssl: { rejectUnauthorized: true }` in the config.

**Recommended ssl settings:**

| Setting | Meaning |
|---|---|
| `ssl: { rejectUnauthorized: false }` | Encrypts but does not verify the server certificate (not recommended for production) |
| `ssl: { rejectUnauthorized: true }` | Encrypts and verifies the server certificate (recommended) |

### For more information about [Apache AGE](https://age.apache.org/)
* Apache Age : https://age.apache.org/
* GitHub : https://github.com/apache/age
* Document : https://age.apache.org/age-manual/master/index.html
