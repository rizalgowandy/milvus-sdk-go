// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entity

// CompactionState is enum of compaction execution state
// should match definition in common.proto
type CompactionState int32

const (
	//COMPACTION_STATE_UNDEFINED zero value placeholder
	COMPACTION_STATE_UNDEFINED CompactionState = 0
	//COMPACTION_STATE_EXECUTING compaction in progress
	COMPACTION_STATE_EXECUTING CompactionState = 1
	// COMPACTION_STATE_COMPLETED compcation done
	COMPACTION_STATE_COMPLETED CompactionState = 2
)

type CompactionPlanType int32

// See https://wiki.lfaidata.foundation/display/MIL/MEP+16+--+Compaction
const (
	// COMPACTION_PLAN_UNDEFINE zero value placeholder
	COMPACTION_PLAN_UNDEFINE CompactionPlanType = 0
	// COMPACTION_PLAN_APPLY_DELETE apply delete log
	COMPACTION_PLAN_APPLY_DELETE CompactionPlanType = 1
	// COMPACTION_PLAN_MERGE_SEGMENTS merge multiple segments
	COMPACTION_PLAN_MERGE_SEGMENTS CompactionPlanType = 2
)

// CompactionMergePlan contains compaction plan of merging multple segments
type CompactionPlan struct {
	Source   []int64
	Target   int64
	PlanType CompactionPlanType
}
