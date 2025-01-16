import React from 'react';
import PlayingCard from './Card';

interface CardDeckProps {
  onCardSelect: (card: string) => void;
}

const suits = ['♠', '♥', '♦', '♣'];
const values = ['2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K', 'A'];

const CardDeck: React.FC<CardDeckProps> = ({ onCardSelect }) => {
  return (
    <div className="card-deck">
      {suits.map(suit => (
        <div key={suit} className="suit-group">
          <div className="cards">
            {values.map(value => (
              <PlayingCard
                key={`${value}${suit}`}
                value={value}
                suit={suit}
                onClick={() => onCardSelect(`${value}${suit}`)}
              />
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};

export default CardDeck;