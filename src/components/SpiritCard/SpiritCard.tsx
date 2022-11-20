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
      <Card.Img src="https://images.pexels.com/photos/1207875/pexels-photo-1207875.jpeg?cs=srgb&dl=pexels-andre-mouton-1207875.jpg&fm=jpg&_gl=1*1r7ueue*_ga*NDc1NjA5OTE4LjE2Njg5NTMzNjg.*_ga_8JE65Q40S6*MTY2ODk1MzM2OC4xLjEuMTY2ODk1MzUyNi4wLjAuMA.." />
      <Card.Body>
        <ProgressBar now={100} variant="success" label={getSpiritHealth()} />
      </Card.Body>
    </Card>
  );
};

export default SpiritCard;
