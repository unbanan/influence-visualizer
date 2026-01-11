package unit

import (
	"encoding/json"
	"testing"
	"time"

	proto "contest-influence/proto/simulation"
	"contest-influence/server/internal/simulation_types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulation_MarshalJSON(t *testing.T) {
	queuedAt := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	startedAt := time.Date(2024, 1, 1, 10, 1, 0, 0, time.UTC)
	finishedAt := time.Date(2024, 1, 1, 10, 5, 0, 0, time.UTC)

	testSimulation := &simulation_types.Simulation{
		Map: simulation_types.Map{
			Data: proto.Map{
				Nrows: 5,
				Ncols: 5,
				FieldMask: []bool{
					true, true, true, true, true,
					true, false, false, false, true,
					true, false, true, false, true,
					true, false, false, false, true,
					true, true, true, true, true,
				},
				BigCells: []*proto.Position{
					{Row: 1, Col: 1},
					{Row: 3, Col: 3},
				},
				StartPositions: []*proto.Position{
					{Row: 0, Col: 0},
					{Row: 4, Col: 4},
				},
			},
			Name: "TestMap",
		},
		Data: proto.Simulation{
			Rounds: []*proto.Round{
				{
					Attack: &proto.AttackSimulationPhase{
						Results: []*proto.AttackResult{
							{
								From: &proto.Cell{
									Pos:   &proto.Position{Row: 0, Col: 0},
									Value: 10,
								},
								To: &proto.Cell{
									Pos:   &proto.Position{Row: 1, Col: 1},
									Value: 5,
								},
								Win: true,
							},
						},
					},
					Defense: &proto.DefenseSimulationPhase{
						Moves: &proto.Cells{
							Cells: []*proto.Cell{
								{
									Pos:   &proto.Position{Row: 2, Col: 2},
									Value: 8,
								},
							},
						},
					},
				},
			},
			Statistics: []*proto.Statistics{
				{
					Players: map[int64]*proto.PlayerStatistics{
						1: {Score: 100},
						2: {Score: 50},
					},
				},
				{
					Players: map[int64]*proto.PlayerStatistics{
						1: {Score: 120},
						2: {Score: 60},
					},
				},
			},
		},
		Players:    []string{"Player1", "Player2"},
		QueuedAt:   queuedAt,
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
		State:      "completed",
	}

	// Marshal to JSON
	jsonBytes, err := json.Marshal(testSimulation)
	require.NoError(t, err)

	// Unmarshal to verify the structure
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	require.NoError(t, err)

	// Verify the structure
	require.Contains(t, result, "map")
	require.Contains(t, result, "data")
	require.Contains(t, result, "players")

	// Verify players
	players := result["players"].([]interface{})
	assert.Len(t, players, 2)
	assert.Equal(t, "Player1", players[0])
	assert.Equal(t, "Player2", players[1])

	// Verify state
	assert.Equal(t, "completed", result["state"])

	// Verify timestamps
	assert.Contains(t, result, "queued_at")
	assert.Contains(t, result, "starterd_at")
	assert.Contains(t, result, "finished_at")

	queuedAtStr := result["queued_at"].(string)
	assert.NotEmpty(t, queuedAtStr)
	parsedQueuedAt, err := time.Parse(time.RFC3339, queuedAtStr)
	require.NoError(t, err)
	assert.Equal(t, queuedAt.Unix(), parsedQueuedAt.Unix())

	// Verify map structure
	mapData := result["map"].(map[string]interface{})
	assert.Equal(t, "TestMap", mapData["name"])
	mapInnerData := mapData["map"].(map[string]interface{})
	assert.Equal(t, float64(5), mapInnerData["nrows"])
	assert.Equal(t, float64(5), mapInnerData["ncols"])

	// Verify data structure
	data := result["data"].(map[string]interface{})
	require.Contains(t, data, "rounds")
	require.Contains(t, data, "statistics")

	// Verify rounds
	rounds := data["rounds"].([]interface{})
	assert.Len(t, rounds, 1)

	// Verify statistics
	statistics := data["statistics"].([]interface{})
	assert.Len(t, statistics, 2)
}

func TestSimulation_MarshalJSON_New(t *testing.T) {
	// Test with empty simulation
	testSimulation := simulation_types.NewSimulation()

	// Marshal to JSON
	jsonBytes, err := json.Marshal(testSimulation)
	require.NoError(t, err)

	// Unmarshal to verify the structure
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	require.NoError(t, err)

	// Verify the structure
	require.Contains(t, result, "map")
	require.Contains(t, result, "data")
	require.Contains(t, result, "players")

	// Verify empty players
	players := result["players"].([]interface{})
	assert.Len(t, players, 0)

	// Verify state field exists
	assert.Contains(t, result, "state")

	// Verify timestamp fields exist
	assert.Contains(t, result, "queued_at")
	assert.Contains(t, result, "starterd_at")
	assert.Contains(t, result, "finished_at")

	// Verify empty rounds and statistics
	data := result["data"].(map[string]interface{})
	assert.Len(t, data, 0)
}
