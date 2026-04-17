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

#ifndef AG_GUC_H
#define AG_GUC_H

/*
 * AGE configuration parameters.
 *
 * Ideally, these parameters should be documented in a .sgml file.
 *
 * To add a new parameter, add a global variable. Add its definition
 * in the `define_config_params` function. Include this header file
 * to use the global variable. The parameters can be set just like
 * regular Postgres parameters. See guc.h for more details.
 */

/*
 * If set true, MATCH's property filter is transformed into the @>
 * (containment) operator. Otherwise, the -> operator is used. The former case
 * is useful when GIN index is desirable, the latter case is useful for Btree
 * expression index.
 */
extern bool age_enable_containment;

/*
 * Set the maximum recursion depth for vle graph traversal.
 */
extern int age_max_vle_depth;

/*
 * Maximum number of vertices plus edges that may be loaded into a single
 * global graph context.  0 means no limit (the default).  Set this to a
 * positive value to prevent a backend from exhausting memory on very large
 * graphs.  Example: SET age.graph_load_size_limit = 10000000;
 */
extern int age_graph_load_size_limit;

/*
 * Maximum number of entries allowed in the VLE edge state hash table per
 * traversal context.  0 means no limit (the default).  Set this to a
 * positive value to cap memory consumed by large VLE traversals.
 * Example: SET age.vle_edge_state_limit = 1000000;
 */
extern int age_vle_edge_state_limit;

/*
 * Maximum total number of edge-state entries that may be held across all
 * cached VLE local contexts.  0 means only the context-count limit applies
 * (the default).  Setting a positive value evicts the least-recently-used
 * cached contexts when the combined edge-state entry count would exceed this
 * threshold, providing a memory-proportional cache bound.
 * Example: SET age.vle_cache_max_entries = 5000000;
 */
extern int age_vle_cache_max_entries;

void define_config_params(void);

#endif
