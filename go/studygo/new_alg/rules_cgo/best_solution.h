#pragma once

#include <vector>

namespace OKey {

static const int MAX_N = 14;
static const int MAX_M = 4;
static const int MAX_LEN = MAX_N*MAX_M;

static const int ScoreMap[] = {0,11,2,3,4,5,6,7,8,9,10,10,10,10,11};

typedef std::pair<int,int> Card;
typedef std::vector<Card> Cards;
typedef Cards Deck;
struct Play
{
	int count;
	int min;
	int max;
	int interval;
	int suit;
	int suits[MAX_M+1];
	int wild;

	inline bool operator == (const Play &other) const
	{

		if ( this->count != other.count )
			return false;
		if ( this->count <= 0 )
			return true;

		if ( this->interval != other.interval )
			return false;
		if ( this->min != other.min )
			return false;
		if ( this->max != other.max )
			return false;
		if ( this->wild != other.wild )
			return false;

		if ( this->interval )
			return (this->suit == other.suit);

		return true;
	};
};
struct Solution
{
	std::vector<Play> plays;
	int okey;
};

inline int CARD_OFFSET(const Card &card) { return (card.first-1)*MAX_M + (card.second-1); }
inline int CARD_OFFSET(int first, int second) { return (first-1)*MAX_M + (second-1); }

class OKeySolver
{
public:
	Solution Solve(const Deck &cards, int okey = 2);
	int Score(const Solution &solution);
	int Score(const Play& play);
	int Count(const Solution &solution);
	int Count(const Play& play);

	Solution GetBestScoreSolution();
	Solution GetBestCountSolution();

protected:
	void Search();
	void Retrieve();

	int CheckPlay(const Play &play, const Card &card);
	void AddPlay(Play &play, const Card &card);
	void RemovePlay(Play &play, const Card &card);

	void Evaluate(const Solution &solution);
	bool AddBest(const Solution &solution);
	int MakeBestPlay(Play &play, int okey);

	enum PlayPreference
	{
		PlayPreferNone = 0,
		PlayPreferToSequence,
		PlayPreferToTriplet,
	};
	struct WasteMark
	{
		bool triplets[MAX_N+1];
		int triplet_suit_counts[MAX_N+1];
		bool triplet_suit_markups[MAX_LEN];
		bool sequence_upper_bounds[MAX_LEN];
	};
	bool IsNextWasteOnSinglePlay(const Play &play, const Card& card, int okey_cost, int okey_remain);
	bool IsNextWasteOnMultiplePlays(const Play &play, const Card& card, PlayPreference &preference, WasteMark &mark);

private:
	Deck hand_cards;
	std::vector<int> card_counts;
	std::vector<int> card_suit_upper_bounds;
	bool has_ace_card_in_hand;

	Cards retrieved_cards;

	int hand_card_index;
	int retrieved_card_index;

	Solution solution;
	std::vector<int> last_play_indexes;
	std::vector<PlayPreference> play_preferences;

	Solution best_score_solution;
	int best_score;
	int best_count_in_best_score;
	
	Solution best_count_solution;
	int best_count;
	int best_score_in_best_count;
};

Deck Play2Deck(const Play &play);
std::vector<Deck> Solution2Decks(const Solution &solution, const Deck &cards);

void dump(const Deck &cards);
void dump(const Play &play);
void dump(const Solution &solution);
void dump(const std::vector<Deck> &decks);

};