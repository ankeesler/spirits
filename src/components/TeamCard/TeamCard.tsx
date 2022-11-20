import React, {FC} from 'react';
import {Card} from 'react-bootstrap';
import CardGroup from 'react-bootstrap/CardGroup';

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
        <CardGroup>
          {props.team.spirits!.map((spirit: BattleTeamSpirit, i: number) =>
            <SpiritCard spirit={spirit.spirit!} key={i} />,
          )}
        </CardGroup>
      </Card.Body>
    </Card>
  );
};

export default TeamCard;
