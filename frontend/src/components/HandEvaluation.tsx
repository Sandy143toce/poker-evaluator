import React from 'react';
import PlayingCard from './Card';
import PotentialHands from './PotentialHands'; 

interface HandEvaluationProps {
  hand: string;
  handRank: number;
  cards: string[];
  potentialBetterHands: PotentialBetterHand[]; 
}

interface PotentialBetterHand {
  name: string;
  handRank: number;
  neededCards: string[];
}

const HandEvaluation: React.FC<HandEvaluationProps> = ({ hand, handRank, cards, potentialBetterHands }) => {
  // Helper function to convert card format (e.g., "10D" to {value: "10", suit: "♦"})
  const convertCard = (card: string) => {
    const suitMap: { [key: string]: string } = {
      'H': '♥', 
      'S': '♠', 
      'D': '♦', 
      'C': '♣'
    };
    
    // Extract rank and suit
    const rank = card.slice(0, -1);
    const suit = suitMap[card.slice(-1)];
    
    return { value: rank, suit };
  };

  return (
    <div className="hand-evaluation">
      <h3>Hand Evaluation</h3>
      <p><strong>Hand:</strong> {hand}</p>
      <p><strong>Rank:</strong> {handRank}</p>
      <div>
        <strong>Current Hand Cards:</strong>
        <div className="evaluation-cards">
          {cards.map((card, index) => {
            const { value, suit } = convertCard(card);
            return (
              <PlayingCard
                key={index}
                value={value}
                suit={suit}
                small={true}  // New prop for smaller size
              />
            );
          })}
        </div>
      </div>
      <PotentialHands potentialHands={potentialBetterHands} />
    </div>
  );
};

export default HandEvaluation;