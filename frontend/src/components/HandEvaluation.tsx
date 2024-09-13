import React from 'react';

interface Hand {
  name: string;
  rank: number;
  cards: string[];
  sequence: number[];
}

interface HandEvaluationProps {
  playerBestHand: Hand;
  otherBestHands: Hand[];
}

const HandEvaluation: React.FC<HandEvaluationProps> = ({ playerBestHand, otherBestHands }) => {
  const renderHand = (hand: Hand, isPlayerHand: boolean) => (
    <div className={`hand ${isPlayerHand ? 'player-hand' : 'other-hand'}`}>
      <h4>{isPlayerHand ? 'Your Best Hand' : 'Possible Other Hand'}</h4>
      <p><strong>Name:</strong> {hand.name}</p>
      <p><strong>Rank:</strong> {hand.rank}</p>
      <p><strong>Cards:</strong> {hand.cards.join(', ')}</p>
      <p><strong>Sequence:</strong> {hand.sequence.join(', ')}</p>
    </div>
  );

  return (
    <div className="hand-evaluation">
      <h3>Hand Evaluation</h3>
      {renderHand(playerBestHand, true)}
      <h4>Other Possible Best Hands</h4>
      {otherBestHands.map((hand, index) => (
        <React.Fragment key={index}>
          {renderHand(hand, false)}
          {index < otherBestHands.length - 1 && <hr />}
        </React.Fragment>
      ))}
    </div>
  );
};

export default HandEvaluation;