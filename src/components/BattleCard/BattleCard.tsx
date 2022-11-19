import React, {FC} from 'react';
import Card from 'react-bootstrap/Card';

import {Battle, BattleState} from '../../lib/api/spirits/v1/battle.pb';

interface BattleCardProps {
  battle: Battle
}

const BattleCard: FC<BattleCardProps> = (props) => {
  const getBattleState = (): string => {
    switch (props.battle.state) {
      case BattleState.BATTLE_STATE_PENDING:
        return 'Waiting for teams...';
      case BattleState.BATTLE_STATE_STARTED:
        return 'Running...';
      case BattleState.BATTLE_STATE_WAITING:
        return 'Waiting for input...';
      case BattleState.BATTLE_STATE_FINISHED:
        return 'Finished';
      case BattleState.BATTLE_STATE_ERROR:
        return `Error: ${props.battle.errorMessage}`;
      case BattleState.BATTLE_STATE_CANCELLED:
        return 'Cancelled';
      default:
        return '???';
    }
  };

  return (
    <Card>
      <Card.Header>{getBattleState()}</Card.Header>
      <Card.Body>
        <Card.Title>{props.battle.meta?.id}</Card.Title>
      </Card.Body>
    </Card>
  );
};

export default BattleCard;
