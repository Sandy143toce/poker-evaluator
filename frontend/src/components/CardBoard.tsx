import React from 'react';
import PlayingCard from './Card';  // Import the PlayingCard component

interface CardBoardProps {
  handCards: string[];
  tableCards: string[];
}

const CardBoard: React.FC<CardBoardProps> = ({ handCards, tableCards }) => {
  // Helper function to split card string into value and suit
  const splitCard = (card: string): { value: string, suit: string } => {
    const value = card.slice(0, -1);
    const suit = card.slice(-1);
    return { value, suit };
  };

  return (
    <div className="card-board">
      <div className="instructions">
        <h3>Instructions:</h3>
        <p>1. Select 2 hand cards first.</p>
        <p>2. Then select 3 to 5 table cards.</p>
        <p>3. Minimum 3 table cards are required.</p>
      </div>
      <div className="board-section">
        <h3>Hand Cards</h3>
        <div className="card-slots hand-cards">
          {Array.from({ length: 2 }).map((_, index) => (
            <div key={`hand-${index}`} className="card-slot">
              {handCards[index] ? (
                <PlayingCard 
                  {...splitCard(handCards[index])}
                />
              ) : ''}
            </div>
          ))}
        </div>
      </div>
      <div className="board-section">
        <h3>Table Cards</h3>
        <div className="card-slots table-cards">
          {Array.from({ length: 5 }).map((_, index) => (
            <div key={`table-${index}`} className="card-slot">
              {tableCards[index] ? (
                <PlayingCard 
                  {...splitCard(tableCards[index])}
                />
              ) : ''}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default CardBoard;