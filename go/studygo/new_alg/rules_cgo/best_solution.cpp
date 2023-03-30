#include "best_solution.h"

#include <cassert>
#include <iostream>
#include <algorithm>

using namespace std;

namespace OKey {
void dump(const Card &card){
    cout<<"{"<<card.first<<","<<card.second<<"}"<<endl;
}

void dump(const Deck &cards)
{
	cout<<"{";
	bool first = true;
	for ( const auto &card : cards )
	{
		cout<<(first?"":","); first = false;
		cout<<"{"<<card.first<<","<<card.second<<"}";
	}
	cout<<"}";
}

void dump(const Play &play)
{
	dump(Play2Deck(play));
}

void printPlay(const Play &play){
    cout<<"{"<<play.count<<" "<<play.min<<" "<<play.max<<" "<<play.interval<<" "<<play.suit<<" [";
    for(int i=0; i<5;i++){
        cout<<play.suits[i]<<" ";
    }
    cout<<"]"<<play.wild<<"}"<<endl;
}

void dump(const Solution &solution)
{
	bool first = true;
	for ( const auto &play : solution.plays )
	{
		cout<<(first?"":","); first = false;
		dump(play);
	}
	cout<<", "<<solution.okey<<" okey"<<endl;
}

void dump(const vector<Deck> &decks)
{
	cout<<"{";
	bool first = true;
	for ( const auto &deck : decks )
	{
		cout<<(first?"":","); first = false;
		dump(deck);
	}
	cout<<"}";
	cout<<endl;
}


int OKeySolver::MakeBestPlay(Play &play, int okey)
{
	int score = Score(play);
	if ( play.count > 1 )
	{
		assert(okey <= 1);
		okey = (okey >= 1) ? 1 : 0; // hard code for cut down okey
		if ( play.interval ) // Sequence
		{
			int add_to_max = min(okey, MAX_N - play.max);
			int add_to_min = min(okey - add_to_max, play.min - 1);
			// if ( 1 == okey)
			{
				if ( play.max - play.min + 1 >= 13 ) // full meld
				{
					add_to_min = 0;
					add_to_max = 0;
				}
				else
				if ( 2 == play.min && play.max < 13 ) // ace to min
				{
					add_to_min = 1;
					add_to_max = 0;
				}
				// else
				// if ( 2 < play.min && play.max == 13 ) // ace to max
				// {
				// 	add_to_min = 0;
				// 	add_to_max = 1;
				// }
			}
			if ( play.wild ) // not allow use wild card twice
			{
				if ( 0 == card_counts[CARD_OFFSET(play.min - add_to_min, play.suit)] )
					add_to_min = 0;
				if ( 0 == card_counts[CARD_OFFSET(play.max + add_to_max, play.suit)] )
					add_to_max = 0;
			}
			play.min -= add_to_min;
			play.max += add_to_max;
			okey = add_to_min + add_to_max;
			assert(!(1==play.min&&14==play.max)); // cannot both ace
		}
		else // Triplet
		{
			okey = min(okey, MAX_M - play.count);
		}
		play.count += okey;
		play.suits[0] += okey;
	}
	else
	if ( play.count == 1 && okey >= 2 )
	{
		int add_to_max = min(okey, MAX_N - play.max);
		int add_to_min = min(okey - add_to_max, play.min - 1);
		if ( add_to_max > add_to_min ) // to Sequence
		{
			play.min -= add_to_min;
			play.max += add_to_max;
			okey = add_to_min + add_to_max;
		}
		else // to Triplet
		{
			play.interval = 0;
			okey = min(okey, MAX_M - 1);
		}
		play.count += okey;
		play.suits[0] += okey;
	}
	return Score(play) - score;
}
void OKeySolver::Evaluate(const Solution &solution)
{
	// //dump(solution);

	struct BestScore
	{
		int index;
		Play play;
		int inc;
	};
	struct BestCount
	{
		int index;
		Play play;
	};

	if ( 2 <= solution.okey )
	{
		bool has_2_cards_play = false;

		// 1 okey and 1 okey
		{
			BestScore best_score_1, best_score_2;
			best_score_1.index = best_score_2.index = -1;
			best_score_1.inc = best_score_2.inc = 0;
			BestCount best_count_1, best_count_2;
			best_count_1.index = best_count_2.index = -1;
			for ( int i = solution.plays.size()-1; i >= 0; --i )
			{
				Play play = solution.plays[i];
				int score_inc = MakeBestPlay(play, 1);
				if ( score_inc > best_score_1.inc )
				{
					best_score_2 = best_score_1;
					best_score_1.index = i;
					best_score_1.play = play;
					best_score_1.inc = score_inc;
				}
				else
				if ( score_inc > best_score_2.inc )
				{
					best_score_2.index = i;
					best_score_2.play = play;
					best_score_2.inc = score_inc;
				}

				if ( 3 == play.count && best_count_2.index < 0 ) // last 3-cards plays are best count plays
				{
					has_2_cards_play = true;
					if ( best_count_1.index < 0 )
					{
						best_count_1.index = i;
						best_count_1.play = play;
					}
					else
					{
						best_count_2.index = i;
						best_count_2.play = play;
					}
				}
			}

			Solution okey_solution(solution);
//			dump(okey_solution);
			if ( best_score_1.index >= 0 )
			{
				okey_solution.plays[best_score_1.index] = best_score_1.play;
				okey_solution.okey -= 1;
			}
			if ( best_score_2.index >= 0 )
			{
				okey_solution.plays[best_score_2.index] = best_score_2.play;
				okey_solution.okey -= 1;
			}
			AddBest(okey_solution);

			if ( ( best_score_1.index != best_count_1.index && best_count_1.index >= 0 ) ||
				( best_score_2.index != best_count_2.index && best_count_2.index >= 0 ) )
			{
				Solution okey_solution_count(solution);
				{
					okey_solution_count.plays[best_count_1.index] = best_count_1.play;
					okey_solution_count.okey -= 1;
				}
				if ( best_count_2.index >= 0 )
				{
					okey_solution_count.plays[best_count_2.index] = best_count_2.play;
					okey_solution_count.okey -= 1;
				}
				else
				{
					okey_solution_count.plays[best_score_1.index] = best_score_1.play;
					okey_solution_count.okey -= 1;
				}
				AddBest(okey_solution_count);
			}
		}

		/*
		// 2 okeys
		{
			BestScore best_score;
			best_score.index = -1;
			best_score.inc = 0;
			BestCount best_count;
			best_count.index = -1;
			for ( int i = solution.plays.size()-1; i >= 0; --i )
			{
				Play play = solution.plays[i];
				int score_inc = MakeBestPlay(play, 2);
				if ( score_inc > best_score.inc )
				{
					best_score.index = i;
					best_score.play = play;
					best_score.inc = score_inc;
				}

				if ( !has_2_cards_play ) // previous separated 2 okeys is better choice when 2-cards play exists
					if ( 3 == play.count && best_count.index < 0 ) // last 3-cards play is best count play
					{
						best_count.index = i;
						best_count.play = play;
					}
			}

			Solution okey_solution(solution);
			if ( best_score.index >= 0 )
			{
				okey_solution.plays[best_score.index] = best_score.play;
				okey_solution.okey -= 2;
			}
			AddBest(okey_solution);

			if ( best_score.index != best_count.index && best_count.index >= 0 )
			{
				Solution okey_solution_count(solution);
				{
					okey_solution_count.plays[best_count.index] = best_count.play;
					okey_solution_count.okey -= 2;
				}
				AddBest(okey_solution_count);
			}
		}
		*/
	}
	else
	if ( 1 == solution.okey )
	{
		// 1 okey
		{
			BestScore best_score;
			best_score.index = -1;
			best_score.inc = 0;
			BestCount best_count;
			best_count.index = -1;
			for ( int i = solution.plays.size()-1; i >= 0; --i )
			{
				Play play = solution.plays[i];
				int score_inc = MakeBestPlay(play, 1);
				if ( score_inc > best_score.inc )
				{
					best_score.index = i;
					best_score.play = play;
					best_score.inc = score_inc;
				}

				if ( 3 == play.count && best_count.index < 0 ) // last 3-cards play is best count play
				{
					best_count.index = i;
					best_count.play = play;
				}
			}

			Solution okey_solution(solution);
			if ( best_score.index >= 0 )
			{
				okey_solution.plays[best_score.index] = best_score.play;
				okey_solution.okey -= 1;
			}
			AddBest(okey_solution);

			if ( best_score.index != best_count.index && best_count.index >= 0 )
			{
				Solution okey_solution_count(solution);
				{
					okey_solution_count.plays[best_count.index] = best_count.play;
					okey_solution_count.okey -= 1;
				}
				AddBest(okey_solution_count);
			}
		}
	}
	else
	{
		AddBest(solution);
	}
}

int OKeySolver::Score(const Play& play)
{
	int score = 0;
	if ( play.count >= 3 )
	{
		// if ( play.interval )
		// 	score = (play.min + play.max) * play.count / 2;
		// else
		// 	score = play.max * play.count;
		if ( play.interval )
			for ( int i = play.min; i <= play.max; ++i )
				score += ScoreMap[i];
		else
			score = ScoreMap[play.max] * play.count;
	}
	return score;
}
int OKeySolver::Score(const Solution &solution)
{
	int score = 0;
	for ( const auto &play : solution.plays )
		score += Score(play);
	return score;
}

int OKeySolver::Count(const Play& play)
{
	if ( play.count >= 3 )
		return play.count;
	return 0;
}

int OKeySolver::Count(const Solution &solution)
{
	int count = 0;
	for ( const auto &play : solution.plays )
		count += Count(play);
	return count;
}

bool OKeySolver::AddBest(const Solution &solution)
{
	bool add_best = false;

	int score = Score(solution);
	int count = Count(solution);

	if ( score > best_score || (score == best_score && count > best_count_in_best_score) )
	{
		best_score = score;
		best_count_in_best_score = count;
//		cout<<"start"<<endl;
//		dump(solution);
//		cout<<"end"<<endl;
		best_score_solution = solution;
		add_best = true;
	}

	if ( count > best_count || (count == best_count && score > best_score_in_best_count) )
	{
		best_count = count;
		best_score_in_best_count = score;
		best_count_solution = solution;
		add_best = true;
	}

	return add_best;
}

Solution OKeySolver::GetBestScoreSolution()
{
	return best_score_solution;
}
Solution OKeySolver::GetBestCountSolution()
{
	return best_count_solution;
}

int OKeySolver::CheckPlay(const Play &play, const Card &card)
{
	if ( play.count > 1 )
	{
		if ( play.interval ) // Sequence
		{
			if ( play.suit == card.second )
			{
				if ( 1 == play.min && 14 == card.first ) // ace twice
					return -1;
				return card.first - play.max - 1;
			}
		}
		else // Triplet
		if ( play.count < MAX_M )
		{
			if ( play.max == card.first && 0 == play.suits[card.second] )
				return 0;
		}
		return -1;
	}
	else
	if ( 1 == play.count )
	{
		if ( play.suit == card.second ) // Sequence
			return card.first - play.max - 1;
		else // Triplet
		if ( play.max == card.first )
			return 0;
		return -1;
	}
	else
	if ( play.count < 0 )
		return -1;
	return 0;
}
//int i = 0;
void OKeySolver::AddPlay(Play &play, const Card &card)
{
//    i++;
    //cout<<endl;
    //cout<<i<<" ";
    //dump(play);
    //cout<<"{"<<card.first<<","<<card.second<<"}"<<endl;
    //cout<<"playf";
//    printPlay(play);
    //cout<<endl;
	if ( 0 == play.count )
	{
	    //cout<<endl;
	    //cout<<"countfirst"<<endl;
		play.count = 1;
		play.min = play.max = card.first;
		play.interval = 1;
		play.suit = card.second;
		++play.suits[play.suit];
	}
	else
	{
		assert(play.count > 0);
		if ( play.max == card.first ) // Triplet
		{
			assert(1 == play.count || !play.interval);
			assert(0 == card.second || 0 == play.suits[card.second]);
			play.interval = 0;
		}
		else  // Sequence
		{
			assert(1 == play.count || play.interval);
			assert(play.max+1 == card.first);
			assert(0 == card.second || play.suit == card.second);
			++play.max;
		}
		++play.suits[card.second];

	    //cout<<endl;
	    //cout<<"countsecond"<<play.count<<endl;
		++play.count;
	}
	//cout<<"//dump play";
//	printPlay(play);
	//cout<<endl;
}
void OKeySolver::RemovePlay(Play &play, const Card &card)
{
	assert(play.count > 0);
	assert(play.max == card.first);
	assert(play.suits[card.second]);
	if ( play.interval ) // Sequence
	{
		assert(0 == card.second || play.suit == card.second);
		--play.max;
	}
	else // Triplet
	{
		if ( 2 == play.count )
			play.interval = 1;
	}
	--play.suits[card.second];
	--play.count;
}

// it is not easy, but it is able to proof for this method.
/*
	(a,b),...,(...,c,...),... is waste

		case: 1,2,3
		proof:
		(1,2,3,...),...,(),... is better than (1,2),...,(3),...

		case: 1,2,3,4,...
		proof:
		(1,2,3,4,...),...,(),... is better than (1,2),...,(3,4,...),...

		case: 1,2,2,3,...
		proof:
		(1,2,3,...),...,(2),... is better than (1,2),...,(2,3,...),...

		case: 1,2,2,3,4,...
		proof:
		(1,2,3,4,...),...,(2),... is better than (1,2),...,(2,3,4,...),...

		case: 1,2,3,3
		proof:
		(1,2,3,...),...,(3),... is better than (1,2),...,(3,3),...

		case: 1,2,3,3,3
		proof:
		(1,2,3,...),...,(3,3),... is bad than (1,2),...,(3,3,3),...
		(1),(2),...,(3,3,3),... is same as (1,2),...,(3,3,3),...

		case: 1,2,3,3,3,3
		proof:
		(1,2,3,...),...,(3,3,3),... is same as (1,2),...,(3,3,3,3),...

	(a,a),...,(...,a,...),... is waste

		case: 1,1,1
		proof:
		(1,1,1),...,(),... is better than (1,1),...,(1),...

		case: 1,1,1,1
		proof:
		(1,1,1),...,(1),... is better than (1,1),...,(1,1),...
		(1,1,1),...,(1),... is better than (1,1),...,(1),(1),...

		case: 1,1,1,2
		proof:
		(1,1,1),...,(2),... is better than (1,1),...,(1,2),...

		case: 1,1,1,2,3
		proof:
		(1,1,1),...,(2,3),... is bad than (1,1),...,(1,2,3),...
		(1),(1),...,(1,2,3),... is same as (1,1),...,(1,2,3),...

		case: 1,1,1,2,3,4,...
		proof:
		(1,1,1),...,(2,3,4,...),... is better than (1,1),...,(1,2,3,4,...),...
*/
bool OKeySolver::IsNextWasteOnSinglePlay(const Play &play, const Card& card, int okey_cost, int okey_remain)
{
	if ( 0 == okey_cost && 2 == play.count // play.count can't be less than 2 and more than 2.
		&& okey_remain <= 1 ) // avoid any play contains more than two okeys
	{
		if ( play.interval ) // Sequence
		{
			/*
				play can be any sequence.
			*/
			assert(play.max + 1 == card.first && play.suit == card.second);
			return true;
		}
		else // Triplet
		{
			assert(play.max == card.first);
			assert(0 == card.second || 0 == play.suits[card.second]);
			/*
				play can't be any triplet.

				the following examples will be deal with.
					case: {{1,1},{1,2},{1,3},{2,3},{3,3}} with 0 okey
					the best solution is {{(1,1)},{(1,2)},{(1,3),(2,3),(3,3)}}, score is 6.

					case: {{1,1},{1,2},{1,3},{2,3},{3,3}} with 1 okey
					the best solution is {{(1,1),(1,2),(1,3)},{(2,3),(3,3),okey}}, score is 12.

					case: {{11,1},{11,2},{11,3},{12,3},{13,3}} with 1 okey
					the best solution is {{(11,1),(11,2),(11,3)},{okey,(12,3),(13,3)}}, score is 69.

					case: {{1,1},{1,2},{1,3},{1,3},{2,3},{3,3}} with 0 okey
					the best solution is {{(1,1),(1,2),(1,3)},{(1,3),(2,3),(3,3)}}, score is 9.

				the following examples will be not deal with.
					case: {{1,1},{1,2},{1,3},{1,4},{2,3},{3,3}} with 0 okey
					the best solution is {{(1,1),(1,2),(1,4)},{(1,3),(2,3),(3,3)}}, score is 9.
			*/
			if ( card_suit_upper_bounds[card.first] == card.second ) // no more other else will be triplet
				return true;
		}
	}
	return false;
}

// it is not easy, but it is able to proof for this method.
/*
	(a,b),(c,d),...,(...,e,...),... is waste

		case: 1,2,3,3
		proof:
		(1,2,3),...,(3),... is better than (1,2),...,(3),...,(3),...

		case: 1,2,3,4,4
		proof:
		(1,2,3,4),...,(4),... is better than (1,2),...,(3,4),...,(4),...

	(a),(a),...,(...,b,...),... is waste

		case: 1,1,2,...
		proof:
		(1,1),...,(2,...),... is same as (1),(1),...,(2,...),...
		(1,2,...),(1),...,(),... is better than (1),(1),...,(2,...),...

		case: 1,1,1,2,...
		proof:
		(1,1,1),...,(2,...),... is better than (1,1),(1),...,(2,...),...
		(1,1,1),...,(2,...),... is better than (1),(1,1),...,(2,...),...
		(1,1,1),...,(2,...),... is better than (1),(1),(1),...,(2,...),...

		case: 1,1,1,1,2,...
		proof:
		(1,1,1),(1),...,(2,...),... is better than (1,1),(1,1),...,(2,...),...
		(1,1,1),(1),...,(2,...),... is better than (1,1),(1),(1),...,(2,...),...
		(1,1,1),(1),...,(2,...),... is better than (1),(1,1),(1),...,(2,...),...
		(1,1,1),(1),...,(2,...),... is better than (1),(1),(1,1),...,(2,...),...
		(1,1,1),(1),...,(2,...),... is better than (1),(1),(1),(1),...,(2,...),...
*/
bool OKeySolver::IsNextWasteOnMultiplePlays(const Play &play, const Card& card, PlayPreference &preference, WasteMark &mark)
{
	PlayPreference play_preference = PlayPreferNone;
	if ( play.count > 1 )
	{
		play_preference = play.interval ? PlayPreferToSequence : PlayPreferToTriplet;
	}
	else
	// if ( 1 == play.count )
	if ( 1 == play.count && play.max > 1 ) // ace card need to retrieve
	{
		assert(1 == play.count);
		if ( play.max == card.first && card_suit_upper_bounds[card.first] == card.second )
			preference = PlayPreferToSequence;
		else
		if ( play.max < card.first && play.suit == card.second )
			preference = PlayPreferToTriplet;
		play_preference = preference;
	}

	switch ( play_preference )
	{
	case PlayPreferToSequence:
		/*
			play can't be any sequence.

			the following examples will be deal with.
				case: {{1,1},{2,1},{3,1},{4,1},{5,1},{6,1},{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1}} with 0 okey
				the best solution is {(1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1)}, score is 91.

				case: {{9,1},{10,1},{11,1},{12,1},{13,1},{11,3},{12,3},{13,3}} with 0 okey
				the best solution is {(9,1),(10,1),(11,1),(12,1),(13,1)},{(11,3),(12,3),(13,3)}, score is 91.

			the following examples will be not deal with.
				case: {{9,1},{10,1},{11,1},{12,1},{13,1},{11,3},{12,3},{13,3}} with 1 okey
				the best solution is {(9,1),(10,1),(11,1)},{(11,3),(12,3),(13,3)},{okey,(12,1),(13,1)}, score is 102.

				case: {{9,1},{10,1},{11,1},{12,1},{13,1},{11,3},{12,3},{13,3}} with 2 okey
				the best solution is {(9,1),(10,1),(11,1),(12,1)},{(11,3),(12,3),(13,3)},{okey,okey,(13,1)}, score is 117.

				case: {{1,1},{2,1},{3,1},{4,1},{5,1},{6,1},{7,1},{8,1},{9,1},{10,1},{11,1},{12,1}} with 2 okey
				the best solution is {(1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1)},{okey,okey,(12,1)}, score is 102.

				case: {{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1},{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1}} with 1 okey
				the best solution is {{(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1)},{(7,1),(8,1),(9,1),(10,1),(11,1)},{okey,(12,1),(13,1)}}, score is 151.

				case: {{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1},{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1}} with 2 okey
				the best solution is {(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1)},{(7,1),(8,1),(9,1),(10,1),(11,1),(12,1)},{okey,okey,(13,1)}, score is 166.

				case: {{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1},{7,3},{8,3},{9,3},{10,3},{11,3},{12,3},{13,3}} with 1 okey
				the best solution is {(7,1),(8,1),(9,1),(10,1),(11,1),(12,1)},{(7,3),(8,3),(9,3),(10,3),(11,3),(12,3)},{okey,(13,1),(13,3)}, score is 153.

				case: {{7,1},{8,1},{9,1},{10,1},{11,1},{12,1},{13,1},{7,3},{8,3},{9,3},{10,3},{11,3},{12,3},{13,3}} with 2 okey
				the best solution is {(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1)},{(7,3),(8,3),(9,3),(10,3),(11,3),(12,3)},{okey,okey,(13,3)}, score is 166.
		*/
		if ( 1 < play.min && mark.sequence_upper_bounds[CARD_OFFSET(play.min-1, play.suit)] ) // combinable sequence
		{
			// seems play.min can be any number (just a guess), and none of case is able to show it is incorrect.
			if ( play.min <= MAX_N - 2 ) // make sure play.count >= 3 when play.max is enlarged to MAX_N
				return true;
		}
		mark.sequence_upper_bounds[CARD_OFFSET(play.max, play.suit)] = true;

		break;
	case PlayPreferToTriplet:
		/*
			play can't be any triplet.

			the following examples will be deal with.
				case: {{9,1},{9,2},{9,3},{11,2},{10,1},{10,2},{10,3}} with 1 okey
				the best solution is {{(9,1),(9,2),(9,3)},{okey,(10,1),(10,2),(10,3)},{(11,2)}}, score is 67.

				case: {{9,1},{9,2},{9,3},{9,4},{11,2},{10,1},{10,2},{10,3}} with 1 okey
				the best solution is {{(9,1),(9,3),(9,4)},{(9,2),(10,2),(11,2)},{okey,(10,1),(10,3)}}, score is 87.

				case: {{9,1},{9,2},{9,3},{9,4},{9,2},{11,2},{10,1},{10,2},{10,3}} with 1 okey
				the best solution is {{(9,1),(9,2),(9,3),(9,4)},{(9,2),(10,2),(11,2)},{okey,(10,1),(10,3)}}, score is 96.

				case: {{10,1},{10,2},{10,3},{10,4},{11,1},{12,1}} with 0 okey
				the best solution is {{(10,1),(11,1),(12,1)},{(10,2),(10,3),(10,4)}}, score is 63.

				case: {{10,2},{10,2},{10,3},{10,4},{11,2},{12,2},{11,2}} with 0 okey
				the best solution is {{(10,2),(10,3),(10,4)},{(10,2),(11,2),(12,2)},{(11,2)}}, score is 63.

				case: {{10,1},{10,2},{10,3},{10,4},{12,1},{12,2},{13,2},{11,1}} with 1 okey
				the best solution is {{(10,1),(11,1),(12,1)},{(10,2),(10,3),(10,4)},{okey,(12,2),(13,2)}}, score is 99.

			the following examples will be not deal with.
				case: {{10,1},{10,2},{12,1},{12,2},{13,2},{11,3}} with 2 okey
				the best solution is {{(10,1),okey,(12,1)},{(10,2),okey,(12,2),(13,2)},{(11,3)}}, score is 79.

				case: {{10,1},{10,2},{10,3},{10,4},{11,3},{12,3},{11,2}} with 0 okey
				the best solution is {{(10,1),(10,2),(10,4)},{(10,3),(11,3),(12,3)},{(11,2)}}, score is 63.

				case: {{9,2},{9,4},{11,4},{10,1},{10,2},{10,3}} with 1 okey
				the best solution is {(9,2)},{(9,4),okey,(11,4)},{(10,1),(10,2),(10,3)}, score is 60.
		*/
		if ( play.count < 3 && play.max < card.first ) // bad triplet and need add okey to score up
		{
			bool suit_available = false;
			if ( mark.triplet_suit_counts[play.max] < card_suit_upper_bounds[play.max] )
			{
				for (int suit = 1; suit <= card_suit_upper_bounds[play.max]; ++suit)
				{
					if ( play.suits[suit] )
					{
						int card_offset = CARD_OFFSET(play.max, suit);
						if ( !mark.triplet_suit_markups[card_offset] )
						{
							suit_available = true;
							++mark.triplet_suit_counts[play.max];
							mark.triplet_suit_markups[card_offset] = true;

							// two single triplets are incombinable
							if ( 1 == play.count && mark.triplet_suit_counts[play.max] < 3
								&& play.suit < card_suit_upper_bounds[play.max] )
							{	// refer to ( 2 == play.count ) in IsNextWasteOnSinglePlay()
								suit_available = false;
							}
						}
					}
				}
			}
			if ( !mark.triplets[play.max] )
				mark.triplets[play.max] = true;
			else
			{
				if ( suit_available ) // combinable triplet
					return true;
			}
		}

		break;
	default:
		;
	}

	return false;
}

void OKeySolver::Search()
{
//    //cout<<"here"<<endl;
	// if ( hand_card_index >= (int)hand_cards.size() )
	//cout<<"xxxxxxxxxxxxxxxxxxxxxxxx"<<endl;
	if ( hand_card_index >= (int)hand_cards.size() && retrieved_card_index >= (int)retrieved_cards.size() )
	{
	    //cout<<"retrieve"<<endl;
		Retrieve();
		// Evaluate(solution);
		return;
	}

	// const Card &card = hand_cards[hand_card_index];
	//cout<<endl;
	//cout<<"index"<<hand_card_index;
	//dump(hand_cards);
	//cout<<endl;
	const Card &card = ( hand_card_index < (int)hand_cards.size() ) ? hand_cards[hand_card_index] : retrieved_cards[retrieved_card_index];
	int card_offset = CARD_OFFSET(card);
	int last_play_index = last_play_indexes[card_offset];
	auto play_preferences_bak = play_preferences;
	WasteMark waste_mark = {{0},{0},{0},{0}};
	bool is_next_waste = false;
	// if ( 0 < last_play_index )
	// 	is_next_waste = IsNextWasteOnMultiplePlays(solution.plays[last_play_index-1], card, play_preferences[last_play_index-1], waste_mark);
	// ++hand_card_index;
	( hand_card_index < (int)hand_cards.size() ) ? (++hand_card_index) : (++retrieved_card_index);
	int solution_size = solution.plays.size();
	for ( int play_index = last_play_index; play_index < solution_size; ++play_index )
	{
		if ( !(0 < play_index && solution.plays[play_index-1] == solution.plays[play_index]) ) // not same plays
		{
			bool okey_check_pass = true;
			int okey = CheckPlay(solution.plays[play_index], card);
			// if ( okey < 0 || solution.okey < okey )
			if ( okey < 0 || 1 < okey || solution.okey < okey )
			{
				okey_check_pass = false;
			}
			else
			if ( 0 < okey )
			{
				int wild = 0;
				for ( int okey_i = okey; okey_i >= 1; --okey_i )
				{
					Card okey_card = {card.first-okey_i, card.second};
					if ( solution.okey <= 1 )
						if ( play_index < last_play_indexes[CARD_OFFSET(okey_card)] )
						{
							okey_check_pass = false;
							break;
						}
					if ( 0 == card_counts[CARD_OFFSET(okey_card)] )
						++wild;
				}
				if ( wild && solution.plays[play_index].wild ) // not allow use wild card twice
					okey_check_pass = false;
			}

			if ( okey_check_pass )
			{
				for ( int okey_i = okey; okey_i >= 1; --okey_i )
				{
					// AddPlay(solution.plays[play_index], Card{card.first-okey_i, 0});
					Card wild_card = {card.first-okey_i, 0};
					//cout<<"addPlay1"<<play_index;
					//dump(wild_card);
					//cout<<"solution";
					for(int i=0; i<solution.plays.size(); i++){
					    //dump(solution.plays[i]);
					}
					AddPlay(solution.plays[play_index], wild_card);
					if ( 0 == card_counts[CARD_OFFSET(wild_card.first, card.second)] )
						++solution.plays[play_index].wild;
				}
				solution.okey -= okey;
				assert(solution.plays[play_index].wild <= 1);
				//cout<<"addPlay2";
				//dump(card);
				//cout<<play_index;
				AddPlay(solution.plays[play_index], card);
				last_play_indexes[card_offset] = play_index;
				//cout<<"enter0"<<endl;
				Search();
				RemovePlay(solution.plays[play_index], card);
				for ( int okey_i = 1; okey_i <= okey; ++okey_i )
				{
					// RemovePlay(solution.plays[play_index], Card{card.first-okey_i, 0});
					Card wild_card = {card.first-okey_i, 0};
					RemovePlay(solution.plays[play_index], wild_card);
					if ( 0 == card_counts[CARD_OFFSET(wild_card.first, card.second)] )
						--solution.plays[play_index].wild;
				}
				solution.okey += okey;

				// is_next_waste = IsNextWasteOnSinglePlay(solution.plays[play_index], card, okey);
				is_next_waste = IsNextWasteOnSinglePlay(solution.plays[play_index], card, okey, solution.okey);
				if ( is_next_waste )
					break;
			}
		}

		// is_next_waste = IsNextWasteOnMultiplePlays(solution.plays[play_index], card, play_preferences[play_index], waste_mark);
		if ( is_next_waste )
			break;
	}
	if ( !is_next_waste )
	{
		Play play_new = {0,0,0,0,0,{0},0};

		AddPlay(play_new, card);
		//cout<<"addPlay3";
        //dump(card);
		solution.plays.push_back(play_new);
		play_preferences.push_back(PlayPreferNone);
		last_play_indexes[card_offset] = solution_size;
		//cout<<"enter1"<<endl;
		Search();
		solution.plays.pop_back();
		play_preferences.pop_back();
	}
	// --hand_card_index;
	( 0 == retrieved_card_index ) ? (--hand_card_index) : (--retrieved_card_index);
	last_play_indexes[card_offset] = last_play_index;
	play_preferences = play_preferences_bak; // prevent play preferences from become dirt
}

void OKeySolver::Retrieve()
{
	int solo_ace_meld_count = 0;
	if ( has_ace_card_in_hand && retrieved_cards.empty() )
	{
		for ( int i = 0; i < (int)solution.plays.size(); ++i )
			if ( solution.plays[i].count == 1 && solution.plays[i].max == 1 ) // solo ace card
			{
				retrieved_cards.push_back(Card(14, solution.plays[i].suit));
				solution.plays[i].count = -1;
				++solo_ace_meld_count;
			}
	}

	if ( solo_ace_meld_count > 0 ) // retrieving
	{
		Search();

		for ( int i = 0; i < (int)solution.plays.size(); ++i )
			if ( solution.plays[i].count == -1 && solution.plays[i].max == 1 ) // solo ace card
			{
				solution.plays[i].count = 1;
			}
		// hand_cards.resize(hand_cards.size() - solo_ace_meld_count);
		retrieved_cards.clear();
	}
	else
		Evaluate(solution);
}

Solution OKeySolver::Solve(const Deck &cards, int okey)
{
	hand_cards = cards;
	sort(hand_cards.begin(), hand_cards.end());
	has_ace_card_in_hand = false;
	card_counts.resize(MAX_LEN, 0);
	card_suit_upper_bounds.resize(MAX_N+1, 0);
	for ( const auto &card : hand_cards )
	{
		++card_counts[CARD_OFFSET(card)];
		card_suit_upper_bounds[card.first] = max(card_suit_upper_bounds[card.first], card.second);
		if ( 1 == card.first ) // ace card
		{
			card_suit_upper_bounds[14] = max(card_suit_upper_bounds[14], card.second);
			has_ace_card_in_hand = true;
		}
	}
	retrieved_cards.clear();
	last_play_indexes.resize(MAX_LEN, 0);
	play_preferences.clear();
	solution.plays.clear();
	solution.okey = okey;
	best_score_solution.plays.clear();
	best_score_solution.okey = 0;
	best_score = 0;
	best_count_in_best_score = 0;
	best_count_solution.plays.clear();
	best_count_solution.okey = 0;
	best_count = 0;
	best_score_in_best_count = 0;
	hand_card_index = 0;
	retrieved_card_index = 0;
	Search();
	return best_score_solution;
}

Deck Play2Deck(const Play &play)
{
	Deck deck;
	Play _play = play;
	if ( !_play.interval )
	{
		_play.suit = 0;
		while ( !_play.suits[_play.suit] )
			++_play.suit;
	}
	while ( 0 < _play.count-- )
	{
		deck.push_back({_play.min, _play.suit});
		_play.min += _play.interval;
		if ( !_play.interval )
		{
			--_play.suits[_play.suit];
//			//cout<<endl;
//			//cout<<"start:"<<_play.suit<<" ";
//			for(int i=0; i<6;i++){
//			    //cout<<_play.suits[i];
//			}
//			//cout<<endl;
			while ( !_play.suits[_play.suit] ){
			    _play.suit++;
//			    for(int i=0; i<10;i++){
//                    //cout<<_play.suits[i];
//                }
//                //cout<<" "<<_play.suit<<endl;
			}
			//cout<<endl;

		}
	}
	return deck;
}

vector<Deck> Solution2Decks(const Solution &solution, const Deck &cards)
{
//    cout<<"cpp Solution2Decks"<<endl;
//    dump(solution);
//    dump(cards);
//    cout<<"cpp Solution2Decks end"<<endl;
	vector<Deck> decks;
	vector<int> card_counts(MAX_LEN, 0);
	vector<pair<int,int>> card_last_indexes(MAX_LEN, {-1,-1});
	for ( const auto &card : cards )
		++card_counts[CARD_OFFSET(card)];
	int meld_index = 0;
	for ( const Play& play : solution.plays )
	{
		int wild = 0;
		int card_in_meld_index = 0;
		Deck deck = Play2Deck(play);
		for ( auto &card : deck )
		{
			if ( card.second > 0 )
			{
				if ( 14 == card.first ) // ace card
					card.first = 1;
				int card_offset = CARD_OFFSET(card);
				if ( card_counts[card_offset] )
				{
					--card_counts[card_offset];
					card_last_indexes[card_offset] = {meld_index, card_in_meld_index};
				}
				else
				{
					// card.second = 0;
					card.second *= -1;
				}
			}
			if ( card.second <= 0 )
				++wild;
			++card_in_meld_index;
		}

		if ( wild > 1 ) // not allow use wild card twice
		{
			// for ( auto &card : deck )
			for ( int i = deck.size()-1; i >=0; --i )
			{
				auto &card = deck[i];
				if ( card.second < 0 ) // wild card
				{
					Card normal_card = {card.first, -1*card.second};
					int card_offset = CARD_OFFSET(normal_card);
					const auto &last_index = card_last_indexes[card_offset];
					if ( last_index.first >= 0 && last_index.second >= 0 )
					{
						// swap normal card and wild card
						assert(last_index.first < meld_index);
						decks[last_index.first][last_index.second].second *= -1;
						card.second *= -1;
						--wild;
						break;
					}
				}
			}
			assert(wild <= 1);
		}

		decks.push_back(move(deck));
		++meld_index;
	}
	for ( auto &deck : decks )
		for ( auto &card : deck )
			if ( card.second < 0 ) // wild card
				card.second = 0;

//	for(auto &deck:decks){
//	    cout<<"cpp after deck"<<endl;
//	    dump(deck);
//	}
	return decks;
}

};