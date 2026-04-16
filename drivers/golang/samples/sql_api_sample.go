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
	"database/sql"
	"fmt"

	"github.com/apache/age/drivers/golang/age"
)

// Do cypher query to AGE with database/sql Tx API transaction control
func doWithSqlAPI(dsn string, graphName string) {

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Confirm graph_path created
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	// Tx begin for execute create vertex
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Create vertices with Cypher
	_, err = age.ExecCypher(tx, graphName, 0,
		"CREATE (n:Person {name: $name, weight: $weight})",
		map[string]interface{}{"name": "Joe", "weight": 67.3})
	if err != nil {
		panic(err)
	}

	_, err = age.ExecCypher(tx, graphName, 0,
		"CREATE (n:Person {name: $name, weight: 77.3, roles: ['Dev','marketing']})",
		map[string]interface{}{"name": "Jack"})
	if err != nil {
		panic(err)
	}

	_, err = age.ExecCypher(tx, graphName, 0,
		"CREATE (n:Person {name: $name, weight: $weight})",
		map[string]interface{}{"name": "Andy", "weight": 59})
	if err != nil {
		panic(err)
	}

	// Commit Tx
	tx.Commit()

	// Tx begin for queries
	tx, err = db.Begin()
	if err != nil {
		panic(err)
	}
	// Query cypher
	cypherCursor, err := age.ExecCypher(tx, graphName, 1, "MATCH (n:Person) RETURN n", nil)
	if err != nil {
		panic(err)
	}
	// Unmarshal result data to Vertex row by row
	for cypherCursor.Next() {
		row, err := cypherCursor.GetRow()
		if err != nil {
			panic(err)
		}
		vertex := row[0].(*age.Vertex)
		fmt.Println(vertex.Id(), vertex.Label(), vertex.Props())
	}

	// Create Paths (Edges)
	_, err = age.ExecCypher(tx, graphName, 0,
		"MATCH (a:Person {name: $aname}), (b:Person {name: $bname}) CREATE (a)-[r:workWith {weight: $weight}]->(b)",
		map[string]interface{}{"aname": "Jack", "bname": "Joe", "weight": 3})
	if err != nil {
		panic(err)
	}

	_, err = age.ExecCypher(tx, graphName, 0,
		"MATCH (a:Person {name: $aname}), (b:Person {name: $bname}) CREATE (a)-[r:workWith {weight: $weight}]->(b)",
		map[string]interface{}{"aname": "Joe", "bname": "Andy", "weight": 7})
	if err != nil {
		panic(err)
	}

	tx.Commit()

	tx, err = db.Begin()
	if err != nil {
		panic(err)
	}
	// Query Paths with Cypher
	cypherCursor, err = age.ExecCypher(tx, graphName, 1, "MATCH p=()-[:workWith]-() RETURN p", nil)
	if err != nil {
		panic(err)
	}

	for cypherCursor.Next() {
		row, err := cypherCursor.GetRow()
		if err != nil {
			panic(err)
		}

		path := row[0].(*age.Path)
		vertexStart := path.GetAsVertex(0)
		edge := path.GetAsEdge(1)
		vertexEnd := path.GetAsVertex(2)

		fmt.Println(vertexStart, edge, vertexEnd)
	}

	// Query with return many columns
	cursor, err := age.ExecCypher(tx, graphName, 3, "MATCH (a:Person)-[l:workWith]-(b:Person) RETURN a, l, b", nil)
	if err != nil {
		panic(err)
	}

	count := 0
	for cursor.Next() {
		row, err := cursor.GetRow()
		if err != nil {
			panic(err)
		}
		count++
		v1 := row[0].(*age.Vertex)
		edge := row[1].(*age.Edge)
		v2 := row[2].(*age.Vertex)
		fmt.Println("ROW ", count, ">>", "\n\t", v1, "\n\t", edge, "\n\t", v2)
	}

	// Delete Vertices
	_, err = age.ExecCypher(tx, graphName, 0, "MATCH (n:Person) DETACH DELETE n RETURN *", nil)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}
