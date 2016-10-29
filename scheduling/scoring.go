package scheduling

// Assumes OrderClasses has been called on the combo
func ScoreCombo(combo Combo, criteria Criteria) int {
	return ScoreBreaks(combo, criteria) +
		ScoreProfs(combo, criteria) +
		ScoreEarliestClass(combo, criteria) +
		ScoreLatestClass(combo, criteria) +
		ScoreDays(combo, criteria)
}

//note: I don't really know how to do these, so some of them may have some inherent bias
func ScoreBreaks(combo Combo, criteria Criteria) int {
	output := 0
	minutes := 0
	for i := 0; i < len(combo.Classes)-1; i++ {
		minutes = MinuteDiff(combo.Classes[i].EndTime, combo.Classes[i+1].StartTime)
		if minutes > 15 { //this promotes a 15 minute gap between every class
			//maybe later make it a user-configurable option
			if criteria.Breaks.Maximize {
				output += (minutes - 15)
			} else {
				output -= (minutes - 15)
			}
		} else {
			output -= (15 - minutes)
		}
	}
	return output * criteria.Breaks.Weight
}

func ScoreProfs(combo Combo, criteria Criteria) int {
	//will have to do some sort of query to rateMyProfessor
	return 0 //temp
}

func ScoreEarliestClass(combo Combo, criteria Criteria) int {
	output := 0
	earliestStart := combo.Classes[0].StartTime
	earliestWanted := criteria.EarliestClass.Time
	diff := earliestStart.Sub(earliestWanted)
	var minutes = diff.Minutes()
	if earliestStart.After(earliestWanted) {
		output += int(minutes)
	} else {
		if criteria.EarliestClass.Manditory {
			return -100000
		} else {
			output -= int(minutes)
		}
	}
	output *= criteria.EarliestClass.Weight
	return output
}

func ScoreLatestClass(combo Combo, criteria Criteria) int {
	output := 0
	latestEnd := combo.Classes[len(combo.Classes)-1].EndTime
	latestWanted := criteria.LatestClass.Time
	diff := latestEnd.Sub(latestWanted)
	var minutes = diff.Minutes()
	if latestEnd.Before(latestWanted) {
		output += int(minutes)
	} else {
		if criteria.EarliestClass.Manditory {
			return -100000
		} else {
			output -= int(minutes)
		}
	}

	return output * criteria.EarliestClass.Weight
}

// TODO: Could use some refactoring
func ScoreDays(combo Combo, criteria Criteria) int {
	output := 0
	for _, class := range combo.Classes {
		for _, ClassDay := range class.MeetingDays {
			for _, CriteriaDay := range criteria.Days.Other {
				if ClassDay == CriteriaDay {
					if criteria.Days.Manditory {
						output = -10000000
					} else {
						output -= 10
					}
				}
			}
		}
	}
	output *= criteria.Days.Weight
	return output
}

