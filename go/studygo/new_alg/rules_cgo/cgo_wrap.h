//
// Created by Tommy on 2020/9/15.
//

#ifndef OKEY_CGO_CGO_WRAP_H
#define OKEY_CGO_CGO_WRAP_H

#include <vector>
#include "best_solution.h"

const int MIN_PROJECT_SIZE = 3;

class Result{
public:
    std::vector<int> solution;              // old version solution
    std::vector<int> bestCountSolution;     // new version solution
};

std::vector<int> packSolution(const OKey::Solution &solution, const OKey::Deck &cards){
    std::vector<int> result;
    for(auto& deck: OKey::Solution2Decks(solution, cards)){
        if(deck.size() < MIN_PROJECT_SIZE){
            continue;
        }
        for(auto c : deck){
            result.emplace_back(c.first);
            result.emplace_back(c.second);
        }
        result.emplace_back(1024);
        result.emplace_back(1024);
    }
    return result;
}

Result GetSolution(std::vector<int>& cards, int okey){
    OKey::Deck cardDeck;
    cardDeck.reserve(cards.size()/2);
    for(auto i=0;i<cards.size()/2;i++){
        cardDeck.emplace_back(cards[i*2], cards[i*2+1]);
    }
    OKey::OKeySolver solver;
    Result result;
    result.solution = packSolution(solver.Solve(cardDeck, okey), cardDeck);
    result.bestCountSolution = packSolution(solver.GetBestCountSolution(), cardDeck);
    return result;
}

#endif //OKEY_CGO_CGO_WRAP_H
