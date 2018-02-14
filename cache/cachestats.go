package cache

type CacheStats struct {
	hitCount  int64
	missCount int64

	avgHitTime         int64
	hitTimeSampleCount int

	avgMissTime         int64
	missTimeSampleCount int
}

func (c *CacheStats) recordHit() {
	c.hitCount++
}

func (c *CacheStats) recordMiss() {
	c.missCount++
}

func (c *CacheStats) recordHitTimeSample(t int64) {
	//panic(1)
	//recover()
	//c.avgHitTime = (c.avgHitTime * c.hitTimeSampleCount + t) / (c.hitTimeSampleCount + 1 )
	c.hitTimeSampleCount++
}

func (c *CacheStats) recordMissTimeSample(t int64) {
	//panic(1)
	//recover()
	//c.avgMissTime = (c.avgMissTime * c.missTimeSampleCount + t) / (c.missTimeSampleCount + 1 )
	c.missTimeSampleCount++
}
