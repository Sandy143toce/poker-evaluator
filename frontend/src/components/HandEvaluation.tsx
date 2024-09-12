import React from 'react';

interface HandEvaluationProps {
  hand: string;
  handRank: number;
  cards: string[];
}

const HandEvaluation: React.FC<HandEvaluationProps> = ({ hand, handRank, cards }) => {
  return (
    <div className="hand-evaluation">
      <h3>Hand Evaluation</h3>
      <p><strong>Hand:</strong> {hand}</p>
      <p><strong>Rank:</strong> {handRank}</p>
      <p><strong>Cards:</strong> {cards.join(', ')}</p>
    </div>
  );
};

export default HandEvaluation;