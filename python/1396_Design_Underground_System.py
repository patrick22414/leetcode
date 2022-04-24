from dataclasses import dataclass
from collections import defaultdict
from typing import Tuple, Dict


@dataclass
class StartRecord:
    name: str
    t: int


@dataclass
class IntervalStats:
    totalTime: int = 0
    count: int = 0


class UndergroundSystem:
    def __init__(self):
        self.ongoing: Dict[int, StartRecord] = defaultdict(StartRecord)
        self.stats: Dict[Tuple[str, str], IntervalStats] = defaultdict(IntervalStats)

    def checkIn(self, id: int, startStation: str, t: int) -> None:
        self.ongoing[id] = StartRecord(startStation, t)

    def checkOut(self, id: int, endStation: str, t: int) -> None:
        start = self.ongoing[id]
        interval = self.stats[(start.name, endStation)]

        interval.totalTime += t - start.t
        interval.count += 1

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        interval = self.stats[(startStation, endStation)]
        return interval.totalTime / interval.count


# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)
