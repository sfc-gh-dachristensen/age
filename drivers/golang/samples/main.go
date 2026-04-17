/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// buildDSN constructs a connection string from standard PG* environment
// variables.  Never hardcode passwords in source — set them in the
// environment instead (e.g. export PGPASSWORD=...).
func buildDSN() string {
	host := os.Getenv("PGHOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("PGPORT")
	if port == "" {
		port = "5432"
	}
	dbname := os.Getenv("PGDATABASE")
	if dbname == "" {
		dbname = "postgres"
	}
	user := os.Getenv("PGUSER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("PGPASSWORD")
	sslmode := os.Getenv("PGSSLMODE")
	if sslmode == "" {
		sslmode = "require"
	}
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, password, sslmode)
}

var dsn string = buildDSN()

// var graphName string = "{graph_path}"
var graphName string = "testGraph"

func main() {

	// Do cypher query to AGE with database/sql Tx API transaction control
	fmt.Println("# Do cypher query with SQL API")
	doWithSqlAPI(dsn, graphName)

	// Do cypher query to AGE with Age API
	fmt.Println("# Do cypher query with Age API")
	doWithAgeWrapper(dsn, graphName)
}
