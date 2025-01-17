import React, { useState, useEffect } from 'react';
import './App.css';
import CardDeck from './components/CardDeck';
import CardBoard from './components/CardBoard';
import HandEvaluation from './components/HandEvaluation';
import { API_URL } from './config';

interface GameResult {
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

function App() {
  const [handCards, setHandCards] = useState<string[]>([]);
  const [tableCards, setTableCards] = useState<string[]>([]);
  const [currentHand, setCurrentHand] = useState<GameResult | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleCardSelect = (card: string) => {
    if (handCards.length < 2) {
      setHandCards(prevCards => [...prevCards, card]);
    } else if (tableCards.length < 5) {
      setTableCards(prevCards => [...prevCards, card]);
    }
  };

  useEffect(() => {
    if (handCards.length === 2 && tableCards.length >= 3) {
      handleEvaluateHand();
    }
  }, [handCards, tableCards]);

  const convertCardFormat = (card: string): string => {
    const rankMap: { [key: string]: string } = {
      'A': 'A', 'K': 'K', 'Q': 'Q', 'J': 'J', '10': '10',
      '9': '9', '8': '8', '7': '7', '6': '6', '5': '5', '4': '4', '3': '3', '2': '2'
    };
    const suitMap: { [key: string]: string } = {
      '♥': 'H', '♠': 'S', '♦': 'D', '♣': 'C'
    };

    const rank = card.slice(0, -1);
    const suit = card.slice(-1);

    return `${rankMap[rank] || rank}${suitMap[suit] || suit}`;
  };

  const handleEvaluateHand = async () => {
    setLoading(true);
    setError(null);

    try {
      const convertedHandCards = handCards.map(convertCardFormat);
      const convertedTableCards = tableCards.map(convertCardFormat);

      const response = await fetch(`${API_URL}/evaluate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          playerCards: convertedHandCards,
          tableCards: convertedTableCards
        }),
      });

      if (!response.ok) {
        throw new Error('Failed to evaluate hand');
      }

      const result = await response.json();
      setCurrentHand(result);
    } catch (error) {
      console.error('Error evaluating hand:', error);
      setError('Failed to evaluate hand. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const resetSelection = () => {
    setHandCards([]);
    setTableCards([]);
    setCurrentHand(null);
  };

  return (
    <div className="App">
      <h1>Poker Hand Evaluator</h1>
      <div className="game-container">
        <div className="left-side">
          <CardDeck onCardSelect={handleCardSelect} />
        </div>
        <div className="right-side">
          <CardBoard handCards={handCards} tableCards={tableCards} />
          <div className="controls">
            <button onClick={resetSelection}>Reset</button>
          </div>
          {error && <div className="error-message">{error}</div>}
          {loading && <p>Evaluating hand...</p>}
          {currentHand && (
            <HandEvaluation
              hand={currentHand.hand}
              handRank={currentHand.handRank}
              cards={currentHand.cards} potentialBetterHands={currentHand.potentialBetterHands}            />
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
