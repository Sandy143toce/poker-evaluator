import React from 'react';

interface GameResult {
  hand: string;
  handRank: number;
  cards: string[];
}

interface RecentResultsProps {
  results: GameResult[];
}

const RecentResults: React.FC<RecentResultsProps> = ({ results }) => {
  if (!results || results.length === 0) {
    return <p>No recent results available.</p>;
  }

  return (
    <div className="recent-results">
      <h2>Recent Results</h2>
      <ul>
        {results.map((result, index) => (
          <li key={index}>
            <strong>Hand:</strong> {result.hand}, <strong>Rank:</strong> {result.handRank},
            <strong>Cards:</strong> {result.cards.join(', ')}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default RecentResults;