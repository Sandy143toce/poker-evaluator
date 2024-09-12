import React, { useState } from 'react';

interface CardInputFormProps {
  onSubmit: (playerCards: string[], tableCards: string[]) => void;
}

const CardInputForm: React.FC<CardInputFormProps> = ({ onSubmit }) => {
  const [playerCards, setPlayerCards] = useState<string[]>(['', '']);
  const [tableCards, setTableCards] = useState<string[]>(['', '', '', '', '']);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const validPlayerCards = playerCards.filter(card => card.trim() !== '');
    const validTableCards = tableCards.filter(card => card.trim() !== '');
    if (validPlayerCards.length === 2 && validTableCards.length >= 3) {
      onSubmit(validPlayerCards, validTableCards);
    } else {
      alert('Please enter 2 player cards and at least 3 table cards.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <h3>Player Cards</h3>
        {playerCards.map((card, index) => (
          <input
            key={`player-${index}`}
            type="text"
            value={card}
            onChange={(e) => {
              const newCards = [...playerCards];
              newCards[index] = e.target.value;
              setPlayerCards(newCards);
            }}
            placeholder={`Player Card ${index + 1}`}
          />
        ))}
      </div>
      <div>
        <h3>Table Cards</h3>
        {tableCards.map((card, index) => (
          <input
            key={`table-${index}`}
            type="text"
            value={card}
            onChange={(e) => {
              const newCards = [...tableCards];
              newCards[index] = e.target.value;
              setTableCards(newCards);
            }}
            placeholder={`Table Card ${index + 1}`}
          />
        ))}
      </div>
      <button type="submit">Evaluate Hand</button>
    </form>
  );
};

export default CardInputForm;