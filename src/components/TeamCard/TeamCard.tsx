import React, {FC} from 'react';
import {Card} from 'react-bootstrap';

import {BattleTeam, BattleTeamSpirit} from '../../lib/api/spirits/v1/battle.pb';
import SpiritCard from '../SpiritCard/SpiritCard';

interface TeamCardProps {
  team: BattleTeam
}

const TeamCard: FC<TeamCardProps> = (props) => {
  return (
    <Card>
      <Card.Header>{props.team.name}</Card.Header>
      <Card.Body>
        {props.team.spirits!.map((spirit: BattleTeamSpirit, i: number) =>
          <SpiritCard spirit={spirit.spirit!} key={i} />,
        )}
      </Card.Body>
    </Card>
  );
};

export default TeamCard;
