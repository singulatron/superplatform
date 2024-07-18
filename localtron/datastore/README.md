# DataStore

DataStore is the database interface of Localtron. It's an interface so Localtron can be used with any database (provided there is a DataStore interface implementation for it).

The current implementation (Localstore) is crufty, but an SQLite implementation for a more reliable local solution is shipping soon.
An implementations for PostgreSQL exists but will be published in a different repo soon. Other databases for distributed setups are in the works as well.
