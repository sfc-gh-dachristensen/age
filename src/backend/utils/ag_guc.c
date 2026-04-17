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

#include "postgres.h"

#include <limits.h>

#include "utils/guc.h"
#include "utils/ag_guc.h"

bool  age_enable_containment = true;
int   age_max_vle_depth = 1000;
int   age_graph_load_size_limit = 0;
int   age_vle_edge_state_limit = 0;
int   age_vle_cache_max_entries = 0;
char *age_csv_directory = NULL;

/*
 * Defines AGE's custom configuration parameters.
 *
 * The name of the parameter must be `age.*`. This name is used for setting
 * value to the parameter. For example, `SET age.enable_containment = on;`.
 */
void define_config_params(void)
{
    DefineCustomBoolVariable("age.enable_containment",
                             "Use @> operator to transform MATCH's filter. Otherwise, use -> operator.",
                             NULL,
                             &age_enable_containment,
                             true,
                             PGC_SUSET,
                             0,
                             NULL,
                             NULL,
                             NULL);

    DefineCustomIntVariable("age.max_vle_depth",
                            "Limit recursion in vle detection to the given number of iterations.",
                            NULL,
                            &age_max_vle_depth,
                            1000,
                            1,
                            INT_MAX,
                            PGC_SUSET,
                            0,
                            NULL,
                            NULL,
                            NULL);

    DefineCustomIntVariable("age.graph_load_size_limit",
                            "Maximum number of vertices plus edges loaded into a global graph context (0 = no limit).",
                            NULL,
                            &age_graph_load_size_limit,
                            0,
                            0,
                            INT_MAX,
                            PGC_SUSET,
                            0,
                            NULL,
                            NULL,
                            NULL);

    DefineCustomIntVariable("age.vle_edge_state_limit",
                            "Maximum number of entries in the VLE edge state hash table per traversal (0 = no limit).",
                            NULL,
                            &age_vle_edge_state_limit,
                            0,
                            0,
                            INT_MAX,
                            PGC_SUSET,
                            0,
                            NULL,
                            NULL,
                            NULL);

    DefineCustomIntVariable("age.vle_cache_max_entries",
                            "Maximum total edge-state entries across all cached VLE contexts (0 = count limit only).",
                            NULL,
                            &age_vle_cache_max_entries,
                            0,
                            0,
                            INT_MAX,
                            PGC_SUSET,
                            0,
                            NULL,
                            NULL,
                            NULL);

    DefineCustomStringVariable("age.csv_directory",
                               "Directory from which CSV files may be loaded (must end with /).",
                               "Only files whose realpath() resolves inside this directory are "
                               "permitted. Set this to a directory owned by the PostgreSQL service "
                               "account with mode 0700 to prevent other local users from staging "
                               "malicious files. The default /tmp/age/ is world-writable.",
                               &age_csv_directory,
                               "/tmp/age/",
                               PGC_SUSET,
                               0,
                               NULL,
                               NULL,
                               NULL);

    EmitWarningsOnPlaceholders("age");
}
