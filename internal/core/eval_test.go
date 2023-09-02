package core

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestEvalGEOADD(t *testing.T) {
	res, err := Decode(evalGEOADD([]string{"vn", "10", "20", "p1"}))
	assert.Nil(t, err)
	assert.EqualValues(t, res, 1)

	res, err = Decode(evalGEOADD([]string{"vn", "10", "20", "p1"}))
	assert.Nil(t, err)
	assert.EqualValues(t, res, 0)

	res, err = Decode(evalGEOADD([]string{"vn", "-10", "20", "p1"}))
	assert.Nil(t, err)
	assert.EqualValues(t, res, 1)

	res, err = Decode(evalGEOADD([]string{"vn", "-10", "20", "p2", "-1", "2", "p3"}))
	assert.Nil(t, err)
	assert.EqualValues(t, res, 2)

	zset, exist := zsetStore["vn"]
	assert.True(t, exist)
	assert.EqualValues(t, 3, zset.Len())

	res, err = Decode(evalGEOADD([]string{"vn"}))
	assert.EqualValues(t, "(error) ERR wrong number of arguments for 'GEOADD' command", res)
	res, err = Decode(evalGEOADD([]string{"vn", "-10", "20", "p4", "20"}))
	assert.EqualValues(t, "(error) ERR wrong number of arguments for 'GEOADD' command", res)
}

func TestEvalGEODIST(t *testing.T) {
	evalGEOADD([]string{"vn", "20", "10", "p1"})
	evalGEOADD([]string{"vn", "40", "30", "p2"})
	evalGEOADD([]string{"vn", "10", "85", "p3"})
	evalGEOADD([]string{"vn", "10", "-85", "p4"})
	evalGEOADD([]string{"vn", "180", "20", "p5"})
	evalGEOADD([]string{"vn", "179.9999", "20", "p6"})
	res, err := Decode(evalGEODIST([]string{"vn", "p1", "p2"}))
	assert.Nil(t, err)
	dist, err := strconv.ParseFloat(res.(string), 64)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(dist-3041460.716138), 1.0)

	res, err = Decode(evalGEODIST([]string{"vn", "p3", "p4"}))
	assert.Nil(t, err)
	dist, err = strconv.ParseFloat(res.(string), 64)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(dist-18908471), 1.0)

	res, err = Decode(evalGEODIST([]string{"vn", "p5", "p6"}))
	assert.Nil(t, err)
	dist, err = strconv.ParseFloat(res.(string), 64)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(dist-10.451853), 1.0)

	res, err = Decode(evalGEODIST([]string{"vn", "p1", "p2", "km"}))
	assert.Nil(t, err)
	dist, err = strconv.ParseFloat(res.(string), 64)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(dist-3041), 1.0)
}
