import React from 'react';
import PlayingCard from './Card';

interface PotentialBetterHand {
  name: string;
  handRank: number;
  neededCards: string[];
}

interface PotentialHandsProps {
  potentialHands: PotentialBetterHand[];
}

const PotentialHands: React.FC<PotentialHandsProps> = ({ potentialHands }) => {
  const convertCard = (card: string) => {
    const suitMap: { [key: string]: string } = {
      'H': '♥', 
      'S': '♠', 
      'D': '♦', 
      'C': '♣'
    };
    
    const rank = card.slice(0, -1);
    const suit = suitMap[card.slice(-1)];
    
    return { value: rank, suit };
  };

  if (!potentialHands || potentialHands.length === 0) {
    return <p>No better hands possible.</p>;
  }

  return (
    <div className="potential-hands">
      <h3>Potential Better Hands</h3>
      <div className="potential-hands-list">
        {potentialHands.map((hand, index) => (
          <div key={index} className="potential-hand-item">
            <p className="potential-hand-name">
              <strong>{hand.name}</strong>
            </p>
            <div className="potential-hand-cards">
              <p>Possible with:</p>
              <div className="evaluation-cards">
                {hand.neededCards.map((card, cardIndex) => {
                  const { value, suit } = convertCard(card);
                  return (
                    <PlayingCard
                      key={cardIndex}
                      value={value}
                      suit={suit}
                      small={true}
                    />
                  );
                })}
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default PotentialHands;