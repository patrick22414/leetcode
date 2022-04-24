package main

func main() {

}

type StartRecord struct {
	name string
	t    int
}

type Interval struct {
	start string
	end   string
}

type IntervalStats struct {
	totalTime int
	count     uint
}

type UndergroundSystem struct {
	ongoing map[int]StartRecord
	stats   map[Interval]IntervalStats
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		ongoing: make(map[int]StartRecord),
		stats:   make(map[Interval]IntervalStats),
	}
}

func (tube *UndergroundSystem) CheckIn(id int, startStation string, t int) {
	tube.ongoing[id] = StartRecord{startStation, t}
}

func (tube *UndergroundSystem) CheckOut(id int, endStation string, t int) {
	start := tube.ongoing[id]
	interval := Interval{start.name, endStation}
	stats := tube.stats[interval]

	stats.totalTime += t - start.t
	stats.count += 1

	tube.stats[interval] = stats
}

func (tube *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	stats := tube.stats[Interval{startStation, endStation}]
	return float64(stats.totalTime) / float64(stats.count)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
