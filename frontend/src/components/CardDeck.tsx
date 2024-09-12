import React from 'react';

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
              <button
                key={`${value}${suit}`}
                className={`card ${suit === '♥' || suit === '♦' ? 'red' : 'black'}`}
                onClick={() => onCardSelect(`${value}${suit}`)}
              >
                <div className="card-inner">
                  <div className="card-top-left">
                    <span className="card-value">{value}</span>
                  </div>
                  <div className="card-center">
                    <span className="card-suit">{suit}</span>
                  </div>
                  <div className="card-bottom-right">
                    <span className="card-value">{value}</span>
                  </div>
                </div>
              </button>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};

export default CardDeck;