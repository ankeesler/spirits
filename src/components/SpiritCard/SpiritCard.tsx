import React, {FC} from 'react';
import {Card} from 'react-bootstrap';
import ProgressBar from 'react-bootstrap/ProgressBar';

import {Spirit} from '../../lib/api/spirits/v1/spirit.pb';

interface SpiritCardProps {
  spirit: Spirit
}

const SpiritCard: FC<SpiritCardProps> = (props) => {
  const getSpiritHealth = () => {
    const health = props.spirit.stats?.health;
    return `${health}/${health}`;
  };

  return (
    <Card>
      <Card.Header>{props.spirit.name}</Card.Header>
      <Card.Body>
        <ProgressBar now={100} variant="success" label={getSpiritHealth()} />
      </Card.Body>
    </Card>
  );
};

export default SpiritCard;
