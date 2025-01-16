import React from 'react';

interface PlayingCardProps {
  value: string;
  suit: string;
  onClick?: () => void;
}

const PlayingCard: React.FC<PlayingCardProps> = ({ value, suit, onClick }) => {
  const isRed = suit === '♥' || suit === '♦';

  return (
    <button 
      onClick={onClick}
      className="card"
    >
      <div className="card-inner">
        <div className="card-top-left">
          <span className={isRed ? 'red' : 'black'}>{value}</span>
        </div>
        <div className="card-center">
          <span className={`card-suit ${isRed ? 'red' : 'black'}`}>{suit}</span>
        </div>
        {/* Removed all rotation-related classes from bottom-right */}
        <div className="card-bottom-right" style={{ transform: 'none' }}>  
          <span className={isRed ? 'red' : 'black'}>{value}</span>
        </div>
      </div>
    </button>
  );
};

export default PlayingCard;